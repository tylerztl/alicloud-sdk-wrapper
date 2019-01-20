package helpers

import (
	"fmt"
	"zig-cloud/commons"
	"math/rand"
	"time"
	"strings"
)

// define the helper functions

func GenerateInstanceName() string {
	return fmt.Sprintf("%s" + commons.SeparatorHype + "%s",commons.AliCloudInstanceName, GenerateRandomString())
}

func GenerateSecurityGroupName() string  {
	return fmt.Sprintf("%s" + commons.SeparatorHype + "%s",commons.AliCloudSecurityGroupName, GenerateRandomString())
}

func GenerateSecurityGroupRuleName() string {
	return fmt.Sprintf("%s" + commons.SeparatorHype + "%s",commons.AliCloudSecurityGroupRuleName, GenerateRandomString())
}

func GenerateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyz" + "0123456789")
	length := 5
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}