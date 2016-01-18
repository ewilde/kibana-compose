package api_test

import (
	"github.com/ewilde/kibana-compose/api"
	"github.com/nu7hatch/gouuid"
	"testing"
)

func TestSearch(t *testing.T) {
	api := api.Create(api.DefaultKibanaUri)
	data, err := api.Search("dashboard")

	if err != nil {
		t.Error(err)
	}

	if data == nil {
		t.Error("Expected a non-zero result")
	}
}

func TestGet(t *testing.T) {
	api := api.Create(api.DefaultKibanaUri)
	dashboards, err := api.Search("dashboard")
	visualizations, err := api.Search("visualization")

	items := append(dashboards, visualizations...)
	data, err := api.Get(items)
	if err != nil {
		t.Error(err)
	}

	if data["docs"] == nil {
		t.Error("Expected a non-zero result")
	}
}

func TestDelete(t *testing.T) {
	var data []byte
	var err error
	if data, err = GetTestData(); err != nil {
		t.Error(err)
	}

	kibanaData, err := api.Parse(data)

	var id *uuid.UUID
	if id, err = uuid.NewV4(); err != nil {
		t.Error(err)
	}

	kibanaData.Searches[0]["_id"] = id.String()

	kApi := api.Create(api.DefaultKibanaUri)
	kApi.Create(kibanaData.Searches[0])

	response, err := kApi.Delete(api.KibanaItem{
		Type: kibanaData.Searches[0]["_type"].(string),
		Id:   kibanaData.Searches[0]["_id"].(string),
	})

	if err != nil {
		t.Error(err)
	}

	if response == nil {
		t.Error("Expected non-nil response")
	}

	if !response["found"].(bool) {
		t.Error("Expected item to be found.")
	}
}
