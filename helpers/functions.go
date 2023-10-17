package helpers

import (
	"fmt"
	"strconv"
	"time"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	uniqueId := uuid.New()
	return uniqueId.String()
}

func ConvertToInt(value string) int {
	intValue, _ := strconv.Atoi(value)
	return intValue
}

func ConvertToString(value any) string {
	return fmt.Sprintf("%v", value)
}

func GenerateCode(prefix string) string {
	timestamp := time.Now().Format("20060102150405") 
    return prefix + timestamp  
}
