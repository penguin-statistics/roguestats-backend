package tempcsvimport

import (
	"log"
	"strings"
)

type IncidentCSVImport struct {
	path string
}

func NewIncidentCSVImport(path string) *IncidentCSVImport {
	return &IncidentCSVImport{
		path: path,
	}
}

func (c *IncidentCSVImport) Run() error {
	records := ReadCSVFile(c.path)
	records = records[1:]
	for _, row := range records {
		log.Default().Printf("importing row '%s'\n", row)
		content := c.convertRowToContent(row)
		userID := c.getUserID(row)
		log.Println(content)
		if len(content) > 0 && userID != "" {
			PostEvent(content, "rsc_01h8yfh5y9mf4ws7ehf1q9n9n5", userID)
		}
	}
	return nil
}

func (c *IncidentCSVImport) convertRowToContent(row []string) map[string]any {
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

	isPortal := columnHandler.HandleBool(strings.TrimSpace(row[3]))
	if isPortal.Valid {
		content["isPortal"] = isPortal.Bool
	}

	incidentType := columnHandler.HandleIncidentType(strings.TrimSpace(row[4]))
	if incidentType != "" {
		content["incidentType"] = incidentType
	}

	return content
}

func (c *IncidentCSVImport) getUserID(row []string) string {
	return GetColumnHandler().HandleUser(strings.TrimSpace(row[5]))
}
