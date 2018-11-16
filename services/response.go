package services

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
)

// Errors stores JSON Api error response
type Errors struct {
	Errors []string `json:"errors"`
}

// TransformError transforms error to JSON Api format
func TransformError(err error) (Errors, int) {
	if svc, ok := err.(IServiceError); ok {
		return Errors{[]string{svc.Error()}}, svc.Status()
	}

	panic(err)
}

// RecoverUnexpectedError recovers from any unexpected error
func RecoverUnexpectedError(c *beego.Controller) {
	if r := recover(); r != nil {
		// Create an unique id to make a relationship between
		// API error and log error
		var u1String string

		if u1, err := uuid.NewV4(); err == nil {
			u1String = fmt.Sprintf("%s", u1)
		} else {
			u1String = "Unknown"
		}

		beego.Error(fmt.Sprintf("%s - %+v", u1String, r))
		c.Ctx.ResponseWriter.WriteHeader(500)
		c.Data["json"] = map[string][]string{
			"errors": []string{fmt.Sprintf("Transaction: '%s'", u1String)},
		}
	}
}
