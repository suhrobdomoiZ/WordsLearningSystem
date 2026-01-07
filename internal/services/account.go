package services

import "github.com/suhrobdomoiZ/WordsLearningSystem/internal/dto"

type Account struct {
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) SignIn(username, password string) *dto.Token{ 

	return &dto.Token{
		Value: "",
		Signature: "",
	}
}

func (a *Account) SignUp() {

}
