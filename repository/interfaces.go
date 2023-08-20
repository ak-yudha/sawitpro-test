// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
	InsertUsers(ctx echo.Context, input generated.RegistrationRequest) (userID int, err error)
	UpdateUserByPhoneNumber(ctx context.Context, req generated.ProfileRequest, phoneNumber string) (err error)
	UpdateCounterLogin(ctx context.Context, loginCounter int, phoneNumber string) (err error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (user Users, err error)

	Generated(ctx context.Context, id int, phoneNumber string) (string, error)
	Verify(authToken string) (*JwtCustomClaims, error)
}
