package user_service

import (
	"errors"
	"http2/app/types/userDB"
)

func (srv *Service) SigninUser(info user_types.User) (*user_types.User, error) {
	err := user_types.UserValidate(info.Login, info.Password, info.Email)
	if err != nil {
		return nil, err
	}

	data, err := srv.storage.SaveUser(info)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) Login(creds user_types.Credential) (*user_types.User, error) {
	data, err := srv.storage.GetUser(creds)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (srv *Service) GetAllUser() ([]user_types.User, error) {
	data, err := srv.storage.GetAllUser()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) GetUserByID(id uint64) (*user_types.User, error) {
	data, err := srv.storage.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) GetUserByIDs(ids []int) ([]user_types.User, error) {
	data, err := srv.storage.GetUserByIDs(ids)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) GetUserByLogin(str string) (*user_types.User, error) {
	data, err := srv.storage.GetUserByLogin(str)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (srv *Service) UpdateUser(upd user_types.UpdateUserRequestDTO) (*user_types.UpdateUserResponseDTO, error) {
	varUser := make(map[string]interface{})

	if upd.Login != nil && len(*upd.Login) != 0 {
		varUser["login"] = *upd.Login
	}

	if upd.Email != nil && len(*upd.Email) != 0 {
		varUser["email"] = *upd.Email
	}

	if upd.Password != nil && len(*upd.Password) != 0 {
		varUser["password"] = *upd.Password
	}

	if upd.Wallet != nil && len(*upd.Wallet) != 0 {
		varUser["wallet"] = *upd.Wallet
	}

	if len(varUser) == 0 {
		return nil, errors.New("Nothing to chenged")
	}

	data, err := srv.storage.Update(uint64(upd.ID), varUser)
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