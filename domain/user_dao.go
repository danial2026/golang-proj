package domain

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (user *User) Save() error {
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (user *User) GetByEmail(email string) ([]*User, error) {
	filter := bson.D{{"email", email}}
	return filterUsers(filter)
}

func (user *User) GetAllUsers() ([]*User, error) {
	// passing bson.D{{}} matches all documents in the collection
	filter := bson.D{{}}
	return filterUsers(filter)
}

func filterUsers(filter interface{}) ([]*User, error) {
	// A slice of Users for storing the decoded documents
	var users []*User

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return users, err
	}

	for cur.Next(ctx) {
		var t User
		err := cur.Decode(&t)
		if err != nil {
			return users, err
		}

		users = append(users, &t)
	}

	if err := cur.Err(); err != nil {
		return users, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}

	return users, nil
}
