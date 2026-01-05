package handlers

import "github.com/suhrobdomoiZ/WordsLearningSystem/internal/services"

type Homepage struct {
	service *services.Homepage
}

func NewHomepage() *Homepage{
	return &Homepage{
		service: services.NewHomepage(),
	}
}
