package api

import (
	"fmt"
)

type Delete struct {
}

func DeleteCreate() *Delete {
	return &Delete{}
}

func (e *Delete) All(api *KibanaApi) error {
	dashboard, err := api.Search("dashboard")
	if err != nil {
		return err
	}

	search, err := api.Search("search")
	if err != nil {
		return err
	}

	visualization, err := api.Search("visualization")
	if err != nil {
		return err
	}

	items := append(append(dashboard, visualization...), search...)

	for _, value := range items {
		_, err = api.Delete(value)
		fmt.Printf("Deleting %s %s\n", value.Type, value.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
