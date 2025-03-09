package jsonutil

type Movie struct {
	Title, Subtitle *string
	Year            []*int32
	TargetFans      []*TargetFan
	Color           *bool
	Star            map[string]string
	Actors          []*map[string]string
	Oscars          []*string
	Sequel          *string
}

type TargetFan struct {
	// 职业类型
	WorkType *string
	// 典型人物信息
	ExampleInfo []*ExampleInfo
}

type ExampleInfo struct {
	// 姓名
	Name string
	// 年龄
	Age *int64
	// 上一次消费
	LastCost *float32
	// 平均消费
	AvgCost *float64
}

var defaultTestMapConfig = map[string]interface{}{
	"Title":    "this is a Title",
	"Subtitle": "this is a Subtitle",
	"Year":     2025,
	"Star": map[string]string{
		"LiWei": "年度最佳男主角 LiWei, 曾担任多次金鸡奖候选人, 如今终于在这部影片中星途璀璨, 为羊村的转速拼搏出挖掘机的纽带",
	},
	"Actors": []*map[string]string{
		{
			"Alice": "女主",
		},
		{
			"Fly": "路人甲",
		},
	},
	"ExampleInfo": ExampleInfo{
		Name: "this is a Name",
	},
	"Age":      18,
	"LastCost": 33141.123,
	"AvgCost":  12354.7217361,
}
