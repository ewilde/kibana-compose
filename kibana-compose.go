package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ewilde/kibana-compose/api"
	"github.com/ewilde/kibana-compose/commands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "kibana-compser"
	app.Usage = "functions to import and export kibana objects"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:   "export",
			Usage:  "exports kibana objects to a file.",
			Action: commands.CmdExport,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file",
					Usage: "File to export json to.",
					Value: "",
				},
				cli.StringFlag{
					Name:  "uri",
					Usage: fmt.Sprintf("Kibana api uri (defaults to %s).", api.DefaultKibanaUri),
					Value: api.DefaultKibanaUri,
				},
			},
		},
		{
			Name:   "import",
			Usage:  "imports kibana objects from a file.",
			Action: commands.CmdImport,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file",
					Usage: "File to import from (required).",
					Value: "",
				},
				cli.StringFlag{
					Name:  "uri",
					Usage: fmt.Sprintf("Kibana api uri (defaults to %s).", api.DefaultKibanaUri),
					Value: api.DefaultKibanaUri,
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "deletes all kibana objects.",
			Action: commands.CmdDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "uri",
					Usage: fmt.Sprintf("Kibana api uri (defaults to %s).", api.DefaultKibanaUri),
					Value: api.DefaultKibanaUri,
				},
			},
		},
	}
	app.Run(os.Args)
}
