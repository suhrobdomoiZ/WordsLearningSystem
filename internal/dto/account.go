package dto

type User struct {
	Username string
}

type Token struct {
	Value     string
	Signature string
}
