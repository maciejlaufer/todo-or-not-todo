package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/maciejlaufer/todoornottodo/db/sqlc"
	"gopkg.in/guregu/null.v4"
)

type createUserRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type createUserResponse struct {
	Email     string      `json:"email" binding:"required"`
	FirstName null.String `json:"first_name"`
	LastName  null.String `json:"last_name"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: db.NewNullString(req.FirstName),
		LastName:  db.NewNullString(req.LastName),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	firstName := null.String{}
	firstName.String = user.FirstName.String
	firstName.Valid = user.FirstName.Valid

	lastName := null.String{}
	lastName.String = user.LastName.String
	lastName.Valid = user.LastName.Valid

	response := createUserResponse{
		Email:     user.Email,
		FirstName: firstName,
		LastName:  lastName,
	}

	ctx.JSON(http.StatusOK, response)
}
