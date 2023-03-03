package leetcode

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	graphUrl       = "https://leetcode.com/graphql/"
	problemUrl     = "https://leetcode.com/problems/"
	allProblemsUrl = "https://leetcode.com/api/problems/all"
	dataLimit      = 100
)

func sendGraphqlRequest(query string, variables map[string]interface{}) (string, error) {
	jsonData := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}

	request, _ := http.NewRequest("POST", graphUrl, bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Http request failed")
	}

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	return string(data), nil
}
