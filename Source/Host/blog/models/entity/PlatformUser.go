package entity

type PlatformUser struct {
	ID          int    `json:"id",gorm:"primary_key"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Name        string `json:"name"`
	DpPath      string `json:"dp_path"`
	Description string `json:"description"`
	UserType    int    `json:"user_type"`
}

func (PlatformUser) TableName() string {
	return "platform_user"
}
