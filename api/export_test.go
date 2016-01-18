package api_test

import (
	"github.com/ewilde/kibana-compose/api"
	"testing"
)

func TestExportAll(t *testing.T) {
	export := api.ExportCreate()
	json, err := export.All(api.Create(api.DefaultKibanaUri))

	if err != nil {
		t.Error(err)
	}

	if len(json) <= 0 {
		t.Error("Expected a non-zero result")
	}
}
