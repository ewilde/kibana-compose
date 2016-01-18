package api

import (
	"encoding/json"
)

type Export struct {
}

func ExportCreate() *Export {
	return &Export{}
}

func (e *Export) All(api *KibanaApi) (data string, err error) {
	dashboard, err := api.Search("dashboard")
	if err != nil {
		return "", err
	}

	search, err := api.Search("search")
	if err != nil {
		return "", err
	}

	visualization, err := api.Search("visualization")
	if err != nil {
		return "", err
	}

	items := append(append(dashboard, search...), visualization...)
	exports, err := api.Get(items)
	if err != nil {
		return "", err
	}

	byteResult, err := json.MarshalIndent(exports, "", "    ")
	if err != nil {
		return "", err
	}

	return string(byteResult), nil
}
