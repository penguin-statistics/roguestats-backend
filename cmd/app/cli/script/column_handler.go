package script

import (
	"strconv"
	"strings"
	"sync"

	"gopkg.in/guregu/null.v3"
)

type ColumnHandler struct {
	bandMap          map[string]string
	nodeTypeMap      map[string]string
	recruitTicketMap map[string]string
	totemMap         map[string]string
	incidentTypeMap  map[string]string
	variationMap     map[string]string
	layoutMap        map[string]string
}

var columnHandlerInstance *ColumnHandler
var columnHanlderOnce sync.Once

func GetColumnHandler() *ColumnHandler {
	columnHanlderOnce.Do(func() {
		columnHandlerInstance = &ColumnHandler{
			bandMap:          initBandMap(),
			nodeTypeMap:      initNodeTypeMap(),
			recruitTicketMap: initRecruitTicketMap(),
			totemMap:         initTotemMap(),
			incidentTypeMap:  initIncidentTypeMap(),
			variationMap:     initVariaionMap(),
			layoutMap:        initLayoutMap(),
		}
	})
	return columnHandlerInstance
}

func (c *ColumnHandler) HandleBand(input string) string {
	val, ok := c.bandMap[input]
	if !ok {
		return ""
	}
	return val
}

func (c *ColumnHandler) HandleInt(input string) null.Int {
	num, err := strconv.Atoi(input)
	if err != nil {
		return null.IntFromPtr(nil)
	}
	return null.IntFrom(int64(num))
}

func (c *ColumnHandler) HandleBool(input string) null.Bool {
	if input == "是" {
		return null.BoolFrom(true)
	}
	if input == "否" {
		return null.BoolFrom(false)
	}
	return null.BoolFromPtr(nil)
}

func (c *ColumnHandler) HandleNodeType(input string) string {
	val, ok := c.nodeTypeMap[input]
	if !ok {
		return ""
	}
	return val
}

func (c ColumnHandler) HandleRecruitTickets(input1 string, input2 string) [][]string {
	if input1 == "" {
		return nil
	}
	tickets := make([][]string, 2)
	tickets[0] = c.handleRecruitTicketsHelper(input1)
	tickets[1] = c.handleRecruitTicketsHelper(input2)
	return tickets
}

func (c ColumnHandler) handleRecruitTicketsHelper(input string) []string {
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

func (c ColumnHandler) HandleTotems(input string) []string {
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

func (c ColumnHandler) HandleIncidentType(input string) string {
	val, ok := c.incidentTypeMap[input]
	if !ok {
		return ""
	}
	return val
}

func (c ColumnHandler) HandleRestChoices(input string) []string {
	if input == "" {
		return nil
	}
	choices := make([]string, 3)
	strs := strings.Split(input, ",")
	for i, str := range strs {
		str = strings.TrimSpace(str)
		val, ok := initRestChoicesMap()[str]
		if !ok {
			continue
		}
		choices[i] = val
	}
	return choices
}

func (c ColumnHandler) HandleVariation(input string) string {
	val, ok := c.variationMap[input]
	if !ok {
		return ""
	}
	return val
}

func (c ColumnHandler) HandleLayout(input string) string {
	val, ok := c.layoutMap[input]
	if !ok {
		return ""
	}
	return val
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

func initIncidentTypeMap() map[string]string {
	return map[string]string{
		"吉兆":            "ro3_res1",
		"雨！":            "ro3_res2",
		"随到随取":          "ro3_res3",
		"定期维护":          "ro3_res4",
		"昏黑之室":          "ro3_res5",
		"疗愈仪式":          "ro3_res2a",
		"黑色足迹":          "ro3_res3a",
		"不冻河":           "ro3_res5a",
		"野外生存专家":        "ro3_rec1",
		"在地公共信号放大站756号": "ro3_rec2",
		"特里蒙旅行社特派团":     "ro3_relic1",
		"在地公共移动站N6号":    "ro3_relic2",
		"度假胜地":          "ro3_spring",
		"负伤的主树":         "ro3_height",
		"不见群山":          "ro3_pick1",
		"远见所向":          "ro3_pick2",
		"萨米之语":          "ro3_normal1",
		"乌萨斯":           "ro3_normal2",
		"无用之物":          "ro3_normal3",
		"沼泽里的抽泣声":       "ro3_bat1",
		"有利可图":          "ro3_bat2",
		"随行人员":          "ro3_bat3",
		"邪恶计划鸭":         "ro3_bat4",
		"百里连营":          "ro3_bat5",
		"请君入戏":          "ro3_bat6",
		"北风女巫":          "ro3_bat7",
		"时刻警惕":          "ro3_bat8",
		"猜疑链":           "ro3_bat9",
	}
}

func initRestChoicesMap() map[string]string {
	return map[string]string{
		"3生命上限":   "ro3_rest_1",
		"高级物资配给券": "ro3_rest_2",
		"3希望":     "ro3_rest_3",
		"可携带干员+1": "ro3_rest_4",
		"抗干扰+1":   "ro3_rest_5",
		"获得密文板":   "ro3_rest_6",
	}
}

func initVariaionMap() map[string]string {
	return map[string]string{
		"己方生命攻击提升，出现国度":   "variation_1",
		"攻击下降，距离2伤害提升":    "variation_2",
		"同时部署人数、再部署减少":    "variation_3",
		"技力消耗降低，专注失调":     "variation_4",
		"敌方移速下降，重量防御法抗提升": "variation_5",
		"生命上限提升，会掉更多血":    "variation_6",
		"立即获得一笔资金，暗藏低价商品": "variation_shop",
		"每前进一步都能获得希望":     "variation_shelter",
	}
}

func initLayoutMap() map[string]string {
	return map[string]string{
		"战斗c：3-2":            "battle_3-2",
		"战斗d：2-3":            "battle_2-3",
		"战斗e：3-3":            "battle_3-3",
		"事件a：2-2-2":          "event_2-2-2",
		"事件b：3-3（6商店）":       "event_3-3",
		"事件c：2-2（不期而遇+失与得）":  "event_2-2",
		"事件d: 2-3-2 （完全无战斗）": "event_2-3-2",
		"混合a：2-2-1":          "mixed_2-2-1",
		"混合b：2-3-1":          "mixed_2-3-1",
		"混合c：2-3-2":          "mixed_2-3-2",
		"混合d：2-2-2":          "mixed_2-2-2",
		"混合e：3-1":            "mixed_3-1",
	}
}
