package script

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"exusiai.dev/roguestats-backend/internal/model"
	"github.com/machinebox/graphql"
)

func ReadCSVFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}

func PostEvent(content map[string]any, researchID string) {
	client := graphql.NewClient("http://localhost:3500/graphql")
	req := graphql.NewRequest(`
	mutation CreateEvent($input: CreateEventInput!) {
		createEvent(input: $input) {
		  content
		}
	  }`,
	)
	userAgent := "cli"
	input := model.CreateEventInput{
		Content:    content,
		ResearchID: researchID,
		UserAgent:  userAgent,
	}
	req.Var("input", input)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJFUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJyb2d1ZXN0YXRzIiwiZXhwIjoxNjk0MjA0MjY1LCJpYXQiOjE2OTI5OTQ2NjUsImlzcyI6InJvZ3Vlc3RhdHMvdjAuMC4wIiwibmJmIjoxNjkyOTk0NjY1LCJzdWIiOiIwMWg4cTVlYnJuNWV0aG0xcDZ6anhyOWVmdyJ9.AHlIYrx7tKj6nnXO4MYRd_0mXqzOVWPyG6FHidPitfI2IbrtZI3-lXA-bZP_nl0Op7d4TgzacdYwJPDgYGLoZcznAfopT-ahoHmDZrflhrK-Soo8ji7OZENjOIH5VetkkTaKl9zuqdAivds4DQPefSYngsn5vqzIgIZhaoR8nJoaq6MT")
	ctx := context.Background()
	var respData any
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	log.Println(respData)
}
