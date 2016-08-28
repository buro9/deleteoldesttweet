package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	registerFlags()
	flag.Parse()
	validateFlags()

	newTwitterAPI()

	err := getUser()
	if err != nil {
		fmt.Println("getUser error: " + err.Error())
		os.Exit(1)
	}

	var id int64
	id, err = getOldestTweetID()
	if err != nil {
		fmt.Println("getOldestTweetID error: " + err.Error())
		os.Exit(1)
	}

	for id > 0 {
		fmt.Printf("Tweet %d will be deleted\n", id)

		_, err = api.DeleteTweet(id, true)
		if err != nil && !strings.Contains(err.Error(), "status 404") {
			fmt.Println("getOldestTweetID error: " + err.Error())
			os.Exit(1)
		}

		fmt.Println("Oldest Tweet Deleted")

		id, err = getOldestTweetID()
		if err != nil {
			fmt.Println("getOldestTweetID error: " + err.Error())
			os.Exit(1)
		}

		time.Sleep(3 * time.Second)
	}

	fmt.Println("All tweets deleted")
}
