package github

import (
	// "errors"
	"fmt"
	"io"
	"net/http"
	// "os"
	// "strconv"

	"encoding/json"
	customerror "github.com/ndk123-web/github-activity/internal/custom-error"
	// "github.com/ndk123-web/github-activity/internal/handlers"
	"github.com/ndk123-web/github-activity/internal/models"
)

func FetchGitHubApiData(url string) ([]models.GitResponseObject, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(customerror.Wrap("http get failed", err).Error())
	}

	// close client socket
	defer response.Body.Close()

	dataa, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(customerror.Wrap("reading response body failed", err).Error())
		return nil, err
	}

	var jsonData []models.GitResponseObject
	if err := json.Unmarshal(dataa, &jsonData); err != nil {
		fmt.Println(customerror.Wrap("json unmarshal failed", err).Error())
	}

	return jsonData, nil
}
