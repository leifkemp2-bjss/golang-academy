package main

import (
	"fmt"
	"reflect"
	"slices"
	"time"
)

type institute interface{
	register(name string, dob string)(error)
	remove(name name)()
	list() string
	calculateAgeFromDOB(dob string, ager Ager)(int, error)
}

type school struct {
	students []student
	ager Ager
}

// Registers a new student to the school's students array
func(s *school) register(name string, dob string)(error){
	studentName, err := createName(name)
	if err != nil {
		return err
	}
	studentAge, err2 := s.calculateAgeFromDOB(dob, s.ager)
	if err2 != nil {
		return err2
	}

	student := student{
		name: *studentName,
		dob: dob,
		age: studentAge,
	}

	s.students = append(s.students, student)

	return nil
}

// Remove the first student that has this name
func(s *school) remove(name name)(){
	s.students = slices.DeleteFunc(s.students, func(s student) bool {
		return reflect.DeepEqual(s.name, name)
	})
}

func(s *school) list() string{
	result := ""

	for i, student := range s.students{
		result += fmt.Sprint(student)
		if i != len(s.students) - 1 {
			result += "\n"
		}
	}

	return result
}

func(s *school) calculateAgeFromDOB(dob string, ager Ager)(int, error){
	const shortForm = "2006-Jan-02"
	date, err := time.Parse(shortForm, dob)
	if err != nil {
		return -1, fmt.Errorf("this date is not valid: %s", dob)
	}

	result := s.ager.Age(date)

	return result, nil
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
		students: []student{},
		ager: &DefaultAger{},
	}

	_ = mySchool.register("Leif Kemp", "2001-Nov-25")
	_ = mySchool.register("Leif Alexander Kemp", "2001-Nov-25")
	_ = mySchool.register("Leif Pemp", "2002-Nov-25")
	_ = mySchool.register("Keif Lemp", "2001-Nov-24")
	_ = mySchool.register("Fiel Pmek", "1002-Nov-25")
	_ = mySchool.register("Evil Leif Kemp", "2001-Oct-31")
	_ = mySchool.register("Leaf Kemp", "2004-Aug-25")
	_ = mySchool.register("Beef Kemp", "1950-Nov-25")
	_ = mySchool.register("Kemp Leif", "1998-Jan-31")
	_ = mySchool.register("Llll Kkkk", "1111-Jun-11")

	pl(mySchool.list())

	pl()
	pl("Evil Leif Kemp got expelled for being evil. Here's the new student list:")
	mySchool.remove(name{firstName: "Evil", middleName: "Leif", lastName: "Kemp"})

	pl(mySchool.list())
}