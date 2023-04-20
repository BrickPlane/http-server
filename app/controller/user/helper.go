package user_controller

import (
	"encoding/json"
	"http2/app/types/userDB"


)

func  convertToType[inputType any, 
outputType user_types.Credential | user_types.UserID | user_types.User | user_types.UpdateUserRequestDTO](inpurtData inputType) (*outputType, error) {
	byteReq, err := json.Marshal(inpurtData)
	if err != nil {
		return nil, err
	}

	var req outputType
	if err := json.Unmarshal(byteReq, &req); err != nil {
		return nil, err
	}

	return &req, nil
}