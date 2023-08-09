package exprutils

import (
	"sync"

	"github.com/rs/zerolog/log"
)

type ExprCommonData struct {
	totemValueMap map[string]int
}

var exprCommonDataInstance *ExprCommonData
var exprCommonDataOnce sync.Once

func GetExprCommonData() *ExprCommonData {
	exprCommonDataOnce.Do(func() {
		exprCommonDataInstance = &ExprCommonData{
			totemValueMap: initTotemValueMap(),
		}
	})
	return exprCommonDataInstance
}

func (e ExprCommonData) GetTotemValueMap() map[string]int {
	return e.totemValueMap
}

func initTotemValueMap() map[string]int {
	log.Debug().Msg("Initializing totem value map")
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
