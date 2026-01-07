package handlers

import (
	"net/http"

	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/services"
)

type Account struct {
	service *services.Account
}

func NewAccount() *Account {
	return &Account{
		service: services.NewAccount(),
	}
}

func (a *Account) SignIn(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		a.signInGet(writer, req)
	case http.MethodPost:
		a.signInPost(writer, req)
	default:
		http.Error(
			writer,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed,
		)
	}
}

func (a *Account) signInGet(writer http.ResponseWriter, req *http.Request) {
}

func (a *Account) signInPost(writer http.ResponseWriter, req *http.Request) {
}
