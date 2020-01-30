package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/patrick-huber-pivotal/ymldiff/formatters"
	"github.com/urfave/cli/v2"

	"github.com/patrick-huber-pivotal/ymldiff/diff"
)

func main() {
	app := &cli.App{
		Name:  "yamldiff",
		Usage: "extracts the differences between two yaml files",
		CustomAppHelpTemplate: `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.HelpName}} [options] <from> <to>
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}
`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Usage:   "the format of the output. Default 'bosh'. (json, bosh, yaml)",
			},
		},
		Action: func(c *cli.Context) error {
			fromFile := c.Args().Get(0)
			toFile := c.Args().Get(1)

			if strings.TrimSpace(fromFile) == "" {
				return fmt.Errorf("missing <from> file argument")
			}

			if strings.TrimSpace(toFile) == "" {
				return fmt.Errorf("missing <to> file argument")
			}

			_, err := os.Stat(fromFile)
			if err != nil {
				return err
			}

			_, err = os.Stat(toFile)
			if err != nil {
				return err
			}

			changeLog, err := diff.NewChangeLogFromFiles(fromFile, toFile)
			if err != nil {
				return err
			}

			format := "bosh"
			if c.IsSet("format") {
				format = c.String("format")
			}
			var formatter formatters.Formatter
			switch format {
			case "bosh":
				formatter = formatters.NewBOSH(changeLog)
				break
			case "json":
				formatter = formatters.NewJSON(changeLog)
				break
			case "yaml":
				formatter = formatters.NewYAML(changeLog)
				break
			default:
				return fmt.Errorf("unrecognized format %s", c.String("format"))
			}

			return formatter.Write(os.Stdout)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
