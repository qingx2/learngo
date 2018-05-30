package parser

import (
	"regexp"

	"github.com/guopuke/learngo/crawler/engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	rs := engine.ParseResult{}

	matches := profileRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		name := string(m[2])
		url := string(m[1])
		// rs.Items = append(rs.Items, "User "+name)

		rs.Requests = append(rs.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, url, name)
			},
		})
	}

	cityMatches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range cityMatches {
		rs.Requests = append(rs.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return rs
}
