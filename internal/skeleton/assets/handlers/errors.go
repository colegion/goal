// Package handlers is generated automatically by goal toolkit.
// Please, do not edit it manually.
package handlers

import (
	"net/http"
	"net/url"

	contr "github.com/colegion/goal/internal/skeleton/controllers"

	"github.com/colegion/goal/strconv"
)

// Errors is an insance of tErrors that is automatically generated from Errors controller
// being found at "github.com/colegion/goal/internal/skeleton/controllers/errors.go",
// and contains methods to be used as handler functions.
//
// Errors is a controller with actions displaying error pages.
var Errors tErrors

// context stores names of all controllers and packages of the app.
var context = url.Values{}

// tErrors is a type with handler methods of Errors controller.
type tErrors struct {
}

// New allocates (github.com/colegion/goal/internal/skeleton/controllers).Errors controller,
// initializes its parents; then returns the controller.
func (t tErrors) New(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Errors {
	c := &contr.Errors{}
	c.Controllers = Controllers.New(w, r, ctr, act)
	return c
}

// Before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what.
func (t tErrors) Before(c *contr.Errors, w http.ResponseWriter, r *http.Request) http.Handler {
	// Execute magic Before actions of embedded controllers.
	if h := Controllers.Before(c.Controllers, w, r); h != nil {
		return h
	}

	return nil
}

// After is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tErrors) After(c *contr.Errors, w http.ResponseWriter, r *http.Request) (h http.Handler) {

	// Execute magic After methods of embedded controllers.

	if h = Controllers.After(c.Controllers, w, r); h != nil {
		return h
	}

	return
}

// NotFound is a handler that was generated automatically.
// It calls Before, After methods, and NotFound action found at
// github.com/colegion/goal/internal/skeleton/controllers/errors.go
// in appropriate order.
//
// NotFound prints an error 404 "Page Not Found" message.
func (t tErrors) NotFound(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Errors.New(w, r, "Errors", "NotFound")
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Errors.After(c, w, r)
	if res := Errors.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.NotFound(); res != nil {
		h = res
		return
	}
}

// InternalError is a handler that was generated automatically.
// It calls Before, After methods, and InternalError action found at
// github.com/colegion/goal/internal/skeleton/controllers/errors.go
// in appropriate order.
//
// InternalError displays an error 500 "Internal Server Error" message.
func (t tErrors) InternalError(w http.ResponseWriter, r *http.Request) {
	var h http.Handler
	c := Errors.New(w, r, "Errors", "InternalError")
	defer func() {
		if h != nil {
			h.ServeHTTP(w, r)
		}
	}()
	defer Errors.After(c, w, r)
	if res := Errors.Before(c, w, r); res != nil {
		h = res
		return
	}
	if res := c.InternalError(); res != nil {
		h = res
		return
	}
}

// Init is used to initialize controllers of "github.com/colegion/goal/internal/skeleton/controllers"
// and its parents.
func Init() {

	initApp()

	initControllers()

	initErrors()

}

func initErrors() {

	context.Add("Errors", "NotFound")

	context.Add("Errors", "InternalError")

}

func init() {
	_ = strconv.MeaningOfLife
}