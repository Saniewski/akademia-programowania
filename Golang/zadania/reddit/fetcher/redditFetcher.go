package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditFetcher interface {
	Fetch(ctx context.Context) error
	Save(io.Writer) error
}

type Fetcher struct {
	response response
	URL      string
}

func (g *Fetcher) Fetch(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, g.URL, nil)
	if err != nil {
		return fmt.Errorf("error while creating request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error while fetching data: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error while fetching data: %v", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&g.response); err != nil {
		return fmt.Errorf("error while decoding data: %v", err)
	}

	return nil
}

func (g *Fetcher) Save(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%s\n\nPOSTS:\n", time.Now())
	if err != nil {
		return fmt.Errorf("error while saving data: %v", err)
	}
	for _, post := range g.response.Data.Children {
		_, err := fmt.Fprintf(w, "%s\n%s\n", post.Data.Title, post.Data.URL)
		if err != nil {
			return fmt.Errorf("error while saving data: %v", err)
		}
	}

	return nil
}

func NewFetcher(url string) *Fetcher {
	return &Fetcher{
		URL: url,
	}
}
