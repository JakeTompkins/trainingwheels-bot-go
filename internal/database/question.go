package database

import (
	"errors"

	c "github.com/ostafen/clover"
)

const QuestionCollectionName = "questions"

type Question struct {
	Id         string `clover:"_id"        json:"_id"`
	Title      string `clover:"title"      json:"title"`
	TitleSlug  string `clover:"titleSlug"  json:"titleSlug"`
	Difficulty int    `clover:"difficulty" json:"difficulty"`
}

func InitQuestionCollection() error {
	db := GetDb()

	if hasCollection, _ := db.HasCollection(QuestionCollectionName); !hasCollection {
		err := db.CreateCollection(QuestionCollectionName)
		if err != nil {
			return err
		}
	}

	return nil
}

func InsertQuestion(inputQuestion *Question) error {
	db := GetDb()

	existingQuestion, err := FindQuestionByTitleSlug(inputQuestion.TitleSlug)

	if existingQuestion != nil {
		return errors.New("Question already exists")
	}

	if err != nil {
		return err
	}
	questionDocument := c.NewDocumentOf(inputQuestion)

	_, err = db.InsertOne(QuestionCollectionName, questionDocument)

	if err != nil {
		return err
	}

	question := new(Question)
	err = questionDocument.Unmarshal(question)

	if err != nil {
		return err
	}

	return nil
}

func InsertQuestions(inputQuestions []*Question) error {
	db := GetDb()
	var questions []*c.Document

	for _, question := range inputQuestions {
		questionDocument := c.NewDocumentOf(question)
		questions = append(questions, questionDocument)
	}

	err := db.Insert(QuestionCollectionName, questions...)
	if err != nil {
		return err
	}

	return nil
}

func FindQuestionByTitleSlug(titleSlug string) (*Question, error) {
	db := GetDb()

	question := new(Question)

	questionDocument, err := db.Query(QuestionCollectionName).
		Where((*c.Criteria)(c.Field("titleSlug").Eq(titleSlug))).
		FindFirst()

	if questionDocument == nil || err != nil {
		return nil, err
	}

	err = questionDocument.Unmarshal(question)

	if err != nil {
		return nil, err
	}

	return question, nil
}
