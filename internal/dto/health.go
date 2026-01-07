package dto

type Health struct {
	Server   string `json:"serverStatus"`
	DataBase string `json:"databaseStatus"`
}
