package parse

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
)

func ParseURL(rawURL string) string {
	urll, err := parseURL(rawURL)
	if err != nil {
		log.Error().
			Caller().
			Err(err).
			Str("url", urll).
			Msg("Couldn't parse url")
		return rawURL
	}
	return urll
}

func parseURL(rawURL string) (string, error) {
	trimmedRawURL := strings.TrimSpace(rawURL)
	parsedURL, err := url.Parse(trimmedRawURL)
	if err != nil {
		return "", fmt.Errorf("parse.parseURL(): failed url.Parse() on url(%v). error: %w", rawURL, err)
	}

	urlString := parsedURL.String()
	if len(urlString) > 0 && urlString[len(urlString)-1] == '/' {
		urlString = urlString[:len(urlString)-1]
	}

	return urlString, nil
}

func ParseTextWithHTML(rawHTML string) string {
	text, err := parseTextWithHTML(rawHTML)
	if err != nil {
		log.Error().
			Caller().
			Err(err).
			Str("html", rawHTML).
			Msg("Failed parsing text with html")
		return rawHTML
	}
	return text
}

func parseTextWithHTML(rawHTML string) (string, error) {
	var result string = ""

	htmlNode, err := html.ParseFragment(strings.NewReader(rawHTML), nil)
	if err != nil {
		return "", fmt.Errorf("Failed html.ParseFragment on %v: %w", rawHTML, err)
	}

	for _, el := range htmlNode {
		sel := goquery.NewDocumentFromNode(el)
		result += sel.Text()
	}

	return result, nil
}
