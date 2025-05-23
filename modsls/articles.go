package modsls

// 品质
type Quality string

const (
	Red    Quality = "红"
	Gold   Quality = "金"
	Purple Quality = "紫"
	Blue   Quality = "蓝"
	Green  Quality = "绿"
	White  Quality = "白"
)

// 获取品质对应的权重
func GetQualityWeight(q Quality) float64 {
	switch q {
	case Red:
		return 0.01 // 1%
	case Gold:
		return 0.04 // 4%
	case Purple:
		return 0.15 // 15%
	case Blue:
		return 0.25 // 25%
	case Green:
		return 0.30 // 30%
	case White:
		return 0.25 // 25%
	default:
		return 0
	}
}

// 空间占用格数
type Space struct {
	Width  int `gorm:"not null"`
	Height int `gorm:"not null"`
}

// 物品的基础属性
type Articles struct {
	ID      int     `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name" gorm:"not null"`
	Quality Quality `json:"quality" gorm:"not null"`
	Space   Space   `json:"space" gorm:"embedded"` // 嵌入Space结构体
	Weight  float64 `json:"weight"`
	Price   int     `json:"price" gorm:"null"`
	Type    string  `json:"type" gorm:"not null"`
}
