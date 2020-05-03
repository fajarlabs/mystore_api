package response_body

import (
    "github.com/restapi_go/entity_model"
)

type ResponseCommands struct {
    Status      string `json:"Status"`
    Description string `json:"Description"`
}

type ResponseLogin struct {
    Status string `json:"Status"`
    Result entity_model.App_user `json:"Result"`
}