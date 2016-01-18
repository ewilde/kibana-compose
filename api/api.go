package api

import (
	"github.com/bndr/gopencils"
)

var DefaultKibanaUri = "http://localhost/elasticsearch/.kibana"

type KibanaApi struct {
	Uri string
}

type KibanaItem struct {
	Id   string `json:"_id"`
	Type string `json:"_type"`
}

func Create(uri string) *KibanaApi {
	return &KibanaApi{Uri: uri}
}

func (k *KibanaApi) Get(items []KibanaItem) (map[string]interface{}, error) {
	api := gopencils.Api(k.Uri)
	var response interface{}

	postData := make(map[string]interface{})
	postData["docs"] = items
	//postJson, _ := json.Marshal(postData);
	_, err := api.Res("/_mget", &response).Post(postData)
	if err != nil {
		return nil, err
	}

	return response.(map[string]interface{}), nil
}

func (k *KibanaApi) Search(itemType string) ([]KibanaItem, error) {
	var response interface{}
	api := gopencils.Api(k.Uri)
	querystring := map[string]string{"size": "100"}
	if _, err := api.Res(itemType+"/_search", &response).SetQuery(querystring).Get(); err != nil {
		return nil, err
	}

	hits := response.(map[string]interface{})["hits"].(map[string]interface{})["hits"]
	items := hits.([]interface{})

	result := make([]KibanaItem, 0)
	for _, x := range items {
		m := x.(map[string]interface{})
		result = append(result, KibanaItem{m["_id"].(string), m["_type"].(string)})
	}

	return result, nil
}

func (k *KibanaApi) Create(item map[string]interface{}) (map[string]interface{}, error) {
	querystring := map[string]string{"op_type": "create"}

	api := gopencils.Api(k.Uri)

	var response interface{}
	var err error
	if api, err = api.Res(item["_type"], &response).Id(item["_id"]).SetQuery(querystring).Post(item["_source"]); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}
	return response.(map[string]interface{}), nil
}

func (k *KibanaApi) Delete(item KibanaItem) (map[string]interface{}, error) {
	api := gopencils.Api(k.Uri)
	var response interface{}
	var err error
	if api, err = api.Res(item.Type, &response).Id(item.Id).Delete(); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}
	return response.(map[string]interface{}), nil
}
