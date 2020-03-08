
package main

import (
	"fmt"
	"time"
	"bytes"
	"io/ioutil"
	"sync"
	"net/http"

    logger "github.com/DarkMetrix/gofra/pkg/logger/seelog"

	log "github.com/cihub/seelog"
)

func main() {
	defer log.Flush()

	// init log
    err := logger.Init("../config/log.config", "comments_test")

	if err != nil {
		log.Warnf("Init logger failed! error:%v", err.Error())
	}

	log.Info("====== Test [comments] begin ======")
	defer log.Info("====== Test [comments] end ======")

	// begin test
	testHealth()

	time.Sleep(time.Second * 1)
}

var (
	httpInitOnce sync.Once
	httpClient *http.Client
)

func testHealth() {
	// init http client
	httpInitOnce.Do(func() {
		httpClient = &http.Client{
			Transport: &http.Transport{
				MaxIdleConns: 20,
			},
			Timeout: 5 * time.Second,
		}

		log.Infof("Init httpClient success!")
	})

	// http call
	for index := 0; index < 1; index++ {
		url := fmt.Sprintf("http://%v%v", "0.0.0.0:9998", "/health")
		reader := bytes.NewReader([]byte("{}"))

		httpReq, err := http.NewRequest("POST", url, reader)

		if err != nil {
			log.Warnf("http.NewRequest failed! error:%v", err.Error())
			return
		}

		// copy header
		httpReq.Header.Add("Content-Type", "application/json")

		resp, err := httpClient.Do(httpReq)

		if err != nil {
			log.Warnf("client.Do failed! error:%v", err.Error())
			return
		}

		// make sure body is closed so that the connection could be reused
		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Warnf("ioutil.ReadAll failed! error:%v", err.Error())
			return
		}

		log.Infof("http request end! url:%v, resp body:%v", url, string(respBody))
	}
}