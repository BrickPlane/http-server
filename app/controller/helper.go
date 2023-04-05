package controller

import (
	"encoding/json"
	"http2/app/types"

)

func  convertToType[inputType any, outputType types.Credential | types.UserID | types.User](inpurtData inputType) (*outputType, error) {
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