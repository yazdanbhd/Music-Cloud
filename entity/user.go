package entity

type User struct {
	ID          uint
	PhoneNumber string
	Name        string
	UserName    string
	Password    string
	PlayList    []Music
}
