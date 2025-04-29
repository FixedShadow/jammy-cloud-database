package common

import (
	"github.com/google/uuid"
	"strings"
)

func UUID() string {
	return uuid.New().String()
}

func Generate32RandomString() string {
	uuidStr := uuid.New().String()
	return strings.ReplaceAll(uuidStr, "-", "")
}
