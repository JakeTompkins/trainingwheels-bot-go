package leetcode

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	d "github.com/jaketompkins/trainingwheels-bot-go/internal/database"
	j "github.com/valyala/fastjson"
)

const (
	graphUrl       = "https://leetcode.com/graphql/"
	problemUrl     = "https://leetcode.com/problems/"
	allProblemsUrl = "https://leetcode.com/api/problems/all"
	dataLimit      = 100
)

type Question = d.Question

func sendGraphqlRequest(
	query string,
	variables map[string]string,
) (*j.Value, error) {
	jsonData := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	request, _ := http.NewRequest("POST", graphUrl, bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Http request failed")
	}

	defer response.Body.Close()
	data, _ := io.ReadAll(response.Body)

	responseBody, err := j.Parse(string(data))
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func sendRestRequest(url string) (*j.Value, error) {
	request, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Http request failed")
	}

	defer response.Body.Close()
	data, _ := io.ReadAll(response.Body)

	responseBody, err := j.Parse(string(data))
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func unmarshalAllQuestionsResponse(response *j.Value) []*Question {
	var questions []*Question
	rawQuestions := response.GetArray("stat_status_pairs")

	for _, rawQuestion := range rawQuestions {
		var question Question

		question.Id = string(rawQuestion.GetStringBytes("stat", "question_id"))
		question.Title = string(rawQuestion.GetStringBytes("stat", "question_title"))
		question.TitleSlug = string(rawQuestion.GetStringBytes("stat", "question_title_slug"))
		question.Difficulty = rawQuestion.GetInt("difficulty", "level")

		questions = append(questions, &question)

	}

	return questions
}

func buildUserRankString(leetcodeId string, response *j.Value) (string, error) {
	ranking := string(response.GetStringBytes("data", "matchedUser", "profile", "ranking"))
	contributionPoints := response.GetInt("data", "matchedUser", "contributions", "points")
	submissions := response.GetArray("data", "matchedUser", "submitStats", "acSubmissionNum")

	var easySubmissions int
	var mediumSubmissions int
	var hardSubmissions int

	for _, submission := range submissions {
		difficulty := string(submission.GetStringBytes("difficulty"))

		if difficulty == "Easy" {
			easySubmissions = submission.GetInt("count")
		}
		if difficulty == "Medium" {
			mediumSubmissions = submission.GetInt("count")
		}
		if difficulty == "Hard" {
			hardSubmissions = submission.GetInt("count")
		}
	}

	if ranking == "" {
		return "", errors.New("could not get all user data")
	}

	return fmt.Sprintf(`
	Name:				 %v
	Ranking:			 %v
	Contribution Points: %v
	Easy Challenges:	 %v
	Medium Challenges:	 %v
	Hard Challenges:	 %v
	`, leetcodeId, ranking, contributionPoints, easySubmissions, mediumSubmissions, hardSubmissions), nil
}

func LoadAllQuestions() ([]*Question, error) {
	response, err := sendRestRequest(allProblemsUrl)
	if err != nil {
		return nil, err
	}

	return unmarshalAllQuestionsResponse(response), nil
}

func LoadUserRank(leetcodeId string) (string, error) {
	query, variables := BuildUserRankQuery(leetcodeId)
	response, err := sendGraphqlRequest(query, variables)
	if err != nil {
		return "", err
	}

	return buildUserRankString(leetcodeId, response)
}
