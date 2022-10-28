package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrorData struct {
	Status string `json:"status"`
	Source string `json:"source"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		var errors []ErrorData

		for _, err := range ctx.Errors {
			fmt.Printf("%v", err)
			errors = append(errors, ErrorData{
				Status: "", // TODO
				Source: "", // TODO { "pointer": "/data/properties/name" }
				Title:  "", // TODO "ValidationError"
				Detail: err.Error(),
			})
		}

		status := -1 // status passthrough
		ctx.JSON(status, gin.H{"errors": errors})
	}
}
