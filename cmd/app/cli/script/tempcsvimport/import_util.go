package tempcsvimport

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"github.com/machinebox/graphql"

	"exusiai.dev/roguestats-backend/internal/model"
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

func PostEvent(content map[string]any, researchID string, userID string) {
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
		UserID:     userID,
	}
	req.Var("input", input)
	req.Header.Set("Authorization", "Bearer <token>")
	ctx := context.Background()
	var respData any
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	log.Println(respData)
}
