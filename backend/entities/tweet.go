package entities

type Tweet struct {
	Id        string
	Message   string
	Timestamp int64
}

func (t Tweet) GetId() string {
	return t.Id
}
