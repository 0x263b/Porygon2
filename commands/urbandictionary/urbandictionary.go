package urbandictionary

import (
	"fmt"
	"github.com/killwhitey/Porygon2"
	"github.com/killwhitey/Porygon2/web"
	"net/url"
	"regexp"
	"strconv"
)

const (
	urbanURL = "http://api.urbandictionary.com/v0/define?term=%s"
)

type DefinitionResults struct {
	Tags       []string `json:"tags"`
	ResultType string   `json:"result_type"`
	List       []struct {
		Defid       int    `json:"defid"`
		Word        string `json:"word"`
		Author      string `json:"author"`
		Permalink   string `json:"permalink"`
		Definition  string `json:"definition"`
		Example     string `json:"example"`
		ThumbsUp    int    `json:"thumbs_up"`
		ThumbsDown  int    `json:"thumbs_down"`
		CurrentVote string `json:"current_vote"`
	} `json:"list"`
	Sounds []interface{} `json:"sounds"`
}

func urban(command *bot.Cmd, matches []string) (msg string, err error) {
	var i int64 = 0
	if matches[1] != "" {
		i, _ = strconv.ParseInt(matches[1], 10, 64)
		i = i - 1
	}

	results := &DefinitionResults{}
	err = web.GetJSON(fmt.Sprintf(urbanURL, url.QueryEscape(matches[2])), results)
	if err != nil {
		return fmt.Sprintf("Urban Dictionary | %s | (No definition found)", matches[2]), nil
	}
	if results.ResultType == "no_results" {
		return fmt.Sprintf("Urban Dictionary | %s | (No definition found)", matches[2]), nil
	}

	word := results.List[i].Word
	definition := results.List[i].Definition
	permalink := results.List[i].Permalink
	short := web.ShortenURL(permalink)

	reg := regexp.MustCompile("\\s+")
	definition = reg.ReplaceAllString(definition, " ") // Strip tabs and newlines

	if len(definition) > 180 {
		definition = fmt.Sprintf("%s...", definition[0:180])
	}

	output := fmt.Sprintf("Urban Dictionary | %s | %s | %s", word, short, definition)

	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^u(?:r(?:ban)?)? (?:([1-7]{1}) )?(.+)$",
		urban)
}
