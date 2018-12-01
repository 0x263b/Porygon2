package url

import (
	"encoding/json"
	"fmt"
	"github.com/0x263b/porygon2"
	"github.com/0x263b/porygon2/web"
	"github.com/PuerkitoBio/goquery"
	"github.com/dustin/go-humanize"
	"github.com/kennygrant/sanitize"
	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
	"html"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	youtubeVideoURL = "https://www.googleapis.com/youtube/v3/videos?part=snippet,contentDetails,statistics&id=%s&key=%s"
)

// 4chan API thread structure
// https://github.com/4chan/4chan-API
type ChanPost struct {
	Posts []struct {
		No  int    `json:"no"`
		Com string `json:"com"`
	} `json:"posts"`
}

// Reddit API thread structure
// https://www.reddit.com/dev/api
type RedditComment []struct {
	Data struct {
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				Body string `json:"body"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

// YouTube API
type youtubeVideo struct {
	Items []struct {
		Snippet struct {
			Title string `json:"title"`
		} `json:"snippet"`
		Contentdetails struct {
			Duration string `json:"duration"`
		} `json:"contentDetails"`
	} `json:"items"`
}

func LoadCredentials() (client *twittergo.Client, err error) {
	config := &oauth1a.ClientConfig{
		ConsumerKey:    bot.Config.TwitterConsumerKey,
		ConsumerSecret: bot.Config.TwitterConsumerSecret,
	}
	client = twittergo.NewClient(config, nil)
	return
}

// Used to parse youtube's ISO 8601 durations
// https://en.wikipedia.org/wiki/ISO_8601#Durations
func ParseDuration(str string) time.Duration {
	durationRegex := regexp.MustCompile(`P(?P<years>\d+Y)?(?P<months>\d+M)?(?P<days>\d+D)?T?(?P<hours>\d+H)?(?P<minutes>\d+M)?(?P<seconds>\d+S)?`)
	matches := durationRegex.FindStringSubmatch(str)

	years := ParseInt64(matches[1])
	months := ParseInt64(matches[2])
	days := ParseInt64(matches[3])
	hours := ParseInt64(matches[4])
	minutes := ParseInt64(matches[5])
	seconds := ParseInt64(matches[6])

	hour := int64(time.Hour)
	minute := int64(time.Minute)
	second := int64(time.Second)
	return time.Duration(years*24*365*hour + months*30*24*hour + days*24*hour + hours*hour + minutes*minute + seconds*second)
}

func ParseInt64(value string) int64 {
	if len(value) == 0 {
		return 0
	}
	parsed, err := strconv.Atoi(value[:len(value)-1])
	if err != nil {
		return 0
	}
	return int64(parsed)
}

var timeout = time.Duration(3) * time.Second

func dialTimeout(network, addr string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, addr, timeout)
	return conn, err
}

func getTweet(arg string) string {
	var (
		err    error
		client *twittergo.Client
		req    *http.Request
		resp   *twittergo.APIResponse
		tweet  *twittergo.Tweet
	)

	client, err = LoadCredentials()
	if err != nil {
		fmt.Printf("Could not parse CREDENTIALS file: %v\n", err)
	}
	url := fmt.Sprintf("/1.1/statuses/show/%s.json?tweet_mode=extended", arg)
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Could not parse request: %v\n", err)
	}
	resp, err = client.SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
	}
	tweet = &twittergo.Tweet{}
	err = resp.Parse(tweet)
	if err != nil {
		if rle, ok := err.(twittergo.RateLimitError); ok {
			fmt.Printf("Rate limited, reset at %v\n", rle.Reset)
		} else if errs, ok := err.(twittergo.Errors); ok {
			for i, val := range errs.Errors() {
				fmt.Printf("Error #%v - ", i+1)
				fmt.Printf("Code: %v ", val.Code())
				fmt.Printf("Msg: %v\n", val.Message())
			}
		} else {
			fmt.Printf("Problem parsing response: %v\n", err)
		}
	}

	reg := regexp.MustCompile("\\s+")
	text := reg.ReplaceAllString(tweet.FullText(), " ") // Strip tabs and newlines
	text = strings.TrimSpace(text)                      // Then trim excessive spaces

	return fmt.Sprintf("%s (@%s): %s", tweet.User().Name(), tweet.User().ScreenName(), text)
}

func extractURL(text string) string {
	extractedURL := ""
	for _, value := range strings.Split(text, " ") {
		parsedURL, err := url.Parse(value)
		if err != nil {
			continue
		}
		if strings.HasPrefix(parsedURL.Scheme, "http") {
			extractedURL = parsedURL.String()
			break
		}
	}
	return extractedURL
}

func openGraphTitle(command *bot.PassiveCmd) (string, error) {
	if !bot.GetChannelKey(command.Channel, "urls") {
		return "", nil
	}

	URL := extractURL(command.Raw)

	if URL == "" {
		return "", nil
	}

	transport := &http.Transport{
		Dial: dialTimeout,
	}

	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar:       cookieJar, // Some sites require cookies to show you anything (nytimes)
		Transport: transport, // Time out if connection hangs
	}

	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set("Accept-Language", "en-US")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:50.0) Gecko/20100101 Firefox/50.0")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	finalURL := response.Request.URL.Host

	var bytes int64 = 40960

	if response.Header.Get("Content-Type") == "" {
		// Some servers don't give us anything to work with
		return fmt.Sprintf("Title | (no title) | %s", finalURL), nil
	} else if !strings.Contains(response.Header.Get("Content-Type"), "text/html") {
		if !bot.GetChannelKey(command.Channel, "files") {
			return "", nil
		}
		contentType := response.Header.Get("Content-Type")
		contentLength := response.Header.Get("Content-Length")
		size, _ := strconv.ParseUint(contentLength, 10, 64)
		return fmt.Sprintf("File | %s %s | %s", contentType, humanize.Bytes(size), finalURL), nil
	}

	defer response.Body.Close()
	body := response.Body
	chunk := io.LimitReader(body, bytes) // Download/Read first 20kB

	doc, err := goquery.NewDocumentFromReader(chunk)
	if err != nil {
		return "", err
	}

	title := doc.Find("title").Text()

	// Generally <meta> tags have more useful titles
	doc.Find("meta[property='og:title']").Each(func(i int, s *goquery.Selection) {
		title = s.AttrOr("content", title)
	})

	// Get tweet content from <meta>
	if finalURL == "twitter.com" || finalURL == "mobile.twitter.com" {
		if response.StatusCode != 200 {
			title = "404 Not Found"
		} else {
			r := regexp.MustCompile(`\/status\/(\d+)`)
			status := r.FindStringSubmatch(response.Request.URL.Path)

			if len(status) > 0 {
				title = getTweet(status[1])
			} else {
				doc.Find("meta[property='og:description']").Each(func(i int, s *goquery.Selection) {
					reg := regexp.MustCompile(`(^“|”\z)`)
					tweet := reg.ReplaceAllString(s.AttrOr("content", title), "") // Strip quotes
					title = fmt.Sprintf("%s: %s", title, tweet)
				})
			}
		}
	}

	// Get YouTube title/duration
	if finalURL == "www.youtube.com" {
		if response.StatusCode != 200 {
			title = "404 Not Found"
		} else {
			r := regexp.MustCompile(`v=(\S{11})`)
			id := r.FindStringSubmatch(response.Request.URL.RawQuery)

			if len(id) > 0 {
				video := &youtubeVideo{}
				err = web.GetJSON(fmt.Sprintf(youtubeVideoURL, id[1], bot.Config.Youtube), video)
				if err == nil {
					if len(video.Items) == 0 {
						title = "404 Not Found"
					} else {
						videoTitle := video.Items[0].Snippet.Title
						duration := ParseDuration(video.Items[0].Contentdetails.Duration)

						title = fmt.Sprintf("%s | %s", videoTitle, duration)
					}
				}
			}
		}
	}

	// Get 4chan post
	if finalURL == "boards.4chan.org" || finalURL == "boards.4channel.org" {
		if strings.Contains(response.Request.URL.Path, "/thread/") {
			path := strings.Split(response.Request.URL.Path, "/")
			postId := response.Request.URL.Fragment

			response, _ := client.Get(fmt.Sprintf("https://a.4cdn.org/%s/thread/%s.json", path[1], path[3]))

			if response.StatusCode != 200 {
				title = "404 Not Found"
			} else {

				defer response.Body.Close()
				body, _ := ioutil.ReadAll(response.Body)

				var posts ChanPost
				json.Unmarshal(body, &posts)

				title = posts.Posts[0].Com

				if postId != "" {
					for _, element := range posts.Posts {
						if strings.Contains(postId, strconv.Itoa(element.No)) {
							title = element.Com
							break
						}
					}
				}

				title = sanitize.HTML(title) // Remove any unwanted html
			}

			if len(title) < 1 {
				title = "(blank post)"
			}
		}
	}

	if finalURL == "www.reddit.com" {
		thread_title := title
		r := regexp.MustCompile(`\/r\/\w+\/comments\/\w+\/\w+\/\w+`)

		if r.MatchString(response.Request.URL.Path) {

			request, _ := http.NewRequest("GET", fmt.Sprintf("https://www.reddit.com%s.json", response.Request.URL.Path), nil)

			request.Header.Set("Accept-Language", "en-US")
			request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:50.0) Gecko/20100101 Firefox/50.0")

			response, _ := client.Do(request)

			if response.StatusCode != 200 {
				title = ""
			} else {

				defer response.Body.Close()

				body, _ := ioutil.ReadAll(response.Body)
				var comments RedditComment
				json.Unmarshal(body, &comments)

				children := comments[len(comments)-1].Data.Children

				if len(children) > 0 {
					title = children[len(children)-1].Data.Body
				}
			}
			if len(title) < 1 {
				title = thread_title
			}
		}
	}

	reg := regexp.MustCompile("\\s+")
	title = reg.ReplaceAllString(title, " ") // Strip tabs and newlines
	title = strings.TrimSpace(title)         // then trim excessive spaces

	if len(title) > 400 {
		title = fmt.Sprintf("%s …", title[0:400])
	} else if len(title) < 1 {
		title = "(no title)"
	}

	return fmt.Sprintf("Title | %s | %s", html.UnescapeString(title), finalURL), nil
}

func toggleURLs(command *bot.Cmd, matches []string) (msg string, err error) {
	if !bot.IsAdmin(command.Nick) || !bot.IsPrivateMsg(command.Channel, command.Nick) {
		return "", nil
	}

	onOff := matches[1]
	channelToToggle := matches[2]

	if onOff == "on" {
		bot.SetChannelKey(channelToToggle, "urls", true)
		return fmt.Sprintf("Reacting to URLs in %s", channelToToggle), nil
	} else if onOff == "off" {
		bot.SetChannelKey(channelToToggle, "urls", false)
		return fmt.Sprintf("No longer displaying page titles in %s", channelToToggle), nil
	}
	return "", nil
}

func toggleFiles(command *bot.Cmd, matches []string) (msg string, err error) {
	if !bot.IsAdmin(command.Nick) || !bot.IsPrivateMsg(command.Channel, command.Nick) {
		return "", nil
	}

	onOff := matches[1]
	channelToToggle := matches[2]

	if onOff == "on" {
		bot.SetChannelKey(channelToToggle, "files", true)
		return fmt.Sprintf("Reacting to file URLs in %s", channelToToggle), nil
	} else if onOff == "off" {
		bot.SetChannelKey(channelToToggle, "files", false)
		return fmt.Sprintf("No longer displaying file info in %s", channelToToggle), nil
	}
	return "", nil
}

func init() {
	bot.RegisterPassiveCommand(
		"opengraph",
		openGraphTitle)

	bot.RegisterCommand(
		"^set urls (\\S+) (\\S+)$",
		toggleURLs)

	bot.RegisterCommand(
		"^set files (\\S+) (\\S+)$",
		toggleFiles)
}
