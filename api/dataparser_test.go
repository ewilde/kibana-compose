package api_test

import (
	"github.com/ewilde/kibana-compose/api"
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {

	data, err := GetTestData()
	if err != nil {
		t.Error(err)
	}

	kibanaData, err := api.Parse(data)
	if err != nil {
		t.Error(err)
	}

	if len(kibanaData.Searches) <= 0 {
		t.Errorf("Expected a search count greater than 0, got %v.", len(kibanaData.Searches))
	}

	if len(kibanaData.Visualizations) <= 0 {
		t.Errorf("Expected a visualizations count greater than 0, got %v.", len(kibanaData.Visualizations))
	}

	if len(kibanaData.Dashboards) <= 0 {
		t.Errorf("Expected a dashboard count greater than 0, got %v.", len(kibanaData.Dashboards))
	}
}

func GetTestData() ([]byte, error) {
	data, err := ioutil.ReadFile("../testdata/saved_objects.json")
	return data, err
}
