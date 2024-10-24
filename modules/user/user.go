package user

type User struct {
	ID       string `gorm:"primaryKey;autoIncrement:true"`
	Email    string
	Password string
}
