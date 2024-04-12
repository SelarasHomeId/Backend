package general

import (
	"regexp"
	"time"
)

func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func IsValidPhone(phone string) bool {
	phoneNumberRegex := `^\+[1-9]\d{1,14}$`
	re := regexp.MustCompile(phoneNumberRegex)
	return re.MatchString(phone)
}

func DateTodayLocal() *time.Time {
	now := time.Now().UTC()
	return &now
}
