package common

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateRandomStringLess32(length int) string {
	uuidStr := uuid.New().String()
	return strings.ReplaceAll(uuidStr, "-", "")[:length]
}
