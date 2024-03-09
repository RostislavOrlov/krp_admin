package repositories

import (
	"context"
	"errors"
	"github.com/RostislavOrlov/krp_admin/internal/dto"
	"github.com/RostislavOrlov/krp_admin/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"log"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) (*UserRepository, error) {
	return &UserRepository{
		db: db,
	}, nil
}

func (repo *UserRepository) AddUser(req dto.AddUserRequest) (*dto.AddUserResponse, error) {
	q := "INSERT INTO users (lastname, firstname, middlename, email, pswd, passport, inn, snils, birthday, role) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *"

	row, err := repo.db.Query(context.Background(), q,
		req.LastName, req.FirstName, req.MiddleName, req.Email, req.Password,
		req.Passport, req.Inn, req.Snils, req.Birthday, req.Role)
	if err != nil && err.Error() != "no row in result set" {
		return nil, err
	}

	defer row.Close()
	var usrDb dto.AddUserResponse
	for row.Next() {
		err = row.Scan(&usrDb.EmployeeId, &usrDb.LastName, &usrDb.FirstName,
			&usrDb.MiddleName, &usrDb.Email, &usrDb.Password,
			&usrDb.Passport, &usrDb.Inn, &usrDb.Snils, &usrDb.Birthday, &usrDb.Role)
		if err != nil {
			log.Fatalf("Unable to scan row: %v\n", err)
		}
	}
	return &usrDb, nil
}

func (repo *UserRepository) EditUser(req dto.EditUserRequest) (*dto.EditUserResponse, error) {
	logrus.Info("req EditUser из слоя репозитория: ", req)
	q := "update users set lastname=$1, firstname=$2, middlename=$3, email=$4, pswd=$5, passport=$6, inn=$7, snils=$8, birthday=$9, role=$10 where user_id=$11 returning *"
	row, err := repo.db.Query(context.Background(), q, req.LastName, req.FirstName, req.MiddleName, req.Email, req.Password,
		req.Passport, req.Inn, req.Snils, req.Birthday, req.Role, req.EmployeeId)
	if err != nil && err.Error() != "no row in result set" {
		return nil, errors.New("error edit user in table:" + err.Error())
	}

	defer row.Close()
	var usrDb dto.EditUserResponse
	for row.Next() {
		err = row.Scan(&usrDb.EmployeeId, &usrDb.LastName, &usrDb.FirstName,
			&usrDb.MiddleName, &usrDb.Email, &usrDb.Password,
			&usrDb.Passport, &usrDb.Inn, &usrDb.Snils, &usrDb.Birthday, &usrDb.Role)
		if err != nil {
			log.Fatalf("Unable to scan row: %v\n", err)
		}
	}

	return &usrDb, nil
}

func (repo *UserRepository) DeleteUser(req dto.DeleteUserRequest) (*dto.DeleteUserResponse, error) {
	q := "delete from users where user_id=$1"
	_, err := repo.db.Query(context.Background(), q, req.EmployeeId)
	if err != nil && err.Error() != "no row in result set" {
		return nil, errors.New("error delete user from table:" + err.Error())
	}

	var resp dto.DeleteUserResponse
	resp.EmployeeId = req.EmployeeId

	return &resp, nil
}

func (repo *UserRepository) ListUsers() (*dto.ListUserResponse, error) {
	q := "select * from users"
	row, err := repo.db.Query(context.Background(), q)
	if err != nil && err.Error() != "no row in result set" {
		return nil, errors.New("error list users in table:" + err.Error())
	}

	defer row.Close()
	var resp dto.ListUserResponse
	for row.Next() {
		var tempUser entities.User
		err = row.Scan(&tempUser.Id, &tempUser.LastName, &tempUser.FirstName,
			&tempUser.MiddleName, &tempUser.Email, &tempUser.Password,
			&tempUser.Passport, &tempUser.Inn, &tempUser.Snils, &tempUser.Birthday, &tempUser.Role)
		if err != nil {
			log.Fatalf("Unable to scan row: %v\n", err)
		}

		resp.Users = append(resp.Users, tempUser)
	}

	return &resp, nil
}
