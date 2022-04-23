package users

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"user-login-api/core/common"
	"user-login-api/infrastructure/configs"
)

type UserAccessor struct {
}

var userCollection *mongo.Collection = configs.DB.Database("mongodb").Collection("user")

func (accessor *UserAccessor) Create(user common.User, ctx context.Context) *echo.Map {
	newUser := common.User{
		Id:    primitive.NewObjectID(),
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
		Admin: user.Admin,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return &echo.Map{"error": err.Error()}
	}

	return &echo.Map{"insertedID": result.InsertedID}
}

func (accessor *UserAccessor) FindAll(ctx context.Context) *echo.Map {
	var user common.User
	cur, err := userCollection.Find(ctx, bson.D{{}})

	var result []common.User
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		if err = cur.Decode(&user); err != nil {
			log.Fatal(err)
		}
		result = append(result, user)
	}

	if err != nil {
		return &echo.Map{"error": err.Error()}
	}

	return &echo.Map{"user": result}
}

func (accessor *UserAccessor) Find(ctx context.Context, userId string) *echo.Map {
	objId, _ := primitive.ObjectIDFromHex(userId)

	var user common.User
	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)

	if err != nil {
		return &echo.Map{"error": err.Error()}
	}

	return &echo.Map{"user": user}
}

func (accessor *UserAccessor) FindByEmail(ctx context.Context, email string) (common.User, error) {
	var user common.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, err
}

func (accessor *UserAccessor) Update(ctx context.Context, userId string, update bson.M) *echo.Map {
	objId, _ := primitive.ObjectIDFromHex(userId)
	result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return &echo.Map{"error": err.Error()}
	}

	//get updated user details
	var updatedUser common.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)

		if err != nil {
			return &echo.Map{"error": err.Error()}
		}
	}

	return &echo.Map{"user": updatedUser}
}

func (accessor *UserAccessor) Delete(ctx context.Context, userId string) *echo.Map {
	objId, _ := primitive.ObjectIDFromHex(userId)

	result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		return &echo.Map{"error": err.Error()}
	}

	if result.DeletedCount < 1 {
		return &echo.Map{"data": "User with specified ID not found!"}
	}

	return &echo.Map{"data": "User successfully deleted!"}
}
