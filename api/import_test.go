package api_test

import (
	"github.com/ewilde/kibana-compose/api"
	"github.com/nu7hatch/gouuid"
	"testing"
)

func TestImportItem(t *testing.T) {
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

	importer := api.ImportCreate()
	var response map[string]interface{}
	if response, err = importer.Item(api.Create(api.DefaultKibanaUri), kibanaData.Searches[0]); err != nil {
		t.Error(err)
	}

	if response["_id"] != id.String() {
		t.Error(
			"For", "importer.Item",
			"expected", id.String(),
			"got", response["_id"],
		)
	}
}

func TestImportExistingItem(t *testing.T) {
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

	importer := api.ImportCreate()
	var response map[string]interface{}
	if response, err = importer.Item(api.Create(api.DefaultKibanaUri), kibanaData.Searches[0]); err != nil {
		t.Error(err)
	}

	if response["_id"] != id.String() {
		t.Error(
			"For", "importer.Item",
			"expected", id.String(),
			"got", response["_id"],
		)
	}

	if response, err = importer.Item(api.Create(api.DefaultKibanaUri), kibanaData.Searches[0]); err != nil {
		t.Error(err)
	}

	if response != nil {
		t.Error("For", "importer.Item with existing",
			"expected", "nil response",
			"got", response)
	}
}
