package usecase

import "github.com/ValeryBMSTU/web-rk2/internal/entities"

type Provider interface {
	GetQuestion() (*entities.Question, error)
}
