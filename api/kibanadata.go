package api

type KibanaData struct {
	Searches       []map[string]interface{}
	Visualizations []map[string]interface{}
	Dashboards     []map[string]interface{}
}

func NewKibanaData() *KibanaData {
	return &KibanaData{
		make([]map[string]interface{}, 0),
		make([]map[string]interface{}, 0),
		make([]map[string]interface{}, 0)}
}
