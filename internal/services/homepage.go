package services

import (
	"time"

	"github.com/suhrobdomoiZ/WordsLearningSystem/internal/dto"
)

type Homepage struct{}

func NewHomepage() *Homepage {
	return &Homepage{}
}

func (h *Homepage) GetBuildData() *dto.Homepage {
	return &dto.Homepage{
		Words: []dto.Word{
			{
				English: "run",
				Russian: "бежать",
			},
			{
				English: "swim",
				Russian: "плыть",
			},
			{
				English: "believe",
				Russian: "верить",
			},
		},
		Year: time.Now().Year(),
	}
}
