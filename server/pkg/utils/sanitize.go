package utils

import (
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

var articleHTMLPolicy = newArticleHTMLPolicy()

func newArticleHTMLPolicy() *bluemonday.Policy {
	p := bluemonday.UGCPolicy()
	p.AllowAttrs("class").Globally()
	p.AllowAttrs("target").OnElements("a")
	p.AllowAttrs("rel").OnElements("a")
	p.RequireNoReferrerOnLinks(false)
	p.AllowImages()
	return p
}

func SanitizeArticleHTML(input string) string {
	return articleHTMLPolicy.Sanitize(input)
}

var (
	nonSlugChars = regexp.MustCompile(`[^a-z0-9]+`)
	dashRuns     = regexp.MustCompile(`-+`)
)

func Slugify(input string) string {
	s := strings.ToLower(strings.TrimSpace(input))
	s = nonSlugChars.ReplaceAllString(s, "-")
	s = dashRuns.ReplaceAllString(s, "-")
	return strings.Trim(s, "-")
}
