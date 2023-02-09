package model

type Movement struct {
	Name      string
	Type      string
	Class     string
	Power     int
	Precision string
	Effect    string
}

var emptyMovement = &Movement{}

func (m *Movement) ResetStruct() {

	*m = *emptyMovement

}
