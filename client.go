package otdbclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseUrl string
	http    *http.Client
	Token   string
}

func New(baseUrl string) Client {
	return Client{
		BaseUrl: baseUrl,
		http:    &http.Client{},
	}
}

func (c *Client) Categories(ctx context.Context) ([]Category, error) {
	categoryResponse := CategoriesResponse{}
	err := c.executeRequest(ctx, "/api_category.php", &categoryResponse)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch session token: %w", err)
	}
	return categoryResponse.TriviaCategories, nil
}

func (c *Client) RandomQuestions(ctx context.Context, amount int) ([]Question, error) {
	url := fmt.Sprintf("/api.php?amount=%d", amount)
	questionResponse := QuestionsResponse{}
	err := c.executeRequest(ctx, url, &questionResponse)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch session token: %w", err)
	}
	return questionResponse.Results, nil
}

func (c *Client) NewSessionToken(ctx context.Context) (string, error) {
	tokenResponse := NewTokenResponse{}
	err := c.executeRequest(ctx, "/api_token.php?command=request", &tokenResponse)
	if err != nil {
		return "", fmt.Errorf("unable to fetch session token: %w", err)
	}
	return tokenResponse.Token, nil
}

func (c *Client) executeRequest(ctx context.Context, endpoint string, res interface{}) error {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s", c.BaseUrl, endpoint), nil)
	if err != nil {
		return fmt.Errorf("unable to construct request: %w", err)
	}
	rs, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("unable to complete request: %w", err)
	}
	dec := json.NewDecoder(rs.Body)

	err = dec.Decode(&res)
	if err != nil {
		return fmt.Errorf("unable to decode response body: %w", err)
	}
	return nil
}
