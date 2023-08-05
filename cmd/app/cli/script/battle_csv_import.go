package script

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/machinebox/graphql"
	"gopkg.in/guregu/null.v3"

	"exusiai.dev/roguestats-backend/internal/model"
)

type BattleCSVImport struct {
	path             string
	bandMap          map[string]string
	nodeTypeMap      map[string]string
	recruitTicketMap map[string]string
	totemMap         map[string]string
}

func NewBattleCSVImport(path string) *BattleCSVImport {
	return &BattleCSVImport{
		path:             path,
		bandMap:          initBandMap(),
		nodeTypeMap:      initNodeTypeMap(),
		recruitTicketMap: initRecruitTicketMap(),
		totemMap:         initTotemMap(),
	}
}

func (c *BattleCSVImport) Run() error {
	records := readCSVFile(c.path)
	records = records[1:]
	for _, row := range records {
		log.Default().Printf("importing row '%s'\n", row)
		content := c.convertRowToContent(row)
		log.Println(content)
		if len(content) > 0 {
			postEvent(content)
		}
	}
	return nil
}

func readCSVFile(filePath string) [][]string {
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

func postEvent(content map[string]interface{}) {
	client := graphql.NewClient("http://localhost:3500/graphql")
	req := graphql.NewRequest(`
	mutation CreateEvent($newEvent: NewEvent!) {
		createEvent(input: $newEvent) {
		  content
		}
	  }`,
	)
	userAgent := "cli"
	newEvent := model.NewEvent{
		Content:    content,
		ResearchID: "battle",
		UserAgent:  &userAgent,
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

func (c *BattleCSVImport) convertRowToContent(row []string) map[string]interface{} {
	content := make(map[string]interface{})

	band := c.handleBand(strings.TrimSpace(row[1]))
	if band != "" {
		content["band"] = band
	}

	grade := c.handleInt(strings.TrimSpace(row[2]))
	if grade.Valid {
		content["grade"] = grade.Int64
	}

	floor := c.handleInt(strings.TrimSpace(row[3]))
	if floor.Valid {
		content["floor"] = floor.Int64
	}

	isPortal := c.handleBool(strings.TrimSpace(row[4]))
	if isPortal.Valid {
		content["isPortal"] = isPortal.Bool
	}

	nodeType := c.handleNodeType(strings.TrimSpace(row[5]))
	if nodeType != "" {
		content["nodeType"] = nodeType
	}

	visionBeforeBattle := c.handleInt(strings.TrimSpace(row[6]))
	if visionBeforeBattle.Valid {
		content["visionBeforeBattle"] = visionBeforeBattle.Int64
	}

	isPerfect := c.handleBool(strings.TrimSpace(row[7]))
	if isPerfect.Valid {
		content["isPerfect"] = isPerfect.Bool
	}

	dropGold := c.handleInt(strings.TrimSpace(row[8]))
	if dropGold.Valid {
		content["dropGold"] = dropGold.Int64
	}

	dropVision := c.handleInt(strings.TrimSpace(row[9]))
	if dropVision.Valid {
		content["dropVision"] = dropVision.Int64
	}

	dropRecruitTickets := c.handleRecruitTickets(strings.TrimSpace(row[10]), strings.TrimSpace(row[11]))
	if dropRecruitTickets != nil {
		content["dropRecruitTickets"] = dropRecruitTickets
	}

	dropTotem := c.handleTotems(strings.TrimSpace(row[12]))
	if dropTotem != nil {
		content["dropTotem"] = dropTotem
	}

	gainExp := c.handleInt(strings.TrimSpace(row[13]))
	if gainExp.Valid {
		content["gainExp"] = gainExp.Int64
	}

	return content
}

func (c *BattleCSVImport) handleBand(input string) string {
	val, ok := c.bandMap[input]
	if !ok {
		return ""
	}
	return val
}

func (c *BattleCSVImport) handleInt(input string) null.Int {
	num, err := strconv.Atoi(input)
	if err != nil {
		return null.IntFromPtr(nil)
	}
	return null.IntFrom(int64(num))
}

func (c *BattleCSVImport) handleBool(input string) null.Bool {
	if input == "是" {
		return null.BoolFrom(true)
	}
	if input == "否" {
		return null.BoolFrom(false)
	}
	return null.BoolFromPtr(nil)
}

func (c *BattleCSVImport) handleNodeType(input string) string {
	val, ok := c.nodeTypeMap[input]
	if !ok {
		return ""
	}
	return val
}

func (c BattleCSVImport) handleRecruitTickets(input1 string, input2 string) [][]string {
	if input1 == "" {
		return nil
	}
	tickets := make([][]string, 2)
	tickets[0] = c.handleRecruitTicketsHelper(input1)
	tickets[1] = c.handleRecruitTicketsHelper(input2)
	return tickets
}

func (c BattleCSVImport) handleRecruitTicketsHelper(input string) []string {
	tickets := make([]string, 0)
	if input == "" {
		return tickets
	}
	strs := strings.Split(input, ",")
	for _, str := range strs {
		str = strings.TrimSpace(str)
		val, ok := c.recruitTicketMap[str]
		if !ok {
			continue
		}
		tickets = append(tickets, val)
	}
	return tickets
}

func (c BattleCSVImport) handleTotems(input string) []string {
	if input == "" {
		return nil
	}
	totems := make([]string, 0)
	if input == "无掉落" {
		return totems
	}
	strs := strings.Split(input, ",")
	for _, str := range strs {
		str = strings.TrimSpace(str)
		val, ok := c.totemMap[str]
		if !ok {
			continue
		}
		totems = append(totems, val)
	}
	return totems
}

func initBandMap() map[string]string {
	return map[string]string{
		"指挥":     "rogue_3_band_1",
		"集群":     "rogue_3_band_2",
		"后勤":     "rogue_3_band_3",
		"矛头":     "rogue_3_band_4",
		"突击（近锋）": "rogue_3_band_5",
		"堡垒（重辅）": "rogue_3_band_6",
		"远程（狙医）": "rogue_3_band_7",
		"破坏（术特）": "rogue_3_band_8",
		"特训":     "rogue_3_band_9",
		"高规格":    "rogue_3_band_10",
		"永恒狩猎":   "rogue_3_band_11",
		"生活至上":   "rogue_3_band_12",
		"科学主义":   "rogue_3_band_13",
	}
}

func initNodeTypeMap() map[string]string {
	return map[string]string{
		"作战":         "BATTLE_NORMAL",
		"紧急作战":       "BATTLE_ELITE",
		"险路恶敌":       "BATTLE_BOSS",
		"不期而遇":       "INCIDENT",
		"不期而遇-黑色足迹":  "INCIDENT_RES3A",
		"不期而遇-鸭爵与雇员": "INCIDENT_MIMIC_ENEMY",
		"失与得":        "SACRIFICE",
		"诡异行商":       "BATTLE_SHOP",
	}
}

func initRecruitTicketMap() map[string]string {
	return map[string]string{
		"先锋":       "rogue_3_recruit_ticket_pioneer",
		"近卫":       "rogue_3_recruit_ticket_warrior",
		"重装":       "rogue_3_recruit_ticket_tank",
		"狙击":       "rogue_3_recruit_ticket_sniper",
		"术师":       "rogue_3_recruit_ticket_caster",
		"辅助":       "rogue_3_recruit_ticket_support",
		"医疗":       "rogue_3_recruit_ticket_medic",
		"特种":       "rogue_3_recruit_ticket_special",
		"突击协议（近锋）": "rogue_3_recruit_ticket_double_1",
		"堡垒协议（重辅）": "rogue_3_recruit_ticket_double_2",
		"远程协议（狙医）": "rogue_3_recruit_ticket_double_3",
		"破坏协议（术特）": "rogue_3_recruit_ticket_double_4",
		"前线统合":     "rogue_3_recruit_ticket_quad_melee",
		"后方协调":     "rogue_3_recruit_ticket_quad_ranged",
		"前线统合资深":   "rogue_3_recruit_ticket_quad_melee_discount",
		"后方协调资深":   "rogue_3_recruit_ticket_quad_ranged_discount",
		"高级人事调度函":  "rogue_3_recruit_ticket_all",
		"高级人事资深":   "rogue_3_recruit_ticket_all_discount",
	}
}

func initTotemMap() map[string]string {
	return map[string]string{
		"黜人": "rogue_3_totem_R_L1",
		"猎手": "rogue_3_totem_R_L2",
		"战士": "rogue_3_totem_R_L3",
		"萨满": "rogue_3_totem_R_L4",
		"雪祀": "rogue_3_totem_R_L5",
		"英雄": "rogue_3_totem_R_L6",
		"歌唱": "rogue_3_totem_R_E1",
		"沉默": "rogue_3_totem_R_E2",
		"朗诵": "rogue_3_totem_R_E3",
		"辩论": "rogue_3_totem_R_E4",
		"慰藉": "rogue_3_totem_R_E5",
		"告解": "rogue_3_totem_R_E6",
		"树冠": "rogue_3_totem_G_L1",
		"水面": "rogue_3_totem_G_L2",
		"眼睛": "rogue_3_totem_G_L3",
		"拱门": "rogue_3_totem_G_L4",
		"光芒": "rogue_3_totem_G_L5",
		"大地": "rogue_3_totem_G_L6",
		"喜悦": "rogue_3_totem_G_E1",
		"惊讶": "rogue_3_totem_G_E2",
		"愤怒": "rogue_3_totem_G_E3",
		"疑惑": "rogue_3_totem_G_E4",
		"憧憬": "rogue_3_totem_G_E5",
		"爱恋": "rogue_3_totem_G_E6",
		"源石": "rogue_3_totem_B_L1",
		"乔木": "rogue_3_totem_B_L2",
		"砂石": "rogue_3_totem_B_L3",
		"灌木": "rogue_3_totem_B_L4",
		"兽类": "rogue_3_totem_B_L5",
		"人类": "rogue_3_totem_B_L6",
		"巡视": "rogue_3_totem_B_E1",
		"筑巢": "rogue_3_totem_B_E2",
		"捕猎": "rogue_3_totem_B_E3",
		"掠夺": "rogue_3_totem_B_E4",
		"繁衍": "rogue_3_totem_B_E5",
		"迁徙": "rogue_3_totem_B_E6",
	}
}
