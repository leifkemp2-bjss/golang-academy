package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDisplayStudent(t *testing.T){
	student := student{
		name: name{firstName: "John", middleName: "", lastName: "Student"},
		dob: "2006-Jan-01",
		age: 18,
	}

	got := fmt.Sprintf("%+v", student)
	want := "Name: John Student, DOB: 2006-Jan-01, Age: 18"

	if got != want {
		t.Errorf("expected %s, got %s", want, got)
	}
}

func TestCreateStudent(t *testing.T){
	school := school{
		students: []student{},
		ager: &MockAger{},
	}

	students := []struct{
		name string
		dob string
		want student
	}{
		{name: "John Student", dob: "2006-Jan-01", want: student{
			name: name{firstName: "John", middleName: "", lastName: "Student"},
			dob: "2006-Jan-01",
			age: 18,
		}},
		{name: "Peter Student", dob: "2007-Jan-01", want: student{
			name: name{firstName: "Peter", middleName: "", lastName: "Student"},
			dob: "2007-Jan-01",
			age: 17,
		}},
		{name: "John Peter Student", dob: "2008-Sep-25", want: student{
			name: name{firstName: "John", middleName: "Peter", lastName: "Student"},
			dob: "2008-Sep-25",
			age: 15,
		}},
	}

	for i, test := range students{
		t.Run(fmt.Sprintf("testing creation of student %s", test.name), func(t *testing.T) {
			err := school.register(test.name, test.dob)
			if err != nil {
				t.Errorf("the program hit an unexpected error: %s", err)
				return
			}
			if !reflect.DeepEqual(school.students[i], test.want){
				t.Errorf("expected %+v, got %+v", test.want, school.students[i])
			}
		})
	}

	if len(school.students) != 3 {
		t.Errorf("the school does not have 3 students, has %v", school.students)
	}
}

func TestCreateStudentInvalid(t *testing.T){
	school := school{
		students: []student{},
		ager: &MockAger{},
	}

	students := []struct{
		name string
		dob string
		want string
	}{
		{name: "InvalidName", dob: "2006-Jan-01", want: "this is not a valid name"},
		{name: "Invalid Birth Date", dob: "2006-Jan-34", want: "this date is not valid: 2006-Jan-34"},
		{name: "Invalid Wrong Birth Date", dob: "nonsensestring", want: "this date is not valid: nonsensestring"},
	}

	for _, test := range students{
		t.Run(fmt.Sprintf("testing creation of student %s", test.name), func(t *testing.T) {
			err := school.register(test.name, test.dob)
			if err == nil {
				t.Errorf("this student should've produced an error: %+v", test)
				return
			}
			
			if fmt.Sprint(err) != test.want {
				t.Errorf("expected %s, got %s", test.want, err)
			}
		})
	}
}

func TestListStudents(t *testing.T){
	cases := []struct{
		school school
		want string
	}{
		{
			school: school{
				students: []student{},
				ager: &MockAger{},
			}, 
			want: "",
		},
		{
			school: school{
				students: []student{
					{
						name: name{firstName: "John", middleName: "", lastName: "Student"},
						dob: "2006-Jan-01",
						age: 18,
					},
				},
				ager: &MockAger{},
			},
			want: "Name: John Student, DOB: 2006-Jan-01, Age: 18",
		},
		{
			school: school{
				students: []student{
					{
						name: name{firstName: "John", middleName: "", lastName: "Student"},
						dob: "2006-Jan-01",
						age: 18,
					},
					{
						name: name{firstName: "John", middleName: "Peter", lastName: "Student"},
						dob: "2006-Jan-31",
						age: 17,
					},
				},
				ager: &MockAger{},
			},
			want: "Name: John Student, DOB: 2006-Jan-01, Age: 18\nName: John Peter Student, DOB: 2006-Jan-31, Age: 17",
		},
	}

	for _, test := range cases{
		t.Run(fmt.Sprintf("testing listing of school with %d students", len(test.school.students)), func(t *testing.T) {
			got := test.school.list()

			if got != test.want {
				t.Errorf("expected %s, got %s", test.want, got)
			}
		})
	}
}

func TestRemoveStudent(t *testing.T){
	school := school{
		students: []student{
			{
				name: name{firstName: "John", middleName: "", lastName: "Student"},
				dob: "2006-Jan-01",
				age: 18,
			},
			{
				name: name{firstName: "John", middleName: "Peter", lastName: "Student"},
				dob: "2006-Jan-31",
				age: 17,
			},
		},
		ager: &MockAger{},
	}

	if len(school.students) != 2 {
		t.Error("the school should have 2 students")
	}

	school.remove(name{firstName: "John", middleName: "", lastName: "Student"})

	if len(school.students) != 1 {
		t.Error("the school should have 1 student")
	}

	if !reflect.DeepEqual(school.students[0].name, name{firstName: "John", middleName: "Peter", lastName: "Student"}){
		t.Error("the remaining student should be John Peter Student")
	}

	// removing a student who doesn't exist should do nothing
	school.remove(name{firstName: "Doesn't", middleName: "", lastName: "Exist"})

	if len(school.students) != 1 {
		t.Error("the school should have 1 student after removing a student that doesn't exist")
	}

	if !reflect.DeepEqual(school.students[0].name, name{firstName: "John", middleName: "Peter", lastName: "Student"}){
		t.Error("the remaining student should be John Peter Student after removing a student that doesn't exist")
	}
}