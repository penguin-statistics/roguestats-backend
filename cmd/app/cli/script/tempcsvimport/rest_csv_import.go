package tempcsvimport

import (
	"log"
	"strings"
)

type RestCSVImport struct {
	path string
}

func NewRestCSVImport(path string) *RestCSVImport {
	return &RestCSVImport{
		path: path,
	}
}

func (c *RestCSVImport) Run() error {
	records := ReadCSVFile(c.path)
	records = records[1:]
	for _, row := range records {
		log.Default().Printf("importing row '%s'\n", row)
		content := c.convertRowToContent(row)
		userID := c.getUserID(row)
		log.Println(content)
		if len(content) > 0 && userID != "" {
			PostEvent(content, "rsc_01h8yfh5yg0wergmjhewagsrq2", userID)
		}
	}
	return nil
}

func (c *RestCSVImport) convertRowToContent(row []string) map[string]any {
	content := make(map[string]any)
	columnHandler := GetColumnHandler()

	grade := columnHandler.HandleInt(strings.TrimSpace(row[1]))
	if grade.Valid {
		content["grade"] = grade.Int64
	}

	floor := columnHandler.HandleInt(strings.TrimSpace(row[2]))
	if floor.Valid {
		content["floor"] = floor.Int64
	}

	restChoices := columnHandler.HandleRestChoices(strings.TrimSpace(row[3]))
	if restChoices != nil {
		content["restChoices"] = restChoices
	}

	return content
}

func (c *RestCSVImport) getUserID(row []string) string {
	return GetColumnHandler().HandleUser(strings.TrimSpace(row[4]))
}
