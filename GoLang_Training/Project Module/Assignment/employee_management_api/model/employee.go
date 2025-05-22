package model

type Employee struct {
	EmpId       int    `json: "empid"`
	Name        string `json: "name"`
	EmailId     string `json: "emailid"`
	ProjectName string `json: "projectname"`
	Location    string `json: "location"`
	MobileNo    string `json: "mobileno"`
}
