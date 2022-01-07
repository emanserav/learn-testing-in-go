package freemail2

import (
	"log"
	"testing"
)

func TestGmail(t *testing.T) {
	t.Parallel()
	if testing.Verbose() {
		log.Printf("Running %s test", t.Name())
	}
	gmail := "user@gmail.com"
	if !isInTheEmails(gmail) {
		t.Fail()
	}
}

func TestGmailPlus(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short tests")
	}
	t.Parallel()
	if testing.Verbose() {
		log.Println("Running tag test")
	}
	gmail := "user+tag@gmail.com"
	if !isInTheEmails(gmail) {
		t.Fail()
	}
}

func TestYahoo(t *testing.T) {
	t.Parallel()
	if testing.Verbose() {
		log.Printf("Running %s test", t.Name())
	}
	yahoo := "user@yahoo.com"
	if !isInTheEmails(yahoo) {
		t.Fail()
	}
}

func TestGoogle(t *testing.T) {
	t.Parallel()
	if testing.Verbose() {
		log.Printf("Running %s test", t.Name())
	}
	google := "user@google.com"
	if isInTheEmails(google) {
		t.Fail()
	}
}
