package database

import (
	c "github.com/ostafen/clover"
)

const ChallengeQuestionCollectionName = "challenge_questions"

type ChallengeQuestion struct {
	Id          string `clover:"_id"         json:"_id"`
	ChallengeId string `clover:"challengeId" json:"challengeId"`
	TitleSlug   string `clover:"titleSlug"   json:"titleSlug"`
	Title       string `clover:"title"       json:"title"`
	Difficulty  int    `clover:"difficulty"  json:"difficulty"`
}

func InitChallengeQuestionCollection() error {
	db := GetDb()

	if hasCollection, _ := db.HasCollection(ChallengeQuestionCollectionName); !hasCollection {
		err := db.CreateCollection(ChallengeQuestionCollectionName)
		if err != nil {
			return err
		}
	}

	return nil
}

func InsertChallengeQuestion(challengeQuestion *ChallengeQuestion) error {
	db := GetDb()

	challengeQuestionDocument := c.NewDocumentOf(challengeQuestion)

	if _, err := db.InsertOne(ChallengeQuestionCollectionName, challengeQuestionDocument); err != nil {
		return err
	}

	return nil
}

func LoadByChallengeId(challengeId string) ([]*ChallengeQuestion, error) {
	db := GetDb()

	var challengeQuestions []*ChallengeQuestion

	challengeQuestionDocuments, err := db.Query(ChallengeQuestionCollectionName).
		Where(c.Field("challengeId").Eq(challengeId)).
		FindAll()
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

	challengeQuestionDocument, err := db.Query(ChallengeQuestionCollectionName).
		Where(c.Field("titleSlug").Eq(titleSlug)).
		FindFirst()
	if err != nil {
		return nil, err
	}

	err = challengeQuestionDocument.Unmarshal(challengeQuestion)

	if err != nil {
		return nil, err
	}

	return &challengeQuestion, nil
}
