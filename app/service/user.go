package service

import (
	"http2/app/types"
)

func (srv *Service) SigninUser(info types.User) (*types.User, error) {
	err := types.UserValidate(info.Login, info.Password, info.Email)
	if err != nil {
		return nil, err
	}

	data, err := srv.storage.SaveUser(info)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) Login(creds types.Credential) (*types.User, error) {
	data, err := srv.storage.GetUser(creds); 
	if  err != nil {
		return nil, err
	} 
	
	return data, nil
}

func (srv *Service) GetAllUser() ([]types.User, error) {
	data, err := srv.storage.GetAllUser()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv* Service) GetUserByID(id uint64) (*types.User, error) {
	data, err := srv.storage.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) GetUserByIDs(ids []int) ([]types.User, error) {
	data, err := srv.storage.GetUserByIDs(ids)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) GetUserByLogin(str string) (*types.User, error) {
	data, err := srv.storage.GetUserByLogin(str)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) UpdateUser(upd types.User) (*types.User, error) {
	data, err := srv.storage.Update(upd)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) DeleteUser(dlt uint64) error {
	err := srv.storage.Delete(dlt)
	if err != nil {
		return err
	}
	return nil
}
