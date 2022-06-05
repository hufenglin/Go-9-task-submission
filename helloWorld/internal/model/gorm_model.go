package model

type User struct {
	ID int32
	Class int32
	Level int32
	Name string
}

func (User) TableName() string {
	return "User"
}