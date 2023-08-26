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

func PostEvent(content map[string]interface{}, researchID string) {
	client := graphql.NewClient("http://localhost:3500/graphql")
	req := graphql.NewRequest(`
	mutation CreateEvent($newEvent: NewEvent!) {
		createEvent(input: $newEvent) {
		  content
		}
	  }`,
	)
	userAgent := "cli"
	newEvent := model.CreateEventInput{
		Content:    content,
		ResearchID: researchID,
		UserAgent:  userAgent,
	}
	req.Var("newEvent", newEvent)
	// req.Header.Set("Authorization", "")
	ctx := context.Background()
	var respData interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	log.Println(respData)
}
