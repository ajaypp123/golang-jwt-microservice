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

func (u *userService) Signup(c *gin.Context, user *models.User) *models.Response {

	validationErr := u.validate.Struct(user)
	if validationErr != nil {
		return &models.Response{
			Code:    http.StatusBadRequest,
			Message: gin.H{"error": validationErr.Error()},
			Status:  "Failed",
		}
	}

	isUserExists, err := repository.NewUserRepo().UserExistsByEmail(c, user.Email)
	if err != nil {
		return &models.Response{
			Code:    http.StatusInternalServerError,
			Message: gin.H{"error": "error occured while checking for the email"},
			Status:  "Failed",
		}
	}
	if isUserExists {
		return &models.Response{
			Code:    http.StatusInternalServerError,
			Message: gin.H{"error": "this email or phone number already exists"},
			Status:  "Failed",
		}
	}
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.UserId = user.ID.Hex()
	id, insertErr := repository.NewUserRepo().IsertUser(c, user)
	if insertErr != nil {
		return &models.Response{
			Code:    http.StatusInternalServerError,
			Message: gin.H{"error": "User item was not created, " + insertErr.Error()},
			Status:  "Failed",
		}
	}
	logger.Logger.Info("Signup for user, ", id)
	return &models.Response{
		Code:    http.StatusOK,
		Message: gin.H{"user_id": id},
		Status:  "Success",
	}
}

func (u *userService) Login(c *gin.Context, user *models.User) *models.Response {

	foundUser, err := repository.NewUserRepo().GetUserByKey(c, "email", user.Email)
	if err != nil {
		return &models.Response{
			Code:    http.StatusInternalServerError,
			Message: gin.H{"error": "email or password is incorrect" + err.Error()},
			Status:  "Failed",
		}
	}
	if foundUser == nil {
		return &models.Response{
			Code:    http.StatusNotFound,
			Message: gin.H{"error": "no user found"},
			Status:  "Failed",
		}
	}

	passwordIsValid := *user.Password == *foundUser.Password
	if !passwordIsValid {
		return &models.Response{
			Code:    http.StatusInternalServerError,
			Message: gin.H{"error": "Not matching"},
			Status:  "Failed",
		}
	}

	if foundUser.Email == nil {
		return &models.Response{
			Code:    http.StatusInternalServerError,
			Message: gin.H{"error": "user not found"},
			Status:  "Failed",
		}
	}

	token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName,
		*foundUser.LastName, foundUser.UserId)

	return &models.Response{
		Code: http.StatusOK,
		Message: &models.UserDetail{
			User:         foundUser,
			Token:        &token,
			RefreshToken: &refreshToken,
		},
		Status: "Success",
	}
}

func (u *userService) GetUser(c *gin.Context, userId *string) *models.Response {

	user, err := repository.NewUserRepo().GetUserByKey(c, "userid", &userId)
	if err != nil {
		return &models.Response{
			Code:    http.StatusInternalServerError,
			Message: gin.H{"error": err},
			Status:  "Failed",
		}
	}
	if user == nil {
		return &models.Response{
			Code:    http.StatusNotFound,
			Message: gin.H{"error": "no user found"},
			Status:  "Failed",
		}
	}
	return &models.Response{
		Code:    http.StatusOK,
		Message: user,
		Status:  "Success",
	}
}
