package hyperion

import (
	"github.com/yl10/mdd"
)

//DimScenario 情景维度
var DimScenario mdd.Dimension

// SCENARIO_ID NUMBER No FK(member).
// START_YR_ID NUMBER No FK(member). Start year of time period
// END_YR_ID NUMBER No FK(member) Ending year of time period
// START_TP_ID NUMBER No FK(time_period). Referenced time period
// defines the start period for the scenario.
// END_TP_ID NUMBER No FK(time_period). Referenced time period
// defines the end period for the scenario.
// FX_TBL NUMBER Yes FK(fx_table). Exchange rate table
// associated with scenario
// USEBEGBAL NUMBER Yes Flag to indicate whether to include
// BegBalance T
type Scenario struct {
	SCENARIO_ID int64
}
