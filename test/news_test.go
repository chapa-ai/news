package test

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/sync/errgroup"
	"news/pkg/models"
	"testing"
)

var (
	url1 = "http://localhost:9993"
)

func TestNews(t *testing.T) {
	params := &models.Params{
		Limit:      5,
		Categories: "tech",
		Language:   "en",
	}
	t.Parallel()

	_, err := News(params)
	if err != nil {
		t.Fatal(err)
	}

}

func News(params *models.Params) ([]*models.MainData, error) {
	output := []*models.MainData{}

	url := fmt.Sprintf("%s/news", url1)

	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&output).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("status code wrong. status: %v. body: %v", resp.StatusCode(), resp.String())
	}

	return output, nil

}

func TestSemaphoreNews(t *testing.T) {
	t.Parallel()

	params := &models.Params{
		Limit:      5,
		Categories: "tech",
		Language:   "en",
	}

	g, _ := errgroup.WithContext(context.Background())

	for i := 0; i < 15; i++ {
		i = i

		g.Go(func() error {
			_, err := SemaphoreNews(params)
			if err != nil {
				t.Fatal(err)
			}

			return err
		})

	}

	err := g.Wait()
	if err != nil {
		t.Fatal(err)
		return
	}

}

func SemaphoreNews(params *models.Params) ([]*models.MainData, error) {
	output := []*models.MainData{}

	url := fmt.Sprintf("%s/news", url1)

	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		SetResult(&output).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("status code wrong. status: %v. body: %v", resp.StatusCode(), resp.String())
	}

	return output, nil

}
