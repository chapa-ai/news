package apiNews

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"news/pkg/models"
	"os"
)

func FetchNews(ctx context.Context, p *models.Params) (*models.InternalNews, error) {
	path, err := url.JoinPath(fmt.Sprintf("https://api.thenewsapi.com/v1/news/top?api_token=%s&locale=us&limit=%d&categories=%s&language=%s", os.Getenv("apiToken"), p.Limit, p.Categories, p.Language))
	if err != nil {
		return nil, fmt.Errorf("failed JoinPath: %w", err)
	}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed NewRequest: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("couldn't make request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code not 200: %w", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read body: %w", err)
	}

	var userStat models.InternalNews
	err = json.Unmarshal(body, &userStat)
	if err != nil {
		return nil, fmt.Errorf("failed Unmarshal: %w", err)
	}

	return &userStat, nil
}

func FetchSimilarNews(ctx context.Context, uuid string) (*models.InternalNews, error) {
	path, err := url.JoinPath(fmt.Sprintf("https://api.thenewsapi.com/v1/news/similar/%s?api_token=%s", uuid, os.Getenv("apiToken")))
	if err != nil {
		return nil, fmt.Errorf("failed JoinPath: %w", err)
	}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed newRequest: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("make NewRequest failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code not 200: %w", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read body failed: %w", err)
	}

	var userStat *models.InternalNews
	err = json.Unmarshal(body, &userStat)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}

	return userStat, nil
}

func NewData(ctx context.Context, news *models.InternalNews) ([]*models.MainData, error) {
	list := make([]*models.MainData, len(news.Data))

	for key, value := range news.Data {

		result := models.MainData{
			Uuid:        value.UUID,
			Headline:    value.Title,
			Description: value.Description,
			Keywords:    value.Keywords,
			Snippet:     value.Snippet,
			Url:         value.URL,
		}
		list[key] = &result

	}

	return list, nil
}
