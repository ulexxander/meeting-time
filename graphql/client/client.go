package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	URL    string
	Logger Logger
}

type Variables map[string]interface{}

type Error struct {
	Message   string `json:"message"`
	Locations []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations"`
	Extensions struct {
		Code string `json:"code"`
	} `json:"extensions"`
}

func (e *Error) Error() string {
	msg := "GraphQL: "
	if e.Extensions.Code != "" {
		msg += e.Extensions.Code + ": "
	}
	msg += e.Message
	if len(e.Locations) != 0 {
		msg += fmt.Sprintf(" (%d:%d)", e.Locations[0].Line, e.Locations[0].Column)
	}
	return msg
}

func (c *Client) Query(ctx context.Context, query string, variables Variables, dest interface{}) error {
	if c.Logger != nil {
		c.Logger.Query(query, variables)
	}

	params := struct {
		Query     string    `json:"query"`
		Variables Variables `json:"variables"`
	}{
		Query:     query,
		Variables: variables,
	}
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("marshaling params: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bytes.NewReader(paramsJSON))
	if err != nil {
		return fmt.Errorf("initializing request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer res.Body.Close()

	resBodyJSON, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("recoding response body: %w", err)
	}

	if c.Logger != nil {
		c.Logger.QueryResponse(resBodyJSON)
	}

	resBody := struct {
		Data   interface{}
		Errors []Error
	}{
		Data: dest,
	}
	if err := json.Unmarshal(resBodyJSON, &resBody); err != nil {
		return fmt.Errorf("unmarshaling response body: %w", err)
	}

	if len(resBody.Errors) != 0 {
		return &resBody.Errors[0]
	}

	return nil
}
