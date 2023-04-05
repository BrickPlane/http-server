package storage

import (
	// "context"
	"database/sql"
	"fmt"
	"http2/app/types"
	"http2/app/types/erors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func (stor *Storage) SaveUser(val types.User) (*types.User, error) {
	data, err := stor.DB.Query(`
		INSERT INTO "user" (email, login, password)
		VALUES ($1, $2, $3)
		RETURNING *`,
		val.Email, val.Login, val.Password,
	)
	if err != nil {
		return nil, erors.InvalidLogin
	}

	cred := &types.User{}
	for data.Next() {
		if err := data.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
		); err != nil {
			return nil, erors.CredError
		}
	}

	return cred, nil
}

func (stor *Storage) GetUser(val types.Credential) (*types.User, error) {
	data, err := stor.DB.Query(`SELECT * FROM "user" WHERE login=$1 AND password=$2`, val.Login, val.Password)
	if err != nil { 
		return nil, err
	}
	
	cred := &types.User{}
	for data.Next() {
		if err := data.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
		); err != nil {
			return nil, erors.CredError
		}
	}
	return cred, nil
}

func (stor *Storage) GetAllUser() ([]types.User, error) {
	val, err := stor.DB.Query(`SELECT * FROM "user"`)
	if err != nil {
		return nil, err
	}

	mk := make([]types.User, 0)
	for val.Next() {
		var cred types.User
		if err := val.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
		); err != nil {
			return nil, err
		}
		mk = append(mk, cred)
	}
	return mk, nil
}

func (stor *Storage) GetUserByID(id uint64) (*types.User, error) {
	val := &types.User{}
	err := stor.DB.Get(val, `SELECT * FROM "user" WHERE id=$1`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	fmt.Println(val)
	return val, nil
}

func (stor *Storage) GetUserByLogin(str string) (*types.User, error) {
	data, err := stor.DB.Query(`SELECT * FROM "user" WHERE login=$1`, str)
	if err != nil {
		return nil, err
	}

	cred := &types.User{}
	for data.Next() {
		if err := data.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
		); err != nil {
			return nil, erors.CredError
		}
	}
	return cred, nil
}

func (stor *Storage) GetUserByIDs(ids []int) ([]types.User, error) {
	query, args, err := sqlx.In(`SELECT * FROM public."user" WHERE id IN (?)`, ids)

	data, err := stor.DB.Query(stor.DB.Rebind(query), args...)
	if err != nil {
		return nil, err
	}

	users := make([]types.User, 0)
	for data.Next() {
		var cred types.User
		if err := data.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
		); err != nil {
			return nil, err
		}
		users = append(users, cred)
	}
	return users, err
}
func (stor *Storage) Update(val types.User) (*types.User, error) {
	data, err := stor.DB.Query(`UPDATE "user" SET password=$1, email=$2 WHERE login=$3 RETURNING *;`, val.Password, val.Email, val.Login)
	if err != nil {
		return nil, err
	}

	cred := &types.User{}
	for data.Next() {
		if err := data.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
		); err != nil {
			return nil, err
		}
	}
	fmt.Println(cred)
	return cred, nil
}

func (stor *Storage) Delete(val uint64) error {
	dest := &types.User{}
	err := stor.DB.Get(dest, `DELETE FROM "user" WHERE id=$1`, val)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}
	fmt.Println(dest)
	return nil
}

func (stor *Storage) FindUser(val uint64) error {
	fmt.Println("..... val ", val)
	_, err := stor.DB.Query(`SELECT * FROM "user" WHERE id=$1`, val)
	if err != nil {
		fmt.Println("[finduser] error")
		return err
	}

	fmt.Println("no error")
	return nil
}
