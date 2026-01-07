package services

import "github.com/suhrobdomoiZ/WordsLearningSystem/internal/dto"

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) CheckServerStatus() *dto.Health {
	return &dto.Health{
		Server:   "ok",
		DataBase: "ok",
	}
}
