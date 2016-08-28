package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os/exec"
	"sync"
)

const oldestTweetCMD string = `curl 
	"%s"`

var (
	oldestTweetURL     *url.URL
	oldestTweetURLOnce sync.Once
)

func getOldestTweetID() (tweetID int64, err error) {
	oldestTweetURLOnce.Do(func() {
		oldestTweetURL, _ = url.Parse(`https://discover.twitter.com/first-tweet`)
		qs := oldestTweetURL.Query()
		qs.Set("username", "buro9")
		oldestTweetURL.RawQuery = qs.Encode()
	})

	cmdName := `curl`
	args := []string{
		`-X`, `GET`,
		`-H`, `Host: discover.twitter.com`,
		`-H`, `User-Agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:48.0) Gecko/20100101 Firefox/48.0`,
		`-H`, `Accept: */*`,
		`-H`, `Accept-Language: en-GB`,
		`-H`, `DNT: 1`,
		`-H`, `X-Requested-With: XMLHttpRequest`,
		`-H`, `Referer: https://discover.twitter.com/first-tweet`,
		`-H`, `Connection: keep-alive`,
		fmt.Sprintf(`%s`, oldestTweetURL.String()),
	}

	cmd := exec.Command(cmdName, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if e := cmd.Run(); e != nil {
		err = fmt.Errorf(fmt.Sprint(e) + ": " + stderr.String())
		return
	}

	type OldestTweet struct {
		Username string `json:"username"`
		StatusID int64  `json:"status_id"`
		HTML     string `json:"html"`
	}

	var oldestTweet OldestTweet
	err = json.Unmarshal(out.Bytes(), &oldestTweet)
	if err != nil {
		return
	}

	tweetID = oldestTweet.StatusID

	return
}
