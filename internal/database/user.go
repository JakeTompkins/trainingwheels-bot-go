package database

import (
	c "github.com/ostafen/clover"
)

const UserCollectionName = "users"

type User struct {
	id         string `clover: "id"`
	discordId  string `clover: "discordId"`
	leetcodeId string `clover: "leetcodeId"`
}

func InitUserCollection() {
	db := GetDb()

	if hasCollection, _ := db.HasCollection("users"); !hasCollection {
		err := db.CreateCollection(UserCollectionName)
		if err != nil {
			panic(err)
		}
	}
}

func InsertUser(discordId string, leetcodeId string) bool {
	db := GetDb()

	userDocument := c.NewDocument()
	userDocument.Set("discordId", discordId)
	userDocument.Set("leetcodeId", leetcodeId)

	_, err := db.InsertOne(UserCollectionName, userDocument)
	if err != nil {
		return false
	}

	return true
}

func FindUserByDiscordId(discordId string) (*User, error) {
	db := GetDb()

	user := new(User)

	userDocument, err := db.Query(UserCollectionName).Where((*c.Criteria)(c.Field("discordId").Eq(discordId))).FindFirst()
	if err != nil {
		return nil, err
	}

	err = userDocument.Unmarshal(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func FindUserByLeetcodeId(leetcodeId string) (*User, error) {
	db := GetDb()

	user := new(User)

	userDocument, err := db.Query(UserCollectionName).Where((*c.Criteria)(c.Field("leetcodeId").Eq(leetcodeId))).FindFirst()
	if err != nil {
		return nil, err
	}

	err = userDocument.Unmarshal(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
