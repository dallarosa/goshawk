package main

import (
	"os"
	"net/mail"
	"github.com/bthomson/mbox"
)

const (
	SPOOL = "/var/spool/mail/"
)

func getInbox(user string) ([]*mail.Message, error) {
 	mboxFile, err := os.Open(SPOOL + user)
	if err != nil {
		return nil,err
 	}
	
 	inbox, err := mbox.Read(mboxFile,true)

	return inbox, err
}
