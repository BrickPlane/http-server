package user_storage

import (
	"database/sql"
	"fmt"
	"http2/app/types/userDB"
	"http2/app/types/errors"
	"http2/app/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sqlx.DB
}

func NewStorage() (*Storage, error) {
	con, err := storage.ConnectDB()
	if err != nil {
		return nil, err
	}

	return &Storage{DB: con}, nil

}

func (stor *Storage) SaveUser(val user_types.User) (*user_types.User, error) {
	data, err := stor.DB.Query(`
		INSERT INTO "users" (email, login, password, wallet)
		VALUES ($1, $2, $3, $4)
		RETURNING *`,
		val.Email, val.Login, val.Password, val.Wallet,
	)
	if err != nil {
		return nil, errors.InvalidLogin
	}

	cred := &user_types.User{}
	for data.Next() {
		if err := data.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
			&cred.Wallet,
		); err != nil {
			return nil, errors.CredError
		}
	}

	return cred, nil
}

func (stor *Storage) GetUser(val user_types.Credential) (*user_types.User, error) {
	data, err := stor.DB.Query(`SELECT * FROM "users" WHERE login=$1 AND password=$2`, val.Login, val.Password)
	if err != nil {
		return nil, err
	}

	cred := &user_types.User{}
	for data.Next() {
		if err := data.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
			&cred.Wallet,
		); err != nil {
			return nil, errors.CredError
		}
	}
	return cred, nil
}

func (stor *Storage) GetAllUser() ([]user_types.User, error) {
	val, err := stor.DB.Query(`SELECT * FROM "users"`)
	if err != nil {
		return nil, err
	}

	mk := make([]user_types.User, 0)
	for val.Next() {
		var cred user_types.User
		if err := val.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
			&cred.Wallet,
		); err != nil {
			return nil, errors.CredError
		}
		mk = append(mk, cred)
	}
	return mk, nil
}

func (stor *Storage) GetUserByID(id uint64) (*user_types.User, error) {
	val := &user_types.User{}
	err := stor.DB.Get(val, `SELECT * FROM "users" WHERE id=$1`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return val, nil
}

func (stor *Storage) GetUserByLogin(str string) (*user_types.User, error) {
	data, err := stor.DB.Query(`SELECT * FROM "users" WHERE login=$1`, str)
	if err != nil {
		return nil, err
	}

	cred := &user_types.User{}
	for data.Next() {
		if err := data.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
			&cred.Wallet,
		); err != nil {
			return nil, errors.CredError
		}
	}
	return cred, nil
}

func (stor *Storage) GetUserByIDs(ids []int) ([]user_types.User, error) {
	query, args, err := sqlx.In(`SELECT * FROM public."users" WHERE id IN (?)`, ids)

	data, err := stor.DB.Query(stor.DB.Rebind(query), args...)
	if err != nil {
		return nil, err
	}

	users := make([]user_types.User, 0)
	for data.Next() {
		var cred user_types.User
		if err := data.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
			&cred.Wallet,
		); err != nil {
			return nil, errors.CredError
		}
		users = append(users, cred)
	}
	return users, err
}
func (stor *Storage) Update(id uint64, val map[string]interface{}) (*user_types.UpdateUserResponseDTO, error) {
	makeVal := storage.HelperForUpdate(val)

	query := fmt.Sprintf(
		`UPDATE "users" SET %[1]s WHERE id=%[2]d RETURNING *`,
		makeVal, id)
	data, err := stor.DB.NamedQuery(query, val)
	if err != nil {
		return nil, err
	}

	cred := &user_types.UpdateUserResponseDTO{}
	for data.Next() {
		if err := data.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
			&cred.Wallet,
		); err != nil {
			return nil, errors.CredError
		}
	}
	return cred, nil
}

func (stor *Storage) Delete(id uint64) error {
	dest := &user_types.User{}
	err := stor.DB.Get(dest, `DELETE FROM "users" WHERE id=$1`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}

	return nil
}
