package belajar

import "fmt"

// Encapsulation
type user struct {
	name       string
	Age        int
	Profession string
	Office     string
}

func (u user) getName() string {
	return u.name
}

func (u user) GetUser() string {
	return fmt.Sprintf("Nama saya %s, umur saya %d, pekerjaan %s, bekerja di %s",
		u.getName(), u.Age, u.Profession, u.Office)
}

func NewUser(name string, age int, proffession string, office string) user {
	return user{name: name, Age: age, Profession: proffession, Office: office}
}
