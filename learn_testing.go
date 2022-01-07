// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

var theEmails = []string{"gmail.com", "outlook.com", "yahoo.com", "hotmail.com"}

func isInTheEmails(email string) bool {
	for _, provider := range theEmails {
		fmt.Printf("value is %s \n", provider)
		if strings.Contains(email, provider) {
			return true
		}
	}

	return false
}
