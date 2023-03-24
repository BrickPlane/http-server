package service

import (
	"http2/app/types"
)

func (srv *Service) SigninUser(creds types.Credential) (*types.Credential, error) {
	err := types.LoginValidate(creds.Login, creds.Password)
	if err != nil {
		return nil, err
	}

	data, err := srv.storage.SaveUser(creds)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) GetUser() ([]types.Credential, error) {
	data, err := srv.storage.GetUser()
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (srv* Service) GetUserByID(get types.Credential) (*types.Credential, error) {
	data, err := srv.storage.GetUserByID(get)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) GetUserByIDs(ids []int) ([]types.Credential, error) {
	data, err := srv.storage.GetUserByIDs(ids)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) UpdateUser(upd types.Credential) (*types.Credential, error) {
	data, err := srv.storage.Update(upd)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) DeleteUser(dlt types.Credential) error {
	err := srv.storage.Delete(dlt)
	if err != nil {
		return err
	}
	return nil
}
