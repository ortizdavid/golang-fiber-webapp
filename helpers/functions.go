package helpers

import (
	
	"fmt"
	"log"
	"time"
	"strconv"
	"crypto/rand"
	"encoding/base64"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	uniqueId := uuid.New()
	return uniqueId.String()
}

func GenerateCode(prefix string) string {
	timestamp := time.Now().Format("20060102150405") 
    return prefix + timestamp  
}

func GenerateRandomToken() string {
	length := 100
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	token := base64.RawURLEncoding.EncodeToString(randomBytes)
	return token
}

func ConvertToInt(value string) int {
	intValue, _ := strconv.Atoi(value)
	return intValue
}

func ConvertToString(value any) string {
	return fmt.Sprintf("%v", value)
}