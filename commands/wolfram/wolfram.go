package wolfram

import (
	"fmt"
	"github.com/0x263b/porygon2"
	"github.com/0x263b/porygon2/web"
	"gopkg.in/xmlpath.v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	wolframURL = "http://api.wolframalpha.com/v2/query?appid=%s&input=%s"
)

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

func wolfram(command *bot.Cmd, matches []string) (msg string, err error) {
	doc, _ := http.Get(fmt.Sprintf(wolframURL, bot.Config.Wolfram, url.QueryEscape(matches[1])))
	defer doc.Body.Close()
	root, err := xmlpath.Parse(doc.Body)

	if err != nil {
		return "Wolfram | Stephen Wolfram doesn't know the answer to this", nil
	}

	short := web.ShortenURL(fmt.Sprintf("https://www.wolframalpha.com/input/?i=%s", url.QueryEscape(matches[1])))

	success := xmlpath.MustCompile("//queryresult/@success")
	input := xmlpath.MustCompile("//pod[@position='100']//plaintext[1]")
	output := xmlpath.MustCompile("//pod[@position='200']/subpod[1]/plaintext[1]")

	suc, _ := success.String(root)

	if suc != "true" {
		return fmt.Sprintf("Wolfram | Stephen Wolfram doesn't know the answer to this | %s", short), nil
	}

	in, _ := input.String(root)
	out, _ := output.String(root)

	in = strings.Replace(in, `\:`, `\u`, -1)
	out = strings.Replace(out, `\:`, `\u`, -1)

	reg := regexp.MustCompile("\\s+")
	in = reg.ReplaceAllString(in, " ")
	out = reg.ReplaceAllString(out, " ")

	in, _ = strconv.Unquote(`"` + in + `"`)
	out, _ = strconv.Unquote(`"` + out + `"`)

	return fmt.Sprintf("Wolfram | %s >>> %s | %s", in, out, short), nil
}

func init() {
	bot.RegisterCommand(
		"^wa (.+)$",
		wolfram)
}
