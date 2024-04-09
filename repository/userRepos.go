package repository

import (
	"context"
	"log"

	"github.com/ajaypp123/golang-jwt-microservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepoI interface {
	GetUserByKey(ctx context.Context, key string, val any) (*models.User, error)
	UserExistsByEmail(ctx context.Context, email *string) (bool, error)
	IsertUser(ctx context.Context, user *models.User) (any, error)
}

type userRepo struct {
	userCollection *mongo.Collection
}

func NewUserRepo() userRepoI {
	return &userRepo{
		userCollection: OpenCollection(Client, "user"),
	}
}

// GetUserById implements userRepoI.
func (u *userRepo) GetUserByKey(ctx context.Context, key string, val any) (*models.User, error) {
	var foundUser models.User
	// Define a filter document based on the provided key and value
	filter := bson.M{key: val}

	// Find one document matching the filter in the user collection
	result := u.userCollection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No document found with the specified key-value pair
			return nil, nil
		}
		// Handle other errors
		return nil, err
	}

	// Decode the found document into a User struct
	err = result.Decode(&foundUser)
	if err != nil {
		return nil, err
	}

	// Return the decoded user object
	return &foundUser, nil
}

// IsertUser implements userRepoI.
func (u *userRepo) IsertUser(ctx context.Context, user *models.User) (any, error) {
	resultInsertionNumber, insertErr := u.userCollection.InsertOne(ctx, user)
	return resultInsertionNumber.InsertedID, insertErr
}

// UserExistsByEmail implements userRepoI.
func (u *userRepo) UserExistsByEmail(ctx context.Context, email *string) (bool, error) {
	count, err := u.userCollection.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		log.Panic("error occured while checking for the email")
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
