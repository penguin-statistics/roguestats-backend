package exprutils

import (
	"sync"
)

type ExprCommonData struct {
	recruitTicketOperatorClassMap map[string][]string
	incidentTypeNameMap           map[string]string
	restChoicesNameMap            map[string]string
	variationNameMap              map[string]string
}

var exprCommonDataInstance *ExprCommonData
var exprCommonDataOnce sync.Once

func GetExprCommonData() *ExprCommonData {
	exprCommonDataOnce.Do(func() {
		exprCommonDataInstance = &ExprCommonData{
			recruitTicketOperatorClassMap: initRecruitTicketOperatorClassMap(),
			incidentTypeNameMap:           initIncidentTypeNameMap(),
			restChoicesNameMap:            initRestChoicesNameMap(),
			variationNameMap:              initVariationNameMap(),
		}
	})
	return exprCommonDataInstance
}

func (e ExprCommonData) GetRecruitTicketOperatorClassMap() map[string][]string {
	return e.recruitTicketOperatorClassMap
}

func (e ExprCommonData) GetIncidentTypeNameMap() map[string]string {
	return e.incidentTypeNameMap
}

func (e ExprCommonData) GetRestChoicesNameMap() map[string]string {
	return e.restChoicesNameMap
}

func (e ExprCommonData) GetVariationNameMap() map[string]string {
	return e.variationNameMap
}

func initRecruitTicketOperatorClassMap() map[string][]string {
	recruitTicketOperatorClassMap := make(map[string][]string)
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_pioneer"] = []string{"先锋"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_warrior"] = []string{"近卫"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_tank"] = []string{"重装"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_sniper"] = []string{"狙击"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_caster"] = []string{"术师"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_support"] = []string{"辅助"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_medic"] = []string{"医疗"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_special"] = []string{"特种"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_pioneer"] = []string{"先锋"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_double_1"] = []string{"先锋", "近卫"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_double_2"] = []string{"重装", "辅助"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_double_3"] = []string{"狙击", "医疗"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_double_4"] = []string{"术师", "特种"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_quad_melee"] = []string{"先锋", "近卫", "重装", "特种"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_quad_ranged"] = []string{"辅助", "狙击", "医疗", "术师"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_quad_melee_discount"] = []string{"先锋", "近卫", "重装", "特种"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_quad_ranged_discount"] = []string{"辅助", "狙击", "医疗", "术师"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_all"] = []string{"先锋", "近卫", "重装", "辅助", "狙击", "医疗", "术师", "特种"}
	recruitTicketOperatorClassMap["rogue_3_recruit_ticket_all_discount"] = []string{"先锋", "近卫", "重装", "辅助", "狙击", "医疗", "术师", "特种"}
	return recruitTicketOperatorClassMap
}

func initIncidentTypeNameMap() map[string]string {
	return map[string]string{
		"ro3_res1":    "吉兆",
		"ro3_res2":    "雨！",
		"ro3_res3":    "随到随取",
		"ro3_res4":    "定期维护",
		"ro3_res5":    "昏黑之室",
		"ro3_res2a":   "疗愈仪式",
		"ro3_res3a":   "黑色足迹",
		"ro3_res5a":   "不冻河",
		"ro3_rec1":    "野外生存专家",
		"ro3_rec2":    "在地公共信号放大站756号",
		"ro3_relic1":  "特里蒙旅行社特派团",
		"ro3_relic2":  "在地公共移动站N6号",
		"ro3_spring":  "度假胜地",
		"ro3_height":  "负伤的主树",
		"ro3_pick1":   "不见群山",
		"ro3_pick2":   "远见所向",
		"ro3_normal1": "萨米之语",
		"ro3_normal2": "乌萨斯",
		"ro3_normal3": "无用之物",
		"ro3_bat1":    "沼泽里的抽泣声",
		"ro3_bat2":    "有利可图",
		"ro3_bat3":    "随行人员",
		"ro3_bat4":    "邪恶计划鸭",
		"ro3_bat5":    "百里连营",
		"ro3_bat6":    "请君入戏",
		"ro3_bat7":    "北风女巫",
		"ro3_bat8":    "时刻警惕",
		"ro3_bat9":    "猜疑链",
	}
}

func initRestChoicesNameMap() map[string]string {
	return map[string]string{
		"ro3_rest_1": "稍事休息",
		"ro3_rest_2": "强化训练",
		"ro3_rest_3": "鼓舞士气",
		"ro3_rest_4": "重整编制",
		"ro3_rest_5": "维护设备",
		"ro3_rest_6": "求得启示",
	}
}

func initVariationNameMap() map[string]string {
	return map[string]string{
		"variation_1":       "己方生命攻击提升，出现国度",
		"variation_2":       "攻击下降，距离2伤害提升",
		"variation_3":       "同时部署人数、再部署减少",
		"variation_4":       "技力消耗降低，专注失调",
		"variation_5":       "敌方移速下降，重量防御法抗提升",
		"variation_6":       "生命上限提升，会掉更多血",
		"variation_shop":    "立即获得一笔资金，暗藏低价商品",
		"variation_shelter": "每前进一步都能获得希望",
	}
}
