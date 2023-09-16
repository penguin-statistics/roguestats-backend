package tempcsvimport

import (
	"log"
	"strings"
)

type PortalCSVImport struct {
	path string
}

func NewPortalCSVImport(path string) *PortalCSVImport {
	return &PortalCSVImport{
		path: path,
	}
}

func (c *PortalCSVImport) Run() error {
	records := ReadCSVFile(c.path)
	records = records[1:]
	for _, row := range records {
		log.Default().Printf("importing row '%s'\n", row)
		content := c.convertRowToContent(row)
		log.Println(content)
		if len(content) > 0 {
			PostEvent(content, "portal")
		}
	}
	return nil
}

func (c *PortalCSVImport) convertRowToContent(row []string) map[string]any {
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

	layout := columnHandler.HandleLayout(strings.TrimSpace(row[3]))
	if layout != "" {
		content["layout"] = layout
	}

	variation := columnHandler.HandleVariation(strings.TrimSpace(row[4]))
	if variation != "" {
		content["variation"] = variation
	}

	return content
}
