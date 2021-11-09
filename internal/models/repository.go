package models

import "context"

type CardTC struct {
	NumTC      string `json:"num_tc"`
	ModelTC    string `json:"model_tc"`
	DriverName string `json:"driver_name"`
	TypeTC     string `json:"type_tc"`
	NumPric    string `json:"num_pric"`
	NumPlomb   string `json:"num_plomb"`
	ContNum    string `json:"cont_num"`
	Remark     string `json:"remark"`
	SizeTC     string `json:"sizetc"`
	UserID     string
}

//Repository interface repo urls.
type Repository interface {
	SaveCard(context.Context, *CardTC) error
	CheckDBConnection(context.Context) error
	CreateUser(context.Context) (string, error)
}
