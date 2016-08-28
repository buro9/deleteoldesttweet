package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

var (
	apiConsumerKey        *string
	apiConsumerSecret     *string
	userAccessToken       *string
	userAccessTokenSecret *string
	username              *string

	flags sync.Once
)

func registerFlags() {
	flags.Do(func() {

		apiConsumerKey = flag.String(
			"consumerKey",
			os.Getenv("TWITTER_CONSUMER_KEY"),
			"application consumer key",
		)
		apiConsumerSecret = flag.String(
			"apiConsumerSecret",
			os.Getenv("TWITTER_CONSUMER_SECRET"),
			"application consumer secret",
		)
		userAccessToken = flag.String(
			"userAccessToken",
			os.Getenv("TWITTER_ACCESS_TOKEN"),
			"user access token",
		)
		userAccessTokenSecret = flag.String(
			"userAccessTokenSecret",
			os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
			"user access token secret",
		)
		username = flag.String(
			"username",
			os.Getenv("TWITTER_USERNAME"),
			"twitter username whose oldest tweet we should delete. This must match the access token credentials",
		)
	})
}

func validateFlags() {
	if *apiConsumerKey == "" ||
		*apiConsumerSecret == "" ||
		*userAccessToken == "" ||
		*userAccessTokenSecret == "" ||
		*username == "" {

		fmt.Println("all flags must be provided")
		fmt.Println("visit https://apps.twitter.com to create your app and keys")
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	}
}
