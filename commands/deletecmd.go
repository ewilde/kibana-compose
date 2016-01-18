package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ewilde/kibana-compose/api"
)

func CmdDelete(c *cli.Context) {
	delete := api.DeleteCreate()
	err := delete.All(api.Create(c.String("uri")))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Delete successful...")
}
