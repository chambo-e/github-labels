package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"

	"gopkg.in/urfave/cli.v2"
)

func get(user, repo, token string) (GithubLabels, error) {
	// Build url with the user and repository
	url := githubAPIRoot + user + "/" + repo + "/labels" + "?access_token=" + token

	// Execute the request
	resp, err := http.Get(url)
	if err != nil {
		return GithubLabels{}, err
	}

	// If the request failed return the error status
	if resp.StatusCode != 200 {
		return GithubLabels{}, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GithubLabels{}, err
	}

	// Unmarshal response body to slice
	labels := GithubLabels{}
	if err := json.Unmarshal(body, &labels); err != nil {
		return GithubLabels{}, err
	}

	return labels, nil
}

// List print the list of labels from a repository
func List(c *cli.Context) error {
	// Check argc
	if c.NArg() != 1 {
		return cli.Exit(missingRepo, 1)
	}

	// Parse argument to retrieve repository
	user, repo, err := parseRepository(c.Args().Get(0))
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	labels, err := get(user, repo, c.String("token"))
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	// Init our tabwriter
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	// Print labels
	for _, label := range labels {
		fmt.Fprintf(w, "label: %s\tcolor: #%s\n", label.Name, label.Color)
	}

	// Flush our writer
	w.Flush()

	return nil
}
