package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"strings"
)
type FieldError struct {
	err validator.FieldError
}
func (q *FieldError) String() string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + q.err.Field() + "'")
	sb.WriteString(", condition: " + q.err.ActualTag())

	// Print condition parameters, e.g. oneof=red blue -> { red blue }
	if q.err.Param() != "" {
		sb.WriteString(" { " + q.err.Param() + " }")
	}

	if q.err.Value() != nil && q.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
	}

	return sb.String()
}
func Decode(c *gin.Context,data interface{}) []string {
	err := c.ShouldBindJSON(data)
	if err != nil {
		errorSlice := make([]string, 0)
		//errs, ok := err.(validator.ValidationErrors)
		//if ok{
		//	for _, fieldErr := range errs {
		//		fErr := &FieldError{err: fieldErr}
		//		errorSlice = append(errorSlice, fErr.String())
		//	}
		//} else {
		//
		//}
		errorSlice = append(errorSlice, err.Error())
		return errorSlice
	}
	return nil
}
