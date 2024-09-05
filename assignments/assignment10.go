package main

import (
	"fmt"
	"time"
)

type institute interface{
	register(string, string) student
	list() string
	calculateAgeFromDOB(string, Ager)
}

type school struct {
	students []*student
	ager Ager
}

func(s school) register(name string, dob string)([]*student, error){
	studentName, err := createName(name)
	if err != nil {
		return nil, err
	}
	studentAge, err2 := s.calculateAgeFromDOB(dob, s.ager)
	if err2 != nil {
		return nil, err2
	}

	student := student{
		name: *studentName,
		dob: dob,
		age: studentAge,
	}

	s.students = append(s.students, &student)

	return s.students, nil
}

func(s school) calculateAgeFromDOB(dob string, ager Ager)(result int, err error){
	const shortForm = "2006-Jan-02"
	date, err := time.Parse(shortForm, dob)
	if err != nil {
		return -1, fmt.Errorf("this date is not valid: %s", dob)
	}

	result = s.ager.Age(date)

	return result, nil
}

func(s school) list() string{
	result := ""

	for i, student := range s.students{
		result += fmt.Sprint(student)
		if i != len(s.students) - 1 {
			result += "\n"
		}
	}

	return result
}

type student struct {
	name name
	dob  string
	age  int
}

func (s student) String() string {
	return fmt.Sprintf("Name: %s, DOB: %s, Age: %d", s.name , s.dob, s.age)
}

func assignment10(){
	mySchool := school{
		students: []*student{},
		ager: &DefaultAger{},
	}

	mySchool.students, _ = mySchool.register("Leif Kemp", "2001-Nov-25")
	mySchool.students, _ = mySchool.register("Leif Alexander Kemp", "2001-Nov-25")
	mySchool.students, _ = mySchool.register("Leif Pemp", "2002-Nov-25")
	mySchool.students, _ = mySchool.register("Keif Lemp", "2001-Nov-24")
	mySchool.students, _ = mySchool.register("Fiel Pmek", "1002-Nov-25")
	mySchool.students, _ = mySchool.register("Evil Leif Kemp", "2001-Oct-31")
	mySchool.students, _ = mySchool.register("Leaf Kemp", "2004-Aug-25")
	mySchool.students, _ = mySchool.register("Beef Kemp", "1950-Nov-25")
	mySchool.students, _ = mySchool.register("Kemp Leif", "1998-Jan-31")
	mySchool.students, _ = mySchool.register("Llll Kkkk", "1111-Jun-11")

	pl(mySchool.list())
}