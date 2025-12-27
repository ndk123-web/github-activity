package github

import (
	// "errors"
	"fmt"
	"io"
	"net/http"

	// "os"

	// "os"
	// "strconv"

	"encoding/json"

	"github.com/ndk123-web/github-activity/internal/config"
	customerror "github.com/ndk123-web/github-activity/internal/custom-error"

	// "github.com/ndk123-web/github-activity/internal/handlers"
	"github.com/ndk123-web/github-activity/internal/models"
)

type GitHubResponse struct {
	Message           string `json:"message"`
	Documentation_url string `json:"documentation_url"`
	Status            string `json:"status"`
}

func FetchGitHubApiData(url string) ([]models.GitResponseObject, error) {

	// var IsUsingGhToken bool = false

	token, err := config.LoadGhToken()

	if err != nil {
		// fmt.Println("Error loading GitHub Token")
	}
	if token != "" {
		// IsUsingGhToken = true
		fmt.Println("- Using GitHub Token for authentication.")
	} else {
		fmt.Println("- No GitHub Token found. Proceeding without authentication.")
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(customerror.Wrap("creating http request failed", err).Error())
		return nil, err
	}

	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(customerror.Wrap("http get failed", err).Error())
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)

		// Try parsing GitHub error JSON for clearer messaging
		var ghErr GitHubResponse
		if len(body) > 0 {
			if err := json.Unmarshal(body, &ghErr); err == nil && ghErr.Message != "" {
				// Special handling for common cases
				if ghErr.Message == "Bad credentials" {
					fmt.Println("🚧 GitHub authentication failed (invalid or expired token).")
					fmt.Println("- Action: removing saved token and retrying without authentication.")
					fmt.Println("- Action taken: token removed and request retried without authentication.")
					fmt.Println("- Note: unauthenticated requests are rate-limited. Set a valid token to avoid this.")
					if err := config.DeleteConfigContents(); err != nil {
						fmt.Println(customerror.Wrap("deleting config contents failed", err).Error())
						return nil, err
					}
					// recursion to retry after deleting config
					return FetchGitHubApiData(url)
				}
				return nil, fmt.Errorf("github api error: status=%d message=%s docs=%s", response.StatusCode, ghErr.Message, ghErr.Documentation_url)
			}
		}
		// Fallback to raw body if not JSON or empty
		return nil, fmt.Errorf("github api error: status=%d body=%s", response.StatusCode, string(body))
	}
	dataa, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(customerror.Wrap("reading response body failed", err).Error())
		return nil, err
	}
	var jsonData []models.GitResponseObject
	if err := json.Unmarshal(dataa, &jsonData); err != nil {
		fmt.Println(customerror.Wrap("json unmarshal failed", err).Error())
		return nil, err
	}
	return jsonData, nil

	// if response.StatusCode != http.StatusOK {
	// 	body, _ := io.ReadAll(response.Body)
	// 	return nil, fmt.Errorf(
	// 		"github api error: status=%d body=%s",
	// 		response.StatusCode,
	// 		string(body),
	// 	)
	// }

	// if err != nil {
	// 	fmt.Println(customerror.Wrap("http get failed", err).Error())
	// 	return nil, err
	// }

	// // close client socket
	// defer response.Body.Close()

	// dataa, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println(customerror.Wrap("reading response body failed", err).Error())
	// 	return nil, err
	// }

	// var jsonData []models.GitResponseObject
	// if err := json.Unmarshal(dataa, &jsonData); err != nil {
	// 	fmt.Println(customerror.Wrap("json unmarshal failed", err).Error())
	// 	return nil, err
	// }

	// return jsonData, nil
}
