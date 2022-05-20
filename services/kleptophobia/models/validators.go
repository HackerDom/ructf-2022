package models

import (
	"errors"
	"regexp"
)

var NameRegex = regexp.MustCompile("^[a-zA-Z\\s]{3,30}$")
var MiddleNameRegex = regexp.MustCompile("^[a-zA-Z\\s]{30}$")
var DiagnosisRegex = regexp.MustCompile("^[a-zA-Z\\d\\s]{3,50}$")
var UsernameRegex = regexp.MustCompile("^[a-zA-Z\\d_\\-]{6,30}$")

func ValidatePrivatePerson(person *PrivatePerson) error {
	if !NameRegex.MatchString(person.FirstName) {
		return errors.New("incorrect first name")
	}
	if !NameRegex.MatchString(person.SecondName) {
		return errors.New("incorrect second name")
	}
	if !DiagnosisRegex.MatchString(person.Diagnosis) {
		return errors.New("incorrect diagnosis")
	}
	if person.MiddleName != nil && !MiddleNameRegex.MatchString(*person.MiddleName) {
		return errors.New("incorrect middle name")
	}
	if UsernameRegex.MatchString(person.Username) {
		return errors.New("incorrect username")
	}
	return nil
}
