package models

type Community struct {
	ID   int
	Name string
}

func TestCommunity() *Community {
	return &Community{
		Name: "Some community",
	}
}
