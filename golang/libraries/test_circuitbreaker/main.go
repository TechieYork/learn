package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/afex/hystrix-go/hystrix"
)

func main() {
	testSyncCircuitBreaker()
	testAsyncCircuitBreaker()
	return
}

func testSyncCircuitBreaker() {
	log.Infof("====== test sync circuit breaker ======")

	// sync circuit breaker setting
	hystrix.ConfigureCommand("test_sync_get_user_name", hystrix.CommandConfig{
		Timeout: 1000,
		MaxConcurrentRequests: 10,
		RequestVolumeThreshold: 10,
		SleepWindow: 1000,
		ErrorPercentThreshold: 20,
	})

	// loop to do sync success
	for index := 0; index < 20; index++ {
		getUserNameSync(fmt.Sprintf("%v", index), false)
	}

	// loop to do sync failed
	for index := 0; index < 20; index++ {
		getUserNameSync(fmt.Sprintf("%v", index), true)
	}

	// loop to do sync success
	for index := 0; index < 20; index++ {
		getUserNameSync(fmt.Sprintf("%v", index), false)
	}

}

func getUserNameSync(id string, shouldFail bool) (string, error) {
	time.Sleep(100 * time.Millisecond)
	userName := ""

	// circuit breaker should used at client side
	err := hystrix.Do("test_sync_get_user_name", func() error {
		// do RPC or http request
		if shouldFail {
			return fmt.Errorf("RPC failed! error:%v", "just failed")
		}

		userName = fmt.Sprintf("NAME:%v", id)
		log.Infof("id:%v, name:%v", id, userName)

		return nil
	}, func(err error) error {
		// handle request failed situation
		log.Warnf("getUserNameSync failed! error:%v", err)
		return nil
	})

	if err != nil {
		log.Warnf("hystrix.Do failed! error:%v", err)
		return "", err
	}

	return userName, nil
}

func testAsyncCircuitBreaker() {
	log.Infof("====== test async circuit breaker ======")

	// sync circuit breaker setting
	hystrix.ConfigureCommand("test_async_get_user_name", hystrix.CommandConfig{
		Timeout: 1000,
		MaxConcurrentRequests: 10,
		RequestVolumeThreshold: 10,
		SleepWindow: 1000,
		ErrorPercentThreshold: 20,
	})

	// loop to do sync success
	for index := 0; index < 20; index++ {
		getUserNameAsync(fmt.Sprintf("%v", index), false)
	}

	// loop to do sync failed
	for index := 0; index < 20; index++ {
		getUserNameAsync(fmt.Sprintf("%v", index), true)
	}

	// loop to do sync success
	for index := 0; index < 20; index++ {
		getUserNameAsync(fmt.Sprintf("%v", index), false)
	}

}

func getUserNameAsync(id string, shouldFail bool) (string, error) {
	time.Sleep(100 * time.Millisecond)
	userName := ""

	finishedChan := make(chan bool, 1)
	errHandledChan := make(chan bool, 1)

	// circuit breaker should used at client side
	errChan := hystrix.Go("test_async_get_user_name", func() error {
		// do RPC or http request
		if shouldFail {
			return fmt.Errorf("RPC failed! error:%v", "just failed")
		}

		userName = fmt.Sprintf("NAME:%v", id)
		log.Infof("id:%v, name:%v", id, userName)

		finishedChan <- true

		return nil
	}, func(err error) error {
		// handle request failed situation
		log.Warnf("getUserNameSync failed! error:%v", err)

		errHandledChan <- true
		return nil
	})

	select {
	case err := <- errChan:
		log.Warnf("hystrix.Do failed! error:%v", err)
		return "", err
	case <- errHandledChan:
	case <- finishedChan:
	}

	return userName, nil
}
