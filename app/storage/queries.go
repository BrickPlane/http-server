package storage

import (
	"http2/app/types"
	"http2/app/types/erors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TODO email and login must be unique ✓
// TODO handle errors from DB ✓
// TODO must be get user by ID ✓
// TODO must be get user by IDs ✓
// TODO for creating user table use migration package (https://github.com/golang-migrate/migrate) ✓

func (stor *Storage) SaveUser(val types.Credential) (*types.Credential, error) {
	data, err := stor.DB.Query(`
		INSERT INTO "user" (email, login, password)
		VALUES ($1, $2, $3)
		RETURNING *`,
		val.Email, val.Login, val.Password,
	)
	if err != nil {
		return nil, erors.InvalidLogin
	}

	cred := &types.Credential{}
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

func (stor *Storage) GetUser() ([]types.Credential, error) {
	val, err := stor.DB.Query(`SELECT * FROM "user"`)
	if err != nil {
		return nil, err
	}

	mk := make([]types.Credential, 0)
	for val.Next() {
		var cred types.Credential
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
	return mk, err
}

func (stor *Storage) GetUserByID(val types.Credential) (*types.Credential, error) {
	data, err := stor.DB.Query(`SELECT * FROM "user" WHERE id=$1`, val.ID)
	if err != nil {
		return nil, err
	}

	cred := &types.Credential{}
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

func (stor *Storage) GetUserByIDs(ids []int) ([]types.Credential, error) {
	query, args, err := sqlx.In(`SELECT * FROM public."user" WHERE id IN (?)`, ids)
	
	data, err := stor.DB.Query(stor.DB.Rebind(query), args...)
	if err != nil {
		return nil, err
	}

	users := make([]types.Credential, 0)
	for data.Next() {
		var cred types.Credential
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

func (stor *Storage) Update(val types.Credential) (*types.Credential, error) {
	data, err := stor.DB.Query(`UPDATE "user" SET password=$1, login=$2 WHERE email=$3;`, val.Password, val.Login, val.Email)
	if err != nil {
		return nil, err
	}

	cred := &types.Credential{}
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

func (stor *Storage) Delete(val types.Credential) error {
	_, err := stor.DB.Query(`DELETE FROM "user" WHERE login=$1`, val.Login)
	if err != nil {
		return err
	}
	return nil
}
