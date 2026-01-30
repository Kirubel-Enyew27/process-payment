package handlers

import "github.com/gin-gonic/gin"

type Payment interface {
	CreatePayment(c *gin.Context)
	GetTransactionByID(c *gin.Context)
	GetTransactions(c *gin.Context)
}
