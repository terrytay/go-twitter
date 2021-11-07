package entities

type User struct {
	Id       string
	Name     string
	Username string
	Tweets   []Tweet
}

func (u User) GetId() string {
	return u.Id
}
