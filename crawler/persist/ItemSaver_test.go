package persist

import (
	"GoPractice/crawler/engine"
	"GoPractice/crawler/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "zhenai",
		Type: "zhenai",
		Id:   "1089",
		PayLoad: model.Profile{
			Name:       "厌与深情记得笑i",
			Gender:     "女",
			Age:        10,
			Height:     1,
			Weight:     116,
			Income:     "2001-3000元",
			Marriage:   "未婚",
			Education:  "硕士",
			Occupation: "销售",
			Hokou:      "东莞市",
			Xinzuo:     "天秤座",
			House:      "有房",
			Car:        "有豪车",
		},
	}
	// save expected
	//TODO: Try to start up elastic search
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	const index = "dating_test"
	err = Save(client, expected, index)

	if err != nil {
		panic(err)
	}

	// get item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, err := model.FromJsonObj(actual.PayLoad)
	if err != nil {
		panic(err)
	}
	actual.PayLoad = actualProfile
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
