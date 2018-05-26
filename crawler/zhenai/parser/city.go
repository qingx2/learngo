package parser

import (
	"regexp"

	"github.com/guopuke/learngo/crawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	rs := engine.ParseResult{}

	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		rs.Items = append(rs.Items, "User "+string(m[2]))

		rs.Requests = append(rs.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
			// 	func(c []byte) engine.ParseResult {
			// 	return ParseProfile(c, name)
			// },
		})
	}

	return rs
}
