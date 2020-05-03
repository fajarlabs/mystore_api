package entity_model

import (
    "time"
)


// table app_user
type App_user struct {
    Id string              `json:"Id"`
    Username string        `json:"Username"`
    Email string           `json:"Email"`
    First_name string      `json:"First_name"`
    Last_name string       `json:"Last_name"`
    Password string        `json:"Password"`
    Counter int            `json:"Counter"`
    Status int             `json:"Status"`
    Remark string          `json:"Remark"`
    Change_password string `json:"Change_password"`
    Phone string           `json:"Phone"`
    Photo string           `json:"Photo"`
    Is_delete int          `json:"Is_delete"`
    Who_delete_it string   `json:"Who_delete_it"`
}

// table process_order
type Request_order struct {
    Id string                `json:"Id"`
    Chat_id string           `json:"Chat_id"`
    Message string           `json:"Message"`
    Create_order time.Time   `json:"Create_order"`
    Is_process int           `json:"Is_process"`
    Is_delete int            `json:"Is_delete"`
}

// table process_order
type Process_order struct {
    Id string                `json:"Id"`
    Chat_id string           `json:"Chat_id"`
    Message string           `json:"Message"`
    Create_process time.Time `json:"Create_process"`
    Is_delivery int          `json:"Is_delivery"`
    Is_delete int            `json:"Is_delete"`
}