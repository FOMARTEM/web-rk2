package entities

type Question struct {
	ID       int    `json:"id,omitempty"`
	Question string `json:"question"`
	Ans1     string `json:"ans1"`
	Ans2     string `json:"ans2"`
	Ans3     string `json:"ans3"`
	Ans4     string `json:"ans4"`
	Points   int    `json:"points"` //номер ответа
}
