package exprutils

import (
	"sync"
)

type ExprCommonData struct {
	totemValueMap                 map[string]int
	recruitTicketOperatorClassMap map[string][]string
}

var exprCommonDataInstance *ExprCommonData
var exprCommonDataOnce sync.Once

func GetExprCommonData() *ExprCommonData {
	exprCommonDataOnce.Do(func() {
		exprCommonDataInstance = &ExprCommonData{
			totemValueMap:                 initTotemValueMap(),
			recruitTicketOperatorClassMap: initRecruitTicketOperatorClassMap(),
		}
	})
	return exprCommonDataInstance
}

func (e ExprCommonData) GetTotemValueMap() map[string]int {
	return e.totemValueMap
}

func (e ExprCommonData) GetRecruitTicketOperatorClassMap() map[string][]string {
	return e.recruitTicketOperatorClassMap
}

func initTotemValueMap() map[string]int {
	totemValueMap := make(map[string]int)
	totemValueMap["rogue_3_totem_R_L1"] = 8
	totemValueMap["rogue_3_totem_R_L2"] = 8
	totemValueMap["rogue_3_totem_R_L3"] = 8
	totemValueMap["rogue_3_totem_R_L4"] = 12
	totemValueMap["rogue_3_totem_R_L5"] = 12
	totemValueMap["rogue_3_totem_R_L6"] = 16
	totemValueMap["rogue_3_totem_R_E1"] = 8
	totemValueMap["rogue_3_totem_R_E2"] = 8
	totemValueMap["rogue_3_totem_R_E3"] = 8
	totemValueMap["rogue_3_totem_R_E4"] = 12
	totemValueMap["rogue_3_totem_R_E5"] = 12
	totemValueMap["rogue_3_totem_R_E6"] = 16
	totemValueMap["rogue_3_totem_G_L1"] = 8
	totemValueMap["rogue_3_totem_G_L2"] = 8
	totemValueMap["rogue_3_totem_G_L3"] = 8
	totemValueMap["rogue_3_totem_G_L4"] = 12
	totemValueMap["rogue_3_totem_G_L5"] = 12
	totemValueMap["rogue_3_totem_G_L6"] = 16
	totemValueMap["rogue_3_totem_G_E1"] = 8
	totemValueMap["rogue_3_totem_G_E2"] = 8
	totemValueMap["rogue_3_totem_G_E3"] = 8
	totemValueMap["rogue_3_totem_G_E4"] = 12
	totemValueMap["rogue_3_totem_G_E5"] = 12
	totemValueMap["rogue_3_totem_G_E6"] = 16
	totemValueMap["rogue_3_totem_B_L1"] = 8
	totemValueMap["rogue_3_totem_B_L2"] = 8
	totemValueMap["rogue_3_totem_B_L3"] = 8
	totemValueMap["rogue_3_totem_B_L4"] = 12
	totemValueMap["rogue_3_totem_B_L5"] = 12
	totemValueMap["rogue_3_totem_B_L6"] = 16
	totemValueMap["rogue_3_totem_B_E1"] = 8
	totemValueMap["rogue_3_totem_B_E2"] = 8
	totemValueMap["rogue_3_totem_B_E3"] = 8
	totemValueMap["rogue_3_totem_B_E4"] = 12
	totemValueMap["rogue_3_totem_B_E5"] = 12
	totemValueMap["rogue_3_totem_B_E6"] = 16
	return totemValueMap
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
