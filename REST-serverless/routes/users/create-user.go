package users_route

import (
	"REST-serverless/db"
	"REST-serverless/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

type CreateUserRequest struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Img         string `json:"img"`
	Password    string `json:"password"`
	Description string `json:"desc"`
}

func CreateUserFromEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqPayload CreateUserRequest
		err := c.ShouldBindJSON(&reqPayload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = utils.ValidatePassword(reqPayload.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = utils.ValidateUsername(reqPayload.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hashedPassword, err := utils.HashPassword(reqPayload.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userData, err := db.DBQueries.CreateUserOrThrow(c.Request.Context(), db.CreateUserOrThrowParams{
			Name:     reqPayload.Name,
			Username: reqPayload.Username,
			Email:    reqPayload.Email,
			Img: pgtype.Text{
				String: reqPayload.Img,
				Valid:  true,
			},
			Description: pgtype.Text{
				String: reqPayload.Description,
				Valid:  true,
			},
			Password: pgtype.Text{
				String: hashedPassword,
				Valid:  true,
			},
			Provider: "email",
			Verified: false,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"success": true, "id": userData.ID})
	}
}
