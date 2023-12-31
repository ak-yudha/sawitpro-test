// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message string `json:"message"`
}

// FullName defines model for FullName.
type FullName = string

// HelloResponse defines model for HelloResponse.
type HelloResponse struct {
	Message string `json:"message"`
}

// LoginRequest defines model for LoginRequest.
type LoginRequest struct {
	Password    Password    `json:"password" validate:"required,custom_password"`
	PhoneNumber PhoneNumber `json:"phone_number" validate:"required,custom_prefix=+62,min=11,max=14"`
}

// LoginResponse defines model for LoginResponse.
type LoginResponse struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

// Password defines model for Password.
type Password = string

// PhoneNumber defines model for PhoneNumber.
type PhoneNumber = string

// ProfileRequest defines model for ProfileRequest.
type ProfileRequest struct {
	FullName    *FullName    `json:"full_name,omitempty" validate:"required,min=3,max=60"`
	PhoneNumber *PhoneNumber `json:"phone_number,omitempty" validate:"required,custom_prefix=+62,min=11,max=14"`
}

// ProfileResponse defines model for ProfileResponse.
type ProfileResponse struct {
	FullName    *FullName    `json:"full_name,omitempty" validate:"required,min=3,max=60"`
	PhoneNumber *PhoneNumber `json:"phone_number,omitempty" validate:"required,custom_prefix=+62,min=11,max=14"`
}

// ProfileUpdateResponse defines model for ProfileUpdateResponse.
type ProfileUpdateResponse struct {
	Token *string `json:"Token,omitempty"`
}

// RegistrationRequest defines model for RegistrationRequest.
type RegistrationRequest struct {
	FullName    FullName    `json:"full_name" validate:"required,min=3,max=60"`
	Password    Password    `json:"password" validate:"required,custom_password"`
	PhoneNumber PhoneNumber `json:"phone_number" validate:"required,custom_prefix=+62,min=11,max=14"`
}

// RegistrationResponse defines model for RegistrationResponse.
type RegistrationResponse struct {
	UserId *int `json:"user_id,omitempty"`
}

// HelloParams defines parameters for Hello.
type HelloParams struct {
	Id int `form:"id" json:"id"`
}

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = LoginRequest

// UpdateProfileJSONRequestBody defines body for UpdateProfile for application/json ContentType.
type UpdateProfileJSONRequestBody = ProfileRequest

// RegistrationJSONRequestBody defines body for Registration for application/json ContentType.
type RegistrationJSONRequestBody = RegistrationRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// This is just a test endpoint to get you started. Please delete this endpoint.
	// (GET /hello)
	Hello(ctx echo.Context, params HelloParams) error
	// Endpoint User Login
	// (POST /login)
	Login(ctx echo.Context) error
	// Endpoint Get Profile Base on Token Login
	// (GET /profile)
	Profile(ctx echo.Context) error
	// Endpoint Update Profile
	// (PATCH /profile)
	UpdateProfile(ctx echo.Context) error
	// Add New User Registration
	// (POST /registration)
	Registration(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Hello converts echo context to params.
func (w *ServerInterfaceWrapper) Hello(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params HelloParams
	// ------------- Required query parameter "id" -------------

	err = runtime.BindQueryParameter("form", true, true, "id", ctx.QueryParams(), &params.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Hello(ctx, params)
	return err
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Login(ctx)
	return err
}

// Profile converts echo context to params.
func (w *ServerInterfaceWrapper) Profile(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Profile(ctx)
	return err
}

// UpdateProfile converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateProfile(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateProfile(ctx)
	return err
}

// Registration converts echo context to params.
func (w *ServerInterfaceWrapper) Registration(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Registration(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/hello", wrapper.Hello)
	router.POST(baseURL+"/login", wrapper.Login)
	router.GET(baseURL+"/profile", wrapper.Profile)
	router.PATCH(baseURL+"/profile", wrapper.UpdateProfile)
	router.POST(baseURL+"/registration", wrapper.Registration)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xX32/bNhD+VwiuD80q23JsGKuAYEiHdguwBUGavjR2A0Y628wokiGPqT1D//tAyopl",
	"R24KzM5SIE+m5dPdx+++++EFTVWulQSJliYLatMp5Cwc3xujzDlYraQF/0AbpcEgh/BzDtaySfgB5xpo",
	"Qi0aLie0KCJq4NZxAxlNLu8NR1FlqK5vIEVaRPSDE+KU5cFLzmY8dzlNBnFEcy7LL71ow31EZy3FNG+l",
	"KoMJyBbM0LAWskmAdccEzxj6FyoQUc7lUS/K2exoEAd4f4AQas9X+1NNuDyHWwcWH4bQzNqvymT+/MrA",
	"mCb0p84qFZ1lHjpnlV0RUT1VEq6ky6/BPPqetz0tTTdRr/mJVlC+cYttTPGsRhKXCBMfL6Ko/gb5OH88",
	"o5VtU/CzGkkrdfRr6hh4+IhgJE3ol9e/HrV/vjxufR4dhONwmJWHyy/Hrc+s9U/cejs6aC8G0aBfvKI7",
	"UFbqLKr86p5Cf8E69XXg3V4NeDdeRz4cvhkcHrweDrPFL1G3WxzsFJ6BMZ8dvRkchlLodkMtdPslXKPG",
	"XMBWpY6dEFdyWaPfktx9Lf8nqdYRbVPd/wbpk/YMbwd2sU32RUTPYcItGoZcyR2T/Xx6yQr/I31lnY1t",
	"fDoL5qq5xxQPnPpHXI6VtxY8haXHkk3618lFaEwchf/6yYIhH8Hc8dRjvQNjufKl2G3H7dhbKg2SaU4T",
	"2guPQr1OA6zO1M8Pf5pASKLHHK5yktGknC7B3rAcEIylyeWCcu/+1oGZ06hCFRrgik80DqLlDG689Mhb",
	"l2wFJIdx7D9SJRFkgMK0FjwNYDo3VsnVUH8szetDMdCZgU0N11hScwEWiQF0RnqC+nF/Z7HXd42G2KcK",
	"yVg5WbZY6/KcmbnHNOWWcEtunEXCCHqIIDOtuESCikwAyVw5YpEZhKxNzgQwCyQDAQgE/euVfTv47gg/",
	"8oIWlW3IbpiIy6yBxXcqm++MhrWdoYGFeuyVYoo9imJ9/jdA+ujSFKwdO0FKdEEZ8dMp40SGyUdOpHb4",
	"bHT5vtJgaDRLZry6dDlJtjaP5aShe8zp5nz9QbLae7rox+H25IMy1zzLQD4/Wf0OSJZ5JO98Q1OShO1j",
	"lS/NMJ0+FFi5w9Rltvs+trFTbu1kxf5VvrGyvWj9OWvdR377dJF/U3IseNgdm3t30E5VZ2X/NrXFdfuS",
	"UF9v91RjTf8nmtKbZYSRU/gaJtF37A7dPSH8rgKsv0JSAwwhe9kogip9Iqs0rvFUOrFg7qp/Gs4ImtAp",
	"ok46HaFSJqZepcWo+DcAAP//zlAc0O8TAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
