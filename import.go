package main

import "gopkg.in/urfave/cli.v2"

// Import imports labels from a repository to another
func Import(c *cli.Context) error {
	// Check argc
	if c.NArg() != 2 {
		return cli.Exit(missingRepo, 1)
	}

	// Parse argument to retrieve repository
	userFrom, repoFrom, err := parseRepository(c.Args().Get(0))
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	labels, err := get(userFrom, repoFrom, c.String("token"))
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	// Parse argument to retrieve repository
	userTo, repoTo, err := parseRepository(c.Args().Get(1))
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	set(userTo, repoTo, c.String("token"), labels)

	return nil
}
