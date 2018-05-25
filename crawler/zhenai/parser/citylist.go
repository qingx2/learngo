package parser

import (
	"regexp"

	"github.com/guopuke/learngo/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	// compile := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)
	compile := regexp.MustCompile(cityListRe)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 10
	for _, m := range matches {
		result.Items = append(
			result.Items, "City"+string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
		// fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		limit--
		if limit == 0 {
			break
		}
	}
	// fmt.Printf("Matches found: %d\n", len(matches))
	return result
}
