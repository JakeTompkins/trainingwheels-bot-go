package database

import (
	"errors"

	c "github.com/ostafen/clover"
)

const QuestionCollectionName = "questions"

type Question struct {
	id         string `clover:"_id"`
	title      string `clover:"title"`
	titleSlug  string `clover:"titleSlug"`
	difficulty int    `clover:"difficulty"`
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

func InsertQuestion(inputQuestion *Question) (*Question, error) {
	db := GetDb()

	existingQuestion, err := FindQuestionByTitleSlug(inputQuestion.titleSlug)

	if existingQuestion != nil {
		return nil, errors.New("Question already exists")
	}

	if err != nil {
		return nil, err
	}

	questionDocument := c.NewDocumentOf(inputQuestion)

	_, err = db.InsertOne(QuestionCollectionName, questionDocument)

	if err != nil {
		return nil, err
	}

	question := new(Question)
	err = questionDocument.Unmarshal(question)

	if err != nil {
		return nil, err
	}

	return question, nil
}

func InsertQuestions(inputQuestions []*Question) ([]*Question, error) {
	db := GetDb()
}

func FindQuestionByTitleSlug(titleSlug string) (*Question, error) {
	db := GetDb()

	question := new(Question)

	questionDocument, err := db.Query(QuestionCollectionName).Where((*c.Criteria)(c.Field("titleSlug").Eq(titleSlug))).FindFirst()

	if questionDocument == nil || err != nil {
		return nil, err
	}

	questionDocument.Unmarshal(question)

	return question, nil
}
