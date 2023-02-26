package database

import (
	"errors"

	c "github.com/ostafen/clover"
)

const UserCollectionName = "users"

type User struct {
	id         string `clover:"_id"`
	discordId  string `clover:"discordId"`
	leetcodeId string `clover:"leetcodeId"`
}

func InitUserCollection() error {
	db := GetDb()

	if hasCollection, _ := db.HasCollection(UserCollectionName); !hasCollection {
		err := db.CreateCollection(UserCollectionName)
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateUser(user *User) (*User, error) {
	db := GetDb()

	updates := make(map[string]interface{})
	updates["leetcodeId"] = user.leetcodeId
	updates["discordId"] = user.discordId

	err := db.Query(UserCollectionName).Where((*c.Criteria)(c.Field("id").Eq(user.id))).Update(updates)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// TODO: Improve error handling
func InsertUser(inputUser *User) (*User, error) {
	db := GetDb()

	userWithDiscordId, _ := FindUserByDiscordId(inputUser.discordId)
	userWithLeetcodeId, _ := FindUserByLeetcodeId(inputUser.leetcodeId)

	if userWithDiscordId != nil {
		userWithDiscordId.leetcodeId = inputUser.leetcodeId
		return UpdateUser(userWithDiscordId)
	}

	if userWithLeetcodeId != nil {
		return nil, errors.New("Leetcode Id already claimed")
	}

	user := new(User)

	userDocument := c.NewDocumentOf(inputUser)

	_, err := db.InsertOne(UserCollectionName, userDocument)
	if err != nil {
		return nil, err
	}

	err = userDocument.Unmarshal(user)

	if err != nil {
		return nil, err
	}

	return user, nil
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
