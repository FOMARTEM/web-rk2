package provider

import (
	"database/sql"
	"errors"

	"github.com/ValeryBMSTU/web-rk2/internal/entities"
)

func (p *Provider) GetQuestion() ([]*entities.Question, error) {
	questions := []*entities.Question{}

	rows, err := p.conn.Query(`SELECT id, question, ans1, ans2, ans3, ans4, points FROM "My_game"`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return questions, nil
		}
		return nil, err
	}

	for rows.Next() {
		var question entities.Question
		if err := rows.Scan(&question.ID, &question.Question, &question.Ans1, &question.Ans2, &question.Ans3, &question.Ans4, &question.Points); err != nil {
			return nil, err
		}
		questions = append(questions, &question)
	}

	return questions, nil
}
