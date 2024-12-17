package entities

type User struct {
	ID       int    `json:"id,omitempty"`
	QUESTION string `json:"question"`
	ANS1     string `json:"ans1"`
	ANS2     string `json:"ans2"`
	ANS3     string `json:"ans3"`
	ANS4     string `json:"ans4"`
	POINTS   int    `json:"points"` //номер ответа
}
