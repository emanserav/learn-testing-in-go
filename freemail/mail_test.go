// You can edit this code!
// Click here and start typing.
package freemail

import "testing"

func TestGmail(t *testing.T) {
	gmail := "gmail.com"
	if !isInTheEmails(gmail) {
		t.Fail()
	}
}
