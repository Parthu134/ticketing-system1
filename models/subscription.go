package models

type Subscription struct {
	ID       uint   `gorm:"primarykey"`
	Role     string `json:"role"`
	Endpoint string `json:"endpoint"`
	P256dh   string `json:"p256dh"`
	Auth     string `json:"auth"`
}
