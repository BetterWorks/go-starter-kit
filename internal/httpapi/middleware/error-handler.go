package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ErrorData
type ErrorData struct {
	Status string `json:"status"`
	Source string `json:"source"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// ErrorHandler
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
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

			fmt.Printf("ERRORS in ErrorHandler: %+v\n", errors)

			// status := -1 // TODO: status passthrough ?
			// ctx.JSON(status, gin.H{"errors": errors})
		}
	}
}
