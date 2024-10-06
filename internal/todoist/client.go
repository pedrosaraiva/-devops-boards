package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	APIKey  string
	BaseURL string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey, BaseURL: "https://api.todoist.com/sync/v9"}
}

func (c *Client) GetCompletedTasks(ctx context.Context, projectID string, limit, offset int, until, since string, annotateNotes, annotateItems bool) (*CompletedTasksResponse, error) {
	completedTaskURL := "completed/get_all"
	params := url.Values{}
	if projectID != "" {
		params.Add("project_id", projectID)
	}
	if limit > 0 {
		params.Add("limit", strconv.Itoa(limit))
	}
	if offset > 0 {
		params.Add("offset", strconv.Itoa(offset))
	}
	if until != "" {
		params.Add("until", until)
	}
	if since != "" {
		params.Add("since", since)
	}
	params.Add("annotate_notes", strconv.FormatBool(annotateNotes))
	params.Add("annotate_items", strconv.FormatBool(annotateItems))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s?%s", c.BaseURL, completedTaskURL, params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var completedTasksResponse CompletedTasksResponse
	if err = json.NewDecoder(resp.Body).Decode(&completedTasksResponse); err != nil {
		return nil, err
	}

	return &completedTasksResponse, nil
}

func (c *Client) AddTask(task Task) (*CreateTaskResponse, error) {
	addTaskURL := "sync"
	command := Command{
		Type:   "item_add",
		TempID: "43f7ed23-a038-46b5-b2c9-4abda9097ffa",
		UUID:   "997d4b43-55f1-48a9-9e66-de5785dfd69b",
		Args:   task,
	}

	requestBody := Request{
		Commands: []Command{command},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.BaseURL, addTaskURL), bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req = req.WithContext(context.Background())

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var response CreateTaskResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
