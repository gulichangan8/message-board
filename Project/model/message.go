package model

type Mess struct {
	OwnerName string
	WriteName string
	Message   string
	Respond   string
}

type Messes []Mess

type Mes struct {
	Mess struct {
		OwnerName string
		WriteName string
		Message   string
		Respond   string
	}
}
