package main

import (
	"fmt"
	"myapp/staff"
)

var employess=[]staff.Employee{
	{FirstName:"John",LastName: "Smith",Salary: 30000,FullTime: true},
	{FirstName:"Sally",LastName: "Jones",Salary: 50000,FullTime: true},
	{FirstName:"Mark",LastName: "Smithers",Salary: 60000,FullTime: true},
	{FirstName:"Allan",LastName: "Anderson",Salary: 15000,FullTime: false},
	{FirstName:"margaret",LastName: "Carter",Salary: 100000,FullTime: false},
}
func main(){
	myStaff:=staff.Office{
		AllStaff: employess,
	}
}