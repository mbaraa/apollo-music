package enums

import "strings"

// Plan defines the subscription plan size
// Free is a subscription with size of 500MB
// 20GB well, it's 20GB
// 100GB same as 20GB, but it's a 100GB
type Plan string

const (
	FreePlan       Plan = "FREE"
	TwentyGBsPlan  Plan = "20GB"
	HundredGBsPlan Plan = "100GB"
)

func GetPlan(planText string) Plan {
	switch strings.ToUpper(planText) {
	case string(TwentyGBsPlan):
		return TwentyGBsPlan
	case string(HundredGBsPlan):
		return HundredGBsPlan
	case string(FreePlan):
		fallthrough
	default:
		return FreePlan
	}
}

func (p Plan) Size() int64 {
	return planSize[p]
}

func (p Plan) PlanStripeProductId() string {
	return planStripeProductId[p]
}

var planSize = map[Plan]int64{
	FreePlan:       512,
	TwentyGBsPlan:  20480,
	HundredGBsPlan: 102400,
}

var planStripeProductId = map[Plan]string{
	FreePlan:       "",
	TwentyGBsPlan:  "",
	HundredGBsPlan: "",
}
