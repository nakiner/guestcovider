package userRepository

type User struct {
	ID           uint64 `gorm:"primary_key"`
	Status       string
	Company      string
	Surname      string
	Name         string
	Guest        string
	CovidPass    string
	Rank         string
	ContactPhone string
	ContactMail  string
	Checkin      bool
}

func (User) TableName() string {
	return "users"
}
