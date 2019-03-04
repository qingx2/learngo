package parser

import (
	"regexp"

	"github.com/guopuke/learngo/crawler/engine"
	"github.com/guopuke/learngo/crawler_distributed/config"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	rs := engine.ParseResult{}

	matches := profileRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		// rs.Items = append(rs.Items, "User "+name)

		rs.Requests = append(rs.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	cityMatches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range cityMatches {
		rs.Requests = append(rs.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	return rs
}
