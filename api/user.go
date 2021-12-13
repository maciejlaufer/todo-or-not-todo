package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/maciejlaufer/todoornottodo/db/sqlc"
	"gopkg.in/guregu/null.v4"
)

type createUserRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type userResponse struct {
	Email     string      `json:"email" binding:"required"`
	FirstName null.String `json:"first_name"`
	LastName  null.String `json:"last_name"`
}

func formatUserResponse(user db.User) userResponse {
	firstName := null.String{}
	firstName.String = user.FirstName.String
	firstName.Valid = user.FirstName.Valid

	lastName := null.String{}
	lastName.String = user.LastName.String
	lastName.Valid = user.LastName.Valid

	return userResponse{
		Email:     user.Email,
		FirstName: firstName,
		LastName:  lastName,
	}
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

	response := formatUserResponse(user)
	ctx.JSON(http.StatusOK, response)
}

type getUserRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserById(ctx, uuid.Must(uuid.Parse(req.ID)))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := formatUserResponse(user)
	ctx.JSON(http.StatusOK, response)
}

type getUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=20"`
}

func (server *Server) getUsers(ctx *gin.Context) {
	var req getUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.GetUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := []userResponse{}

	for _, user := range users {
		response = append(response, formatUserResponse(user))
	}
	ctx.JSON(http.StatusOK, response)
}
