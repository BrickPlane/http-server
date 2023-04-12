package storage

import (
	"database/sql"
	"fmt"
	"http2/app/types"
	"http2/app/types/erors"

	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func (stor *Storage) SaveUser(val types.User) (*types.User, error) {
	data, err := stor.DB.Query(`
		INSERT INTO "user" (email, login, password, wallet)
		VALUES ($1, $2, $3, $4)
		RETURNING *`,
		val.Email, val.Login, val.Password, val.Wallet,
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
			&cred.Wallet,
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
			&cred.Wallet,
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
			&cred.Wallet,
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
			&cred.Wallet,
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
			&cred.Wallet,
		); err != nil {
			return nil, err
		}
		users = append(users, cred)
	}
	return users, err
}
func (stor *Storage) Update(id uint64, val map[string]interface{}) (*types.UpdateUserResponseDTO, error) {
	makeVal := helperForUpdate(val)

	query := fmt.Sprintf(
		`UPDATE "user" SET %[1]s WHERE id=%[2]d RETURNING *`,
		makeVal, id)
	data, err := stor.DB.NamedQuery(query, val)
	if err != nil {
		return nil, err
	}

	cred := &types.UpdateUserResponseDTO{}
	for data.Next() {
		if err := data.Scan(
			&cred.ID,
			&cred.Email,
			&cred.Login,
			&cred.Password,
			&cred.Wallet,
		); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	fmt.Println("cred", cred)
	return cred, nil
}

func helperForUpdate(data map[string]interface{}) string {
	query := make([]string, 0, len(data))
	for key := range data {
		query = append(query, fmt.Sprintf("%[1]s=:%[1]s", key))
		// fmt.Println("val",val)
		// fmt.Println("key", key)
	}
	// fmt.Println("query", query)

	return strings.Join(query, ", ")
}

func (stor *Storage) Delete(id uint64) error {
	dest := &types.User{}
	err := stor.DB.Get(dest, `DELETE FROM "user" WHERE id=$1`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}
	fmt.Println(dest)
	return nil
}
