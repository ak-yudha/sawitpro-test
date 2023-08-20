package handler

import (
	"fmt"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
	"net/http"
)

// This is just a test endpoint to get you started. Please delete this endpoint.
// (GET /hello)
func (s *Server) Hello(ctx echo.Context, params generated.HelloParams) error {

	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) Login(ctx echo.Context) error {

	u := new(generated.LoginRequest)
	if err := ctx.Bind(u); err != nil {
		return err
	}
	validate, err := ValidateRequest()
	if err := validate.Struct(u); err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	user, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(), u.PhoneNumber)
	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}
	err = ComparePasswords(user.Password, u.Password)
	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	token, err := s.Repository.Generated(ctx.Request().Context(), user.ID, user.PhoneNumber)
	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	err = s.Repository.UpdateCounterLogin(ctx.Request().Context(), user.LoginCounter+1, u.PhoneNumber)
	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	res := generated.LoginResponse{
		Id:    user.ID,
		Token: token,
	}

	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) Profile(ctx echo.Context) error {
	t, err := ExtractJWTToken(ctx)
	authToken := t[1]
	userData, err := s.Repository.Verify(authToken)
	if err != nil {
		err := ctx.JSON(http.StatusUnauthorized, generated.ErrorResponse{Message: err.Error()})
		return err
	}
	user, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(), userData.PhoneNumber)
	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}
	res := generated.ProfileResponse{
		FullName:    &user.FullName,
		PhoneNumber: &user.PhoneNumber,
	}

	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) UpdateProfile(ctx echo.Context) error {
	u := new(generated.ProfileRequest)
	if err := ctx.Bind(u); err != nil {
		return err
	}

	validate, err := ValidateRequest()
	if err := validate.Struct(u); err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	t, err := ExtractJWTToken(ctx)
	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	authToken := t[1]
	userData, err := s.Repository.Verify(authToken)
	if err != nil {
		err := ctx.JSON(http.StatusForbidden, generated.ErrorResponse{Message: err.Error()})
		return err
	}
	// check if exist
	userExist, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(), *u.PhoneNumber)
	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	// exist
	if userExist.PhoneNumber != "" {
		err := ctx.JSON(http.StatusConflict, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	err = s.Repository.UpdateUserByPhoneNumber(ctx.Request().Context(), *u, userData.PhoneNumber)
	if err != nil {
		err := ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	user, err := s.Repository.GetUserByPhoneNumber(ctx.Request().Context(), *u.PhoneNumber)
	if err != nil {
		err := ctx.JSON(http.StatusInternalServerError, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	// regenerate token
	token, err := s.Repository.Generated(ctx.Request().Context(), user.ID, user.PhoneNumber)
	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}
	res := generated.ProfileUpdateResponse{
		Token: &token,
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) Registration(ctx echo.Context) error {
	u := new(generated.RegistrationRequest)
	if err := ctx.Bind(u); err != nil {
		return err
	}

	validate, err := ValidateRequest()
	if err := validate.Struct(u); err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	if err := validate.Struct(u); err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	u.Password, err = HashPassword(u.Password)
	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}

	// do
	userID, err := s.Repository.InsertUsers(ctx, *u)

	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, generated.ErrorResponse{Message: err.Error()})
		return err
	}
	res := generated.RegistrationResponse{
		UserId: &userID,
	}

	return ctx.JSON(http.StatusCreated, res)
}
