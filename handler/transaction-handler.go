package handler

import (
	"net/http"
	"strconv"

	"github.com/khalidalhabibie/APIWITHGO/repository"
	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	TransactionProduct(*gin.Context)
}

type transactionHandler struct {
	repo repository.TransactionRepository
}


func NewTransactionHandler() TransactionHandler {
	return &transactionHandler{
		repo: repository.NewTransactionRepository(),
	}
}

func (h *transactionHandler) TransactionProduct(ctx *gin.Context) {
	prodIDStr := ctx.Param("product")
	if prodID, err := strconv.Atoi(prodIDStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		quantityIDStr := ctx.Param("quantity")
		if quantityID, err := strconv.Atoi(quantityIDStr); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			userID := ctx.GetFloat64("userID")
			if err := h.repo.TransactionProduct(int(userID), prodID, quantityID); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.String(http.StatusOK, "Transaction Successfully")
			}
		}
	}

}