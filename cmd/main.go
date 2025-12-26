package main

import (
	"errors"
	"fmt"
	"strconv"

	"os"

	"github.com/ndk123-web/github-activity/internal/config"
	customerror "github.com/ndk123-web/github-activity/internal/custom-error"
	"github.com/ndk123-web/github-activity/internal/github"
	"github.com/ndk123-web/github-activity/internal/handlers"
	"github.com/ndk123-web/github-activity/internal/models"
	// "golang.org/x/text/cases"
)

func main() {

	// these are rules
	rules := models.Rules()
	scopes := models.Scopes()

	// scopeUser := false
	// scopeRepo := false

	if len(os.Args) < 4 && os.Args[1] != "set" && os.Args[1] != "get" {
		fmt.Println(customerror.Wrap("Insufficient Arguments", errors.New("Insufficient Arguments Error")))
		return
	}

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
			url := fmt.Sprintf("https://api.github.com/users/%s/events?per_page=60", username)

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

				// verify that each flag has an accompanying value
				if (len(os.Args)-4)%2 != 0 {
					fmt.Println(customerror.Wrap("Some Flag Data Missing", errors.New("Flag Data Missing Error")))
					return
				}

				// process flags as pairs: flag then value
				for i := 4; i < len(os.Args); i += 2 {
					currentFlag := os.Args[i]
					if !github.IsValidFlag(currentFlag, rules[currentScope][currentCommand]) {
						fmt.Println(customerror.Wrap(fmt.Sprintf("Invalid Flag For %s", currentCommand), fmt.Errorf(fmt.Sprintf("Invalid Flag %s", currentFlag))))
						return
					}
					flags[currentFlag] = os.Args[i+1]
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
			case "pulls":
				{
					pull_handler := handlers.NewPullHandler(url)

					// process limit flag
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

					// process state flag
					// state := "all" // default value

					// its mandatory to provide state flag now
					if _, ok := flags["--state"]; !ok {
						fmt.Println(customerror.Wrap("State Flag Missing", errors.New("State Flag Missing Error")))
						return
					}

					state := flags["--state"]

					if state != "open" && state != "closed" && state != "merged" {
						fmt.Println(customerror.Wrap("Invalid State Value", errors.New("State Value Should be one of open, closed, merged")))
						return
					}

					err := pull_handler.GetAllPullRequests(jsonData)
					if err != nil {
						fmt.Println(customerror.Wrap("Pull Handler Issue", err))
						return
					}

					err = pull_handler.GetPullRequestRepoWise(limit, state, jsonData)
					if err != nil {
						fmt.Println(customerror.Wrap("Pull Handler Issue", err))
						return
					}
				}
			case "issues":
				{
					// check the flags
					// flags that are possible --limit and --state
					// where limit is default 2 and state is mandatory
					if _, ok := flags["--state"]; !ok {
						fmt.Println(customerror.Wrap("State Flag Missing", errors.New("State Flag Missing Error")))
						return
					}

					// state := flags["--state"]

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

					issue_handler := handlers.NewIssueEventHandler(url)
					if err := issue_handler.GetAllIssueEvents(jsonData); err != nil {
						fmt.Println(customerror.Wrap("Issue Handler Issue", err))
						return
					}
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
	case "set":
		{
			// get the command which will be token
			currentCommand := os.Args[2]
			if currentCommand != "token" {
				fmt.Println(customerror.Wrap("Invalid Command For set", errors.New("Invalid Command Error")))
				return
			}

			// data will be i 	n os.Args[3]
			if len(os.Args) < 4 {
				fmt.Println(customerror.Wrap("Token Value Missing", errors.New("Token Value Missing Error")))
				return
			}
			tokenValue := os.Args[3]

			config.SetGhToken(tokenValue)
		}

	case "get":
		{
			// get the command which will be token
			currentCommand := os.Args[2]
			if currentCommand != "token" {
				fmt.Println(customerror.Wrap("Invalid Command For get", errors.New("Invalid Command Error")))
				return
			}
			token, err := config.LoadGhToken()
			if err != nil {
				fmt.Println(customerror.Wrap("Loading Token Failed", err))
				return
			}
			if token == "" {
				fmt.Println("- GitHub Token is not set.")
			}
			fmt.Println("GitHub Token:", token)
		}
	default:
		{
			fmt.Println(customerror.Wrap("Scope Not Implemented", errors.New("Scope Not Implemented")))
		}
	}

	// // process user scope
	// if scopeUser {

	// }

}
