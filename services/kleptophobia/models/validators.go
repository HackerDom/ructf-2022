package models

import (
	"errors"
	"regexp"
)

var NameRegex = regexp.MustCompile(`^[a-zA-Z\s]{3,30}$`)
var DiagnosisRegex = regexp.MustCompile(`^[!-~]{3,50}$`)
var UsernameRegex = regexp.MustCompile(`^[a-zA-Z\d_\-]{6,30}$`)

func ValidatePrivatePerson(person *PrivatePerson) error {
	if !NameRegex.MatchString(person.FirstName) {
		return errors.New("incorrect first name: " + person.FirstName)
	}
	if !NameRegex.MatchString(person.MiddleName) {
		return errors.New("incorrect middle name" + person.MiddleName)
	}
	if !NameRegex.MatchString(person.SecondName) {
		return errors.New("incorrect second name: " + person.SecondName)
	}
	if !DiagnosisRegex.MatchString(person.Diagnosis) {
		return errors.New("incorrect diagnosis: " + person.Diagnosis)
	}
	if !UsernameRegex.MatchString(person.Username) {
		return errors.New("incorrect username: " + person.Username)
	}
	return nil
}
