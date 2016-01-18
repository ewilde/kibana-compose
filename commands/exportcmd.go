package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ewilde/kibana-compose/api"
	"os"
)

func CmdExport(c *cli.Context) {
	export := api.ExportCreate()
	json, err := export.All(api.Create(c.String("uri")))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var file string
	if file = c.String("file"); file == "" {
		fmt.Println(json)
		return
	}

	os.Remove(file)
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	logFile.WriteString(json)
	fmt.Println("Export successful...")
}
