package persist

import (
	"GoPractice/crawler/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
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
	}
	id, err := save(expected)
	if err != nil {
		panic(err)
	}
	// TODO: Try to start up elastic search
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual model.Profile
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
