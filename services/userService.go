package services

import (
	"net/http"
	"time"

	helpers "github.com/ajaypp123/golang-jwt-microservice/helpers"
	"github.com/ajaypp123/golang-jwt-microservice/helpers/logger"
	models "github.com/ajaypp123/golang-jwt-microservice/models"
	"github.com/ajaypp123/golang-jwt-microservice/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	validate *validator.Validate
}

func NewUserService() *userService {
	return &userService{
		validate: validator.New(),
	}
}

func (u *userService) Signup(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := u.validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	isUserExists, err := repository.NewUserRepo().UserExistsByEmail(c, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
		return
	}
	if isUserExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exists"})
		return
	}
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.UserId = user.ID.Hex()
	id, insertErr := repository.NewUserRepo().IsertUser(c, &user)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User item was not created, " + insertErr.Error()})
		return
	}
	logger.Logger.Info("Signup for user, ", id)
	c.JSON(http.StatusOK, gin.H{"user_id": id})
}

func (u *userService) Login(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := repository.NewUserRepo().GetUserByKey(c, "email", user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect" + err.Error()})
		return
	}
	if foundUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no user found"})
		return
	}

	passwordIsValid := *user.Password == *foundUser.Password
	if !passwordIsValid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not matching"})
		return
	}

	if foundUser.Email == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, foundUser.UserId)
	c.JSON(http.StatusOK, &models.UserDetail{
		User:         foundUser,
		Token:        &token,
		RefreshToken: &refreshToken,
	})
}

func (u *userService) GetUser(c *gin.Context) {
	userId := c.Param("user_id")

	user, err := repository.NewUserRepo().GetUserByKey(c, "userid", userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no user found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
