package main

import "github.com/ChimeraCoder/anaconda"

var api *anaconda.TwitterApi

func newTwitterAPI() {
	anaconda.SetConsumerKey(*apiConsumerKey)
	anaconda.SetConsumerSecret(*apiConsumerSecret)
	api = anaconda.NewTwitterApi(*userAccessToken, *userAccessTokenSecret)
}
