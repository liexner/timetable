package togglclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"timetable/internal/types"
)

func (c *Client) GetProjects(ctx context.Context, token string) ([]types.Project, error) {
	url := c.BaseURL + "/me/projects"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", encodeToken(token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	var projects []types.Project
	if err := json.Unmarshal(body, &projects); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	return projects, nil
}

func (c *Client) GetTimeEntries(ctx context.Context, token string) ([]types.TimeEntry, error) {
	endDate := time.Now().AddDate(0, 0, 1)
	startDate := endDate.AddDate(0, 0, -16)

	url := fmt.Sprintf("%s/me/time_entries?start_date=%s&end_date=%s",
		c.BaseURL,
		startDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", encodeToken(token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	var entries []types.TimeEntry
	if err := json.Unmarshal(body, &entries); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	return entries, nil
}
