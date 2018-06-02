package persist

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/guopuke/learngo/crawler/engine"
	"github.com/guopuke/learngo/crawler/model"
	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
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
		},
	}

	// Need Docker elasticsearch service support

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	client, e := elastic.NewClient(elastic.SetSniff(false))
	if e != nil {
		panic(e)
	}

	result, i := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if i != nil {
		panic(i)
	}

	t.Logf("%s", result.Source)

	var actual engine.Item
	err = json.Unmarshal(*result.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}

}
