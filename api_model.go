package otdbclient

type CategoriesResponse struct {
	TriviaCategories []Category `json:"trivia_categories"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type QuestionsResponse struct {
	ResponseCode int        `json:"response_code"`
	Results      []Question `json:"results"`
}
type Question struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

type NewTokenResponse struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_message"`
	Token           string `json:"token"`
}
