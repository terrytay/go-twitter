package entities

type User struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	HashedPassword string `json:"-"`
}

func (u User) GetId() string {
	return u.Id
}

type NewUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
