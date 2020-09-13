package account

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	accountDB "integration-ginkgo-example/internal/account/database"
	"integration-ginkgo-example/internal/account/model"
	"integration-ginkgo-example/internal/database"
	"integration-ginkgo-example/pkg/logging"
	"net/http"
	"strconv"
)

type Handler struct {
	db *accountDB.AccountDB
}

func (h *Handler) GetAccounts(ctx *gin.Context) {
	logger := logging.FromContext(ctx)
	accounts, err := h.db.FindAccounts(ctx)
	if err != nil {
		logger.Errorw("failed to read accounts", "err", err)
		abortWithMessage(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

func (h *Handler) GetAccount(ctx *gin.Context) {
	logger := logging.FromContext(ctx)
	// binding
	type PathVariable struct {
		AccountID string `uri:"id" binding:"numeric"`
	}
	var path PathVariable
	if err := ctx.ShouldBindUri(&path); err != nil {
		logger.Errorw("invalid account id", "err", err)
		responseBadRequest(ctx, err)
		return
	}
	accountId, _ := strconv.ParseUint(path.AccountID, 10, 64)

	// get account by id
	account, err := h.db.FindById(ctx, uint(accountId))
	if err != nil {
		logger.Errorw("failed to find account", "accountID", path.AccountID, "err", err)
		if err == gorm.ErrRecordNotFound {
			abortWithMessage(ctx, http.StatusNotFound, "not found account: "+path.AccountID)
			return
		}
	}
	ctx.JSON(http.StatusOK, account)
}

func (h *Handler) SaveAccount(ctx *gin.Context) {
	logger := logging.FromContext(ctx)
	// binding
	type RequestBody struct {
		Username string `form:"username" binding:"required"`
		Email    string `form:"email" binding:"required"`
	}
	var body RequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		logger.Errorw("invalid account", "err", err)
		responseBadRequest(ctx, err)
		return
	}

	// save a account
	saved := &model.Account{
		Username: body.Username,
		Email:    body.Email,
	}
	if err := h.db.Save(ctx, saved); err != nil {
		logger.Errorw("failed to save a account", "err", err)
		if database.IsDuplicateEntryError(err) {
			abortWithMessage(ctx, http.StatusBadRequest, "duplicate email: "+body.Email)
			return
		}
		abortWithMessage(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, saved)
}

func (h *Handler) UpdateAccount(ctx *gin.Context) {
	logger := logging.FromContext(ctx)
	// binding
	type PathVariable struct {
		AccountID string `uri:"id" binding:"numeric"`
	}
	type RequestBody struct {
		Username string `form:"username" binding:"required"`
	}
	var (
		path PathVariable
		body RequestBody
	)
	if err := ctx.ShouldBindUri(&path); err != nil {
		logger.Errorw("invalid account id", "err", err)
		responseBadRequest(ctx, err)
		return
	}
	if err := ctx.ShouldBind(&body); err != nil {
		logger.Errorw("invalid request body", "err", err)
		responseBadRequest(ctx, err)
		return
	}
	accountID, _ := strconv.ParseUint(path.AccountID, 10, 64)
	account := model.Account{
		ID:       uint(accountID),
		Username: body.Username,
	}
	updated, err := h.db.Update(ctx, &account)
	if err != nil {
		logger.Errorw("failed to update account", "err", err)
		abortWithMessage(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if updated == 0 {
		logger.Errorw("not found account to update", "accountID", accountID)
		abortWithMessage(ctx, http.StatusNotFound, "not found account id: "+path.AccountID)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}

func (h *Handler) DeleteAccount(ctx *gin.Context) {
	logger := logging.FromContext(ctx)
	// binding
	type PathVariable struct {
		AccountID string `uri:"id" binding:"numeric"`
	}
	var path PathVariable
	if err := ctx.ShouldBindUri(&path); err != nil {
		logger.Errorw("invalid account id", "err", err)
		responseBadRequest(ctx, err)
		return
	}
	accountID, _ := strconv.ParseUint(path.AccountID, 10, 64)

	deleted, err := h.db.Delete(ctx, uint(accountID))
	if err != nil {
		logger.Errorw("failed to delete account", "err", err)
		abortWithMessage(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if deleted == 0 {
		logger.Errorw("not found account to delete", "accountID", accountID)
		abortWithMessage(ctx, http.StatusNotFound, "not found account id: "+path.AccountID)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
}

// abortWithMessage response with given status code and body({"message": "{given_message}"})
func abortWithMessage(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"message": message,
	})
}

// responseBadRequest response 400 bad request with error message
func responseBadRequest(ctx *gin.Context, err error) {
	switch err.(type) {
	case validator.ValidationErrors:
		ve := err.(validator.ValidationErrors)
		abortWithMessage(ctx, http.StatusBadRequest, fmt.Sprintf("field:%s, value:%s", ve[0].Field(), ve[0].Value()))
	default:
		abortWithMessage(ctx, http.StatusBadRequest, err.Error())
	}
}

func NewHandler(db *gorm.DB) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	h := Handler{
		db: accountDB.NewAccountDB(db),
	}
	h.initRoute(e)

	return e
}
