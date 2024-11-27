package ctfd

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"

	"github.com/teebow1e/ctfd-useradd/utils"
)

func Ping() {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.SetRequestURI(viper.GetString("config.CTFD_URL") + "/api/v1/tokens")
	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", viper.GetString("config.CTFD_API_KEY")))

	if err := fasthttp.Do(req, resp); err != nil {
		log.Fatalf("error while pinging CTFd: %v\n", err)
	}

	if resp.StatusCode() == 200 {
		log.Println("Pong! CTFd replied and your token is still good.")
		log.Print(utils.B2S(resp.Body()))
	} else {
		log.Println("Server responded with bad request. Here is the response:")
		log.Print(utils.B2S(resp.Body()))
	}
}
