package services

import "github.com/suhrobdomoiZ/WordsLearningSystem/internal/dto"

type Homepage struct {
}

func NewHomepage() *Homepage {
	return &Homepage{}
}

func (h *Homepage) GetBuildData() *dto.Homepage {
	return &dto.Homepage{}
}
