package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ewilde/kibana-compose/api"
	"io/ioutil"
)

func CmdImport(c *cli.Context) {
	var file string
	var err error
	if file = c.String("file"); file == "" {
		fmt.Println("file flag missing")
		cli.ShowAppHelp(c)
		return
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	importer := api.ImportCreate()
	kibanaData, err := api.Parse(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = importer.All(api.Create(c.String("uri")), kibanaData)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Import successful...")
}
