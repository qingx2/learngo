package persist

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/guopuke/learngo/crawler/model"
	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
		Name:       "冰之泪",
		Age:        47,
		Height:     170,
		Weight:     189,
		Marriage:   "离异",
		Income:     "8001-12000元",
		Education:  "大专",
		Occupation: "其他职业",
		Gender:     "男",
		House:      "有",
		Car:        "无",
		Hukou:      "无",
		Xingzuo:    "无",
	}

	// Need Docker elasticsearch service support

	id, err := save(profile)
	if err != nil {
		panic(err)
	}

	client, e := elastic.NewClient(elastic.SetSniff(false))
	if e != nil {
		panic(e)
	}

	result, i := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if i != nil {
		panic(i)
	}

	t.Logf("%s", result.Source)

	var actual model.Profile
	unmarshal := json.Unmarshal(*result.Source, &actual)
	if unmarshal != nil {
		panic(unmarshal)
	}

}
