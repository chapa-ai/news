package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"news/pkg/apiNews"
	"news/pkg/db"
	"news/pkg/models"
)

var sem = make(chan bool, 2)
var log = logrus.StandardLogger()

func News(w http.ResponseWriter, r *http.Request) {

	sem <- true

	buyerLog := log.WithContext(r.Context()).WithField("News", gofakeit.Phrase())
	buyerLog.Info("One more fresh news came")

	var params models.Params

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		models.Error(w, 400, "failed decode")
		return
	}

	news, err := apiNews.FetchNews(r.Context(), &params)
	if err != nil {
		models.Error(w, 400, "fetchNews failed")
		return
	}

	data, err := apiNews.NewData(r.Context(), news)
	if err != nil {
		models.Error(w, 500, "newData failed")
		return
	}

	g, ctx := errgroup.WithContext(r.Context())

	for _, val := range data {
		value := val

		g.Go(func() error {

			similarNews, err := apiNews.FetchSimilarNews(ctx, value.Uuid)
			if err != nil {
				return fmt.Errorf("fetchSimilarNews failed: %w", err)
			}

			value.SimilarNews, err = apiNews.NewData(ctx, similarNews)
			if err != nil {
				return fmt.Errorf("fetchSimilarNews failed: %w", err)
			}

			return err
		})

	}

	err = g.Wait()
	if err != nil {
		models.Error(w, 400, "couldn't get similar new(s)")
		return
	}

	err = db.InsertData(r.Context(), data)
	if err != nil {
		log.Printf("Error: %v", err)
		models.Error(w, 400, fmt.Sprintf("insertData failed: %v", err))
		return
	}

	j, err := json.Marshal(data)
	if err != nil {
		models.Error(w, 400, "json Marshalling failed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(j)
	if err != nil {
		models.Error(w, 400, "w.Write failed")
		return
	}

	buyerLog.Info("This fresh news is gone")
	<-sem

}
