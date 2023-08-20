package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) InsertUsers(ctx echo.Context, input generated.RegistrationRequest) (userID int, err error) {
	sqlStatement := `INSERT INTO users (full_name, phone_number, password) VALUES ($1, $2,$3) RETURNING id`
	stmt, err := r.Db.Prepare(sqlStatement)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(input.FullName, input.PhoneNumber, input.Password).Scan(&userID)
	if err != nil {
		fmt.Println("errr", err)
		return 0, err
	}
	return userID, nil
}

func (r *Repository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (user Users, err error) {
	sqlStatement := `SELECT id, full_name, phone_number, password, login_counter from users where phone_number = $1`
	r.Db.QueryRowContext(ctx, sqlStatement, phoneNumber).Scan(&user.ID, &user.FullName, &user.PhoneNumber, &user.Password, &user.LoginCounter)
	if err != nil {
		panic(err)
	}

	return user, nil
}

func (r *Repository) UpdateCounterLogin(ctx context.Context, loginCounter int, phoneNumber string) (err error) {
	sqlStatement := `UPDATE users set login_counter = $1 where phone_number = $2`
	stmt, err := r.Db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(loginCounter, phoneNumber)
	if err != nil {
		fmt.Println("errr", err)
		return err
	}
	return nil
}

func (r *Repository) UpdateUserByPhoneNumber(ctx context.Context, req generated.ProfileRequest, phoneNumber string) (err error) {
	sqlStatement := `UPDATE users set `
	params := []interface{}{}
	if req.PhoneNumber != nil {
		sqlStatement += " phone_number = $1,"
		params = append(params, req.PhoneNumber)
	}

	if req.FullName != nil {
		sqlStatement += " full_name = $2,"
		params = append(params, req.FullName)
	}

	sqlStatement = sqlStatement[:len(sqlStatement)-1]
	sqlStatement += " WHERE phone_number = $3"
	params = append(params, phoneNumber)

	x, _ := json.Marshal(params)
	fmt.Println(fmt.Sprintf("%s", string(x)))
	stmt, err := r.Db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(params...)
	if err != nil {
		fmt.Println("errr", err)
		return err
	}
	return nil
}
