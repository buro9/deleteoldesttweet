package main

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

var user anaconda.User

func getUser() error {
	users, err := api.GetUsersLookup(*username, url.Values{})
	if err != nil {
		return err
	}

	if len(users) != 1 {
		return fmt.Errorf("Expected a single user, got %d", len(users))
	}

	user = users[0]

	return nil
}
