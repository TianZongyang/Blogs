package entity

type SystemConfig struct {
	ID  int    `json:"id",gorm:"primary_key"`
	Key string `json:"key"`
	Val string `json:"val"`
}

func (SystemConfig) TableName() string {
	return "system_config"
}
