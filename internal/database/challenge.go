package database

import (
	"errors"
	"time"

	c "github.com/ostafen/clover"
)

const ChallengeCollectionName = "weekly_challenge"

type Challenge struct {
	Id   string    `clover:"_id"  json:"_id"`
	Date time.Time `clover:"date" json:"date"`
}

func InitChallengeCollection() error {
	db := GetDb()

	if hasCollection, _ := db.HasCollection(ChallengeCollectionName); !hasCollection {
		err := db.CreateCollection(ChallengeCollectionName)
		if err != nil {
			return err
		}
	}

	return nil
}

func InsertChallenge(challenge *Challenge) error {
	db := GetDb()

	if challenge.Id == "" {
		return errors.New("Cannot insert challenge with non-empty ID")
	}

	challengeDocument := c.NewDocumentOf(challenge)

	if _, err := db.InsertOne(ChallengeCollectionName, challengeDocument); err != nil {
		return err
	}

	return nil
}

func LoadChallenge(id string) (*Challenge, error) {
	db := GetDb()
	var challenge Challenge

	challengeDocument, err := db.Query(ChallengeCollectionName).
		Where(c.Field("id").Eq(id)).
		FindFirst()
	if err != nil {
		return nil, err
	}

	if err = challengeDocument.Unmarshal(challenge); err != nil {
		return nil, err
	}

	return &challenge, nil
}

func LoadLatestChallenge() (*Challenge, error) {
	db := GetDb()
	var challenge Challenge

	challengeDocument, err := db.Query(ChallengeCollectionName).
		Sort(c.SortOption{Field: "date", Direction: -1}).
		FindFirst()
	if err != nil {
		return nil, err
	}

	if err = challengeDocument.Unmarshal(challenge); err != nil {
		return nil, err
	}

	return &challenge, nil
}
