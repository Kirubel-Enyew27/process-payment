package response

import (
	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(ctx *gin.Context, statusCode int, data interface{},
	metaData *MetaData) {

	ctx.JSON(
		statusCode,
		Response{
			OK:       true,
			MetaData: metaData,
			Data:     data,
		},
	)
}

func SendErrorResponse(ctx *gin.Context, err *ErrorResponse) {
	ctx.AbortWithStatusJSON(err.StatusCode, Response{
		OK:    false,
		Error: err,
	})
}
