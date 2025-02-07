package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"strings"
	"unsafe"

	"github.com/valyala/fasthttp"
)

const (
	lowerChars  = "abcdefghijklmnopqrstuvwxyz"
	upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars = "0123456789"
	allChars    = lowerChars + upperChars + numberChars
)

func genPasswd(length int) string {
	required := []byte{
		lowerChars[randInt(len(lowerChars))],
		upperChars[randInt(len(upperChars))],
		numberChars[randInt(len(numberChars))],
	}

	password := make([]byte, length)
	copy(password, required)

	for i := len(required); i < length; i++ {
		password[i] = allChars[randInt(len(allChars))]
	}

	shuffle(password)

	return string(password)
}

func randInt(max int) int {
	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return int(b[0]) % max
}

func shuffle(slice []byte) {
	for i := len(slice) - 1; i > 0; i-- {
		j := randInt(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func b2s(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func PostJson(path string, data []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.SetRequestURI(CTFD_URL + path)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", CTFD_API_KEY))
	req.SetBody(data)

	if err := fasthttp.Do(req, resp); err != nil {
		log.Fatalln("[-] Failed to perform req", err)
		return []byte{}, err
	}
	return resp.Body(), nil
}

func AppendWithNewline(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data + "\n")
	if err != nil {
		return err
	}

	return writer.Flush()
}

func LoadEmails(filename string) map[string]struct{} {
	emailMap := make(map[string]struct{})

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("email list file not found. returning an empty list.")
			return emailMap
		}
		log.Fatalln("failed to open email list file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		email := strings.TrimSpace(scanner.Text())
		if email != "" {
			emailMap[email] = struct{}{}
		}
	}

	return emailMap
}
