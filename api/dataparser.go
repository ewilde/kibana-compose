package api

import (
	"encoding/json"
	"fmt"
)

func Parse(data []byte) (*KibanaData, error) {
	var f interface{}
	result := NewKibanaData()
	err := json.Unmarshal(data, &f)
	if err != nil {
		return nil, err
	}

	doc := f.(map[string]interface{})
	items := doc["docs"].([]interface{})

	for _, x := range items {
		m := x.(map[string]interface{})
		switch m["_type"] {
		case "dashboard":
			result.Dashboards = append(result.Dashboards, m)
		case "search":
			result.Searches = append(result.Searches, m)
		case "visualization":
			result.Visualizations = append(result.Visualizations, m)
		default:
			fmt.Println("unknown")
		}
	}

	return result, nil
}
