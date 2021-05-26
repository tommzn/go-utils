package utils

import (
	"regexp"

	uuid "github.com/satori/go.uuid"
)

func NewId() string {
	return uuid.NewV4().String()
}

func IsId(id string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(id)
}
