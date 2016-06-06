package main

import (
	"errors"
	"os"
	"strings"

	"gopkg.in/urfave/cli.v2"
)

const missingRepo = "Error: missing repository."
const githubAPIRoot = "https://api.github.com/repos/"

type GithubLabel struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type GithubLabels []GithubLabel

func parseRepository(repository string) (string, string, error) {
	res := strings.Split(repository, "/")

	if len(res) != 2 {
		return "", "", errors.New("Invalid repository name")
	}

	return res[0], res[1], nil
}

func main() {
	app := cli.NewApp()
	app.Name = "github-labels"
	app.Usage = "Manage github labels easily"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "token",
			Aliases: []string{"t"},
			Usage:   "Github Token (Mandatory)",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:   "list",
			Usage:  "Print labels from a repository",
			Action: List,
		},
		{
			Name:  "set",
			Usage: "Set repository labels from a file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "labels",
					Aliases: []string{"l"},
					Usage:   "Labels json file",
				},
			},
			Action: Set,
		},
		{
			Name:   "import",
			Usage:  "Import labels from a repository to another",
			Action: Import,
		},
	}

	app.Run(os.Args)
}
