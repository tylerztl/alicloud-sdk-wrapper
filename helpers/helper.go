package helpers

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
	"zig-cloud/commons"
)

// define the helper functions

func GenerateInstanceName() string {
	return fmt.Sprintf("%s"+commons.SeparatorHype+"%s", commons.AliCloudInstanceName, GenerateRandomString())
}

func GenerateSecurityGroupName() string {
	return fmt.Sprintf("%s"+commons.SeparatorHype+"%s", commons.AliCloudSecurityGroupName, GenerateRandomString())
}

func GenerateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyz" + "0123456789")
	length := 5
	b := new(bytes.Buffer)
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func GetTimeoutMessage(product, status string) string {
	return fmt.Sprintf("Waitting for %s %s is timeout.", product, status)
}

func ConvertListToJsonString(configured []string) string {
	if len(configured) < 1 {
		return ""
	}
	result := "["
	for i, v := range configured {
		result += "\"" + v + "\""
		if i < len(configured)-1 {
			result += ","
		}
	}
	result += "]"
	return result
}
