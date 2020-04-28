package models

type Community struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestCommunity() *Community {
	return &Community{
		Name: "Some community",
	}
}

func (m *Community) Validate() error {
	return nil
}
