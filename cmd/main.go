package main

import (
	"errors"
	"fmt"
	"strconv"

	// "io"
	// "net/http"
	"os"
	// "strconv"

	// "encoding/json"
	customerror "github.com/ndk123-web/github-activity/internal/custom-error"
	"github.com/ndk123-web/github-activity/internal/github"
	"github.com/ndk123-web/github-activity/internal/handlers"
	"github.com/ndk123-web/github-activity/internal/models"
	// "google.golang.org/grpc/benchmark/flags"
	// "github.com/ndk123-web/github-activity/internal/handlers"
	// "github.com/ndk123-web/github-activity/internal/models"
)

func main() {

	// these are rules
	rules := models.Rules()
	scopes := models.Scopes()

	// scopeUser := false
	// scopeRepo := false

	// command -> gh-activity user username command options
	currentScope := os.Args[1]

	// check scope is valid or not
	if !github.IsValidScope(currentScope, scopes) {
		fmt.Println(customerror.Wrap(("Invalid Scope"), errors.New("Invalid Scope Error")))
		return
	}

	switch currentScope {
	case "user":
		{
			username := os.Args[2]
			if username == "" {
				fmt.Println(customerror.Wrap("Username is required", errors.New("Username Missing Error")))
				return
			}

			// get the user url
			url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

			jsonData, err := github.FetchGitHubApiData(url)
			if err != nil {
				fmt.Println(customerror.Wrap("Json Issue", err))
				return
			}

			// process commmand
			currentCommand := os.Args[3]
			if !github.IsValidCommand(currentCommand, rules, currentScope) {
				fmt.Println(customerror.Wrap(fmt.Sprintf("Invalid Scope For %s", currentScope), fmt.Errorf(fmt.Sprintf("Invalid Command %s", currentCommand))))
				return
			}

			flags := make(map[string]string)
			// process flags if any
			if len(os.Args) > 4 {

				// verify that each flag have data or not
				if (len(os.Args)-4)%2 != 0 {
					fmt.Println(customerror.Wrap("Some Flag Data Missing", errors.New("Flag Data Missing Error")))
					return
				}

				// then process each flag
				for i := 4; i < len(os.Args); i++ {
					currentFlag := os.Args[4]
					if !github.IsValidFlag(currentFlag, rules[currentScope][currentCommand]) {
						fmt.Println(customerror.Wrap(fmt.Sprintf("Invalid Flag For %s", currentCommand), errors.New(fmt.Sprintf("Invalid Flag %s", currentFlag))))
					}
					flags[currentFlag] = os.Args[i+1]
					i++
				}
			}

			switch currentCommand {
			case "pushes":
				{
					push_handler := handlers.NewGitHandler(url)
					var limit int64 = 0 // default value

					if l, ok := flags["--limit"]; ok {
						limit, err = strconv.ParseInt(l, 10, 64)
						if err != nil {
							fmt.Println(customerror.Wrap("Limit Flag Parsing Issue", err))
							return
						}
					}

					// set default limit if limit is zero
					if limit == 0 {
						limit = 2
					}

					push_handler.GetAllResponseObjects(jsonData)
					push_handler.GetResponseRepoWise(limit, jsonData)
				}
			default:
				{
					fmt.Println(customerror.Wrap("Command Not Implemented", errors.New("Command Not Implemented")))
				}
			}
		}
		// case "repo": {
		// 	scopeRepo = true
		//
	default:
		{
			fmt.Println(customerror.Wrap("Scope Not Implemented", errors.New("Scope Not Implemented")))
		}
	}

	// // process user scope
	// if scopeUser {

	// }

}
