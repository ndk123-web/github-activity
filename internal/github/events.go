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

func FetchGitHubApiData(url string) ([]models.GitResponseObject, error) {

	var IsUsingGhToken bool = false

	token, err := config.LoadGhToken()

	if err != nil {
		fmt.Println("- Using GitHub Token for authentication.")
	}
	if token != "" {
		IsUsingGhToken = true
		fmt.Println("- Using GitHub Token for authentication.")
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
		if IsUsingGhToken {
			fmt.Println("- GitHub Token might be invalid or has insufficient scopes.")
		}
		return nil, fmt.Errorf(
			"github api error: status=%d body=%s",
			response.StatusCode,
			string(body),
		)
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
