package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

const (
	passwdLength  = 16
	emailListFile = "email-list.txt"
	infoLogFile   = "created-creds.txt"
)

func main() {
	var formFilename string
	registeredEmails := LoadEmails(emailListFile)

	if len(os.Args) > 1 {
		formFilename = os.Args[1]
	} else {
		fmt.Printf("usage: %s <form_response.csv>", os.Args[0])
		os.Exit(1)
	}

	formFile, err := os.Open(formFilename)
	if err != nil {
		log.Fatalln("failed to open csv file:", err)
	}

	reader := csv.NewReader(formFile)
	reader.FieldsPerRecord = -1

	formData, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("failed to read data from csv:", err)
	}

	// Microsoft Form
	// ID = 0
	// StartTime = 1
	// EndTime = 2
	// Email = 3 (*)
	// MicrosoftName (hust_name) = 4
	// LastModTime = 5
	// RealName = 6 (*)
	// MSSV = 7 (*)
	// LopSV = 8 (*)
	// DiscordUname = 9 (*)
	// CTFPreferences = 10

	for idx, row := range formData {
		if idx != 0 {
			tempUser := User{
				Email:           row[3],
				Password:        genPasswd(passwdLength),
				Fullname:        row[6],
				MSSV:            row[7],
				LopSV:           row[8],
				DiscordUsername: row[9],
			}
			log.Println(tempUser)
			if _, exists := registeredEmails[tempUser.Email]; !exists {
				registeredEmails[tempUser.Email] = struct{}{}
				tempUser.AddUser()
				AppendWithNewline(emailListFile, tempUser.Email)
				tempUser.LogUserInfo(infoLogFile)
			} else {
				log.Printf("found an user attempting to sign-up more than once: %s\n", tempUser.Email)
			}
		}
	}
}
