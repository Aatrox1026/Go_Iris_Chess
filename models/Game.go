package models

type Game struct {
	ID      int64  `gorm:"type:BIGINT NOT NULL AUTO_INCREMENT; PRIMARY_KEY;" json:"id"`
	Turns   int64  `gorm:"type:BIGINT NOT NULL;" json:"turns"`
	Status  string `gorm:"type:VARCHAR(10) NOT NULL;" json:"status"`
	Winner  string `gorm:"type:VARCHAR(10) NOT NULL;" json:"winner"`
	WhiteID string `gorm:"type:VARCHAR(30) NOT NULL;" json:"white_id"`
	BlackID string `gorm:"type:VARCHAR(30) NOT NULL;" json:"black_id"`
	Moves   string `gorm:"type:VARCHAR(3000) NOT NULL;" json:"moves"`
	ECOCode string `gorm:"type:VARCHAR(10) NOT NULL;" json:"eco_code"`
	Opening string `gorm:"type:VARCHAR(100) NOT NULL;" json:"opening"`
}
