package api

type Import struct {
}

func ImportCreate() *Import {
	return &Import{}
}

func (e *Import) Item(api *KibanaApi, item map[string]interface{}) (response map[string]interface{}, err error) {
	return api.Create(item)
}

func (e *Import) All(api *KibanaApi, kibanaData *KibanaData) error {
	var err error
	for _, value := range kibanaData.Searches {
		_, err = api.Create(value)
		if err != nil {
			return err
		}
	}

	for _, value := range kibanaData.Visualizations {
		_, err = api.Create(value)
		if err != nil {
			return err
		}
	}

	for _, value := range kibanaData.Dashboards {
		_, err = api.Create(value)
		if err != nil {
			return err
		}
	}

	return nil
}
