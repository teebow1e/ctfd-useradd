package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tidwall/gjson"
)

type User struct {
	Email           string
	Password        string
	Fullname        string
	MSSV            string
	LopSV           string
	DiscordUsername string
}

func (user *User) AddUser() {
	data := map[string]interface{}{
		"name":        fmt.Sprintf("%s - %s", user.Fullname, user.MSSV),
		"email":       user.Email,
		"password":    user.Password,
		"affiliation": user.LopSV,
		"type":        "user",
		"verified":    true,
		"hidden":      false,
		"banned":      false,
		"fields":      []string{},
	}

	builder, err := json.Marshal(data)
	if err != nil {
		log.Fatalln("[-] failed to build json")
	}

	log.Printf("[+] Adding user %s\n", user.Fullname)

	body, err := PostJson("/api/v1/users", builder)
	if err != nil {
		log.Printf("Error during adding user: %s\n", err)
	}

	isSuccess := gjson.GetBytes(body, "success")
	if isSuccess.Value() == true {
		log.Println("[+] Add user successfully")
	} else {
		log.Println("something went wrong during adding user:", b2s(body))
	}
}

func (user *User) LogUserInfo(logFile string) {
	// email, password, ten, mssv, lopsv, discord
	logData := fmt.Sprintf("%s,%s,%s,%s,%s,%s", user.Email, user.Password, user.Fullname, user.MSSV, user.LopSV, user.DiscordUsername)
	AppendWithNewline(logFile, logData)
}
