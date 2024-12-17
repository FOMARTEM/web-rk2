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

func (p *Provider) UpdateScore(score int) error {
	var updateScore entities.Score
	err := p.conn.QueryRow(`UPDATE \"My_game_Score\" SET score = score + $1`,
		score).
		Scan(&updateScore.Score)

	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) SetZeroScore() error {
	var updateScore entities.Score
	err := p.conn.QueryRow(`UPDATE \"My_game_Score\" SET score = 0`).
		Scan(&updateScore.Score)

	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) GetScore() (int, error) {
	var score entities.Score
	err := p.conn.QueryRow(`SELECT score FROM "My_game_Score" LIMIT 1`).
		Scan(&score.Score)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, nil
		}
		return -1, err
	}

	return score.Score, nil
}
