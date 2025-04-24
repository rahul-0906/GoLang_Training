package main

import "fmt"

/*
Assume we have a func getting output as emailID and mob no if its a emailID user should call a Greet()
method with reference to emailing if its mob no it should call same Greet method but
with the reference of sending message
*/
type Contact interface {
	Greet() string
}

type Email struct {
	emailId string
}

type Mobile struct {
	mobileNo string
}

func (e Email) Greet() string {
	return e.emailId
}

func (m Mobile) Greet() string {
	return m.mobileNo
}

func displayInformation(c Contact) {

	fmt.Println("Details: ", c.Greet())
}

func main() {

	emailId := Email{emailId: "rahulbabu.mv@gmail.com"}

	mobileNo := Mobile{mobileNo: "9876543210"}

	fmt.Println("Email Id:")
	displayInformation(emailId)

	fmt.Println("Mobile No:")
	displayInformation(mobileNo)

}
