package main

import (
	"errors"
	"fmt"

	// "io"
	// "net/http"
	"os"
	"strconv"

	// "encoding/json"
	customerror "github.com/ndk123-web/github-activity/internal/custom-error"
	"github.com/ndk123-web/github-activity/internal/github"
	"github.com/ndk123-web/github-activity/internal/handlers"
	// "github.com/ndk123-web/github-activity/internal/models"
)

func main() {

	var args []string
	var mapp = make(map[string]string)

	args = os.Args

	// logic is always
	// 0 idx = gh-ndk , 1 idx = --command , 2 idx data , 3 idx --command , 4 data
	// we can see that 2 , 4, 6 are going to be data
	// and 1, 3, 5, are going to be commands
	// it means we can say that , odd idxs ones are command , even idxs onces are data of that previous command

	var cmd string
	var data string
	for idx, str := range args {
		if idx > 0 && idx%2 == 1 {
			// these are commands
			fmt.Printf("Idx: %v Command: %s\n", idx, str)
			cmd = str
		} else if idx > 0 {
			// these are data of previous commands
			fmt.Printf("Idx: %v Data: %s\n", idx, str)
			data = str
			mapp[cmd] = data
		}
	}

	// first fetch username
	key := "--u"
	username, ok := mapp[key]
	if !ok {
		fmt.Println(customerror.Wrap("Please Provide Valid Username", errors.New("Username Not Exist")).Error())
	}

	// get the user url
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	jsonData, err := github.FetchGitHubApiData(url)
	if err != nil {
		fmt.Println(customerror.Wrap("Json Issue", err))
	}

	for key, value := range mapp {
		switch key {
		case "--push":
			{
				// logic for push limit
				limit := value
				var intLimit int64
				var err error
				if limit != "" {
					intLimit, err = strconv.ParseInt(limit, 10, 64)
					if err != nil {
						fmt.Println(customerror.Wrap("Please Enter Right Limit", errors.New("Limit is Not Integer")).Error())
					}
				} else {
					intLimit = 2
				}

				// create handler
				git_handler := handlers.NewGitHandler(url)
				git_handler.GetAllResponseObjects()
				git_handler.GetResponseRepoWise(intLimit, jsonData)
			}
		}
	}
}
