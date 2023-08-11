package script

import (
	"log"
	"strings"
)

type BattleCSVImport struct {
	path string
}

func NewBattleCSVImport(path string) *BattleCSVImport {
	return &BattleCSVImport{
		path: path,
	}
}

func (c *BattleCSVImport) Run() error {
	records := ReadCSVFile(c.path)
	records = records[1:]
	for _, row := range records {
		log.Default().Printf("importing row '%s'\n", row)
		content := c.convertRowToContent(row)
		log.Println(content)
		if len(content) > 0 {
			PostEvent(content, "battle")
		}
	}
	return nil
}

func (c *BattleCSVImport) convertRowToContent(row []string) map[string]interface{} {
	content := make(map[string]interface{})
	columnHandler := GetColumnHandler()

	band := columnHandler.HandleBand(strings.TrimSpace(row[1]))
	if band != "" {
		content["band"] = band
	}

	grade := columnHandler.HandleInt(strings.TrimSpace(row[2]))
	if grade.Valid {
		content["grade"] = grade.Int64
	}

	floor := columnHandler.HandleInt(strings.TrimSpace(row[3]))
	if floor.Valid {
		content["floor"] = floor.Int64
	}

	isPortal := columnHandler.HandleBool(strings.TrimSpace(row[4]))
	if isPortal.Valid {
		content["isPortal"] = isPortal.Bool
	}

	nodeType := columnHandler.HandleNodeType(strings.TrimSpace(row[5]))
	if nodeType != "" {
		content["nodeType"] = nodeType
	}

	visionBeforeBattle := columnHandler.HandleInt(strings.TrimSpace(row[6]))
	if visionBeforeBattle.Valid {
		content["visionBeforeBattle"] = visionBeforeBattle.Int64
	}

	isPerfect := columnHandler.HandleBool(strings.TrimSpace(row[7]))
	if isPerfect.Valid {
		content["isPerfect"] = isPerfect.Bool
	}

	dropGold := columnHandler.HandleInt(strings.TrimSpace(row[8]))
	if dropGold.Valid {
		content["dropGold"] = dropGold.Int64
	}

	dropVision := columnHandler.HandleInt(strings.TrimSpace(row[9]))
	if dropVision.Valid {
		content["dropVision"] = dropVision.Int64
	}

	dropRecruitTickets := columnHandler.HandleRecruitTickets(strings.TrimSpace(row[10]), strings.TrimSpace(row[11]))
	if dropRecruitTickets != nil {
		content["dropRecruitTickets"] = dropRecruitTickets
	}

	dropTotem := columnHandler.HandleTotems(strings.TrimSpace(row[12]))
	if dropTotem != nil {
		content["dropTotem"] = dropTotem
	}

	gainExp := columnHandler.HandleInt(strings.TrimSpace(row[13]))
	if gainExp.Valid {
		content["gainExp"] = gainExp.Int64
	}

	return content
}
