package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/urfave/cli.v2"
)

func set(user, repo, token string, labels GithubLabels) {
	// Prepare url and client
	url := githubAPIRoot + user + "/" + repo + "/labels" + "?access_token=" + token
	client := &http.Client{}

	// Cache len
	l := len(labels)
	// For each label
	for x, label := range labels {
		fmt.Printf("%d/%d (%s): ", x+1, l, label.Name)
		// Marshal it as json
		b, err := json.Marshal(label)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s\n", err.Error())
			continue
		}

		// Create our request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s\n", err.Error())
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		// Exectue
		resp, err := client.Do(req)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s\n", err.Error())
			continue
		}
		// If status differs from "Created" print error
		if resp.StatusCode != 201 {
			fmt.Fprintf(os.Stderr, "Warning: %s\n", resp.Status)
			continue
		}
		fmt.Println("Ok")
	}
}

// Set reads labels from a json file and set them on the repository passed as parameter
func Set(c *cli.Context) error {
	// Check argc
	if c.NArg() != 1 {
		return cli.Exit(missingRepo, 1)
	}

	// Parse argument to retrieve repository
	user, repo, err := parseRepository(c.Args().Get(0))
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	// If file has not been passed as parameter
	if c.String("labels") == "" {
		return cli.Exit("Missing labels file", 1)
	}

	// Read file
	file, err := ioutil.ReadFile(c.String("labels"))
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	labels := GithubLabels{}
	// Unmarshal labels into slice
	if err := json.Unmarshal(file, &labels); err != nil {
		return cli.Exit(err.Error(), 1)
	}

	set(user, repo, c.String("token"), labels)
	return nil
}
