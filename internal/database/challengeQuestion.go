package database

import (
	c "github.com/ostafen/clover"
)

const ChallengeQuestionCollection = "challenge_questions"

type ChallengeQuestion struct {
	id          string `clover:"_id"`
	challengeId string `clover:"challengeId"`
	titleSlug   string `clover:"titleSlug"`
	title       string `clover:"title"`
	difficulty  int    `clover:"difficulty"`
}

func InsertChallengeQuestion(challengeQuestion *ChallengeQuestion) error {
	db := GetDb()

	challengeQuestionDocument := c.NewDocumentOf(challengeQuestion)

	if _, err := db.InsertOne(ChallengeQuestionCollection, challengeQuestionDocument); err != nil {
		return err
	}

	return nil
}

func LoadByChallengeId(challengeId string) ([]*ChallengeQuestion, error) {
	db := GetDb()

	var challengeQuestions []*ChallengeQuestion

	challengeQuestionDocuments, err := db.Query(ChallengeQuestionCollection).Where(c.Field("challengeId").Eq(challengeId)).FindAll()
	if err != nil {
		return nil, err
	}

	for _, document := range challengeQuestionDocuments {
		var challengeQuestion ChallengeQuestion

		err := document.Unmarshal(challengeQuestion)
		if err != nil {
			return nil, err
		}

		challengeQuestions = append(challengeQuestions, &challengeQuestion)
	}

	return challengeQuestions, nil
}

func LoadByTitleSlug(titleSlug string) (*ChallengeQuestion, error) {
	db := GetDb()

	var challengeQuestion ChallengeQuestion

	challengeQuestionDocument, err := db.Query(ChallengeQuestionCollection).Where(c.Field("titleSlug").Eq(titleSlug)).FindFirst()
	if err != nil {
		return nil, err
	}

	err = challengeQuestionDocument.Unmarshal(challengeQuestion)

	if err != nil {
		return nil, err
	}

	return &challengeQuestion, nil
}
