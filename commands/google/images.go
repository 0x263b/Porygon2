package google

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/0x263b/Porygon2/web"
	"net/url"
)

const (
	imageURL = "https://ajax.googleapis.com/ajax/services/search/images?v=1.0&q=%s"
)

type ImageResults struct {
	Responsedata struct {
		Results []struct {
			Gsearchresultclass  string `json:"GsearchResultClass"`
			Width               string `json:"width"`
			Height              string `json:"height"`
			Imageid             string `json:"imageId"`
			Tbwidth             string `json:"tbWidth"`
			Tbheight            string `json:"tbHeight"`
			Unescapedurl        string `json:"unescapedUrl"`
			URL                 string `json:"url"`
			Visibleurl          string `json:"visibleUrl"`
			Title               string `json:"title"`
			Titlenoformatting   string `json:"titleNoFormatting"`
			Originalcontexturl  string `json:"originalContextUrl"`
			Content             string `json:"content"`
			Contentnoformatting string `json:"contentNoFormatting"`
			Tburl               string `json:"tbUrl"`
		} `json:"results"`
		Cursor struct {
			Resultcount string `json:"resultCount"`
			Pages       []struct {
				Start string `json:"start"`
				Label int    `json:"label"`
			} `json:"pages"`
			Estimatedresultcount string `json:"estimatedResultCount"`
			Currentpageindex     int    `json:"currentPageIndex"`
			Moreresultsurl       string `json:"moreResultsUrl"`
			Searchresulttime     string `json:"searchResultTime"`
		} `json:"cursor"`
	} `json:"responseData"`
	Responsedetails interface{} `json:"responseDetails"`
	Responsestatus  int         `json:"responseStatus"`
}

func image(command *bot.Cmd, matches []string) (msg string, err error) {
	results := &ImageResults{}
	err = web.GetJSON(fmt.Sprintf(imageURL, url.QueryEscape(matches[1])), results)
	if err != nil {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	if len(results.Responsedata.Results) == 0 {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	output := fmt.Sprintf("Google | %s | %s", matches[1], results.Responsedata.Results[0].URL)
	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^img (.+)$",
		image)
}
