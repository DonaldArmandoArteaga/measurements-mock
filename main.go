package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"mesasurements-mock/measurers"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for {
			time.Sleep(1 * time.Minute)

			measurers.TemperatureMeasurer1.Date = time.Now()
			measurers.TemperatureMeasurer1.Values.Temperature = 15 + rand.Intn(34-15)
			SendRequest(measurers.TemperatureMeasurer1)

			measurers.TemperatureMeasurer2.Date = time.Now()
			measurers.TemperatureMeasurer2.Values.Temperature = 25 + rand.Intn(40-25)
			SendRequest(measurers.TemperatureMeasurer2)
		}

	}()

	go func() {
		wg.Add(1)
		for {
			time.Sleep(5 * time.Minute)

			measurers.TemperatureMeasurer3.Date = time.Now()
			measurers.TemperatureMeasurer3.Values.Temperature = 15 + rand.Intn(34-15)
			SendRequest(measurers.TemperatureMeasurer3)

			measurers.TemperatureMeasurer4.Date = time.Now()
			measurers.TemperatureMeasurer4.Values.Temperature = 25 + rand.Intn(40-25)
			SendRequest(measurers.TemperatureMeasurer4)

		}

	}()

	go func() {
		wg.Add(1)
		for {
			time.Sleep(10 * time.Minute)

			measurers.TemperatureMeasurer5.Date = time.Now()
			measurers.TemperatureMeasurer5.Values.Temperature = 15 + rand.Intn(34-15)
			SendRequest(measurers.TemperatureMeasurer5)
		}

	}()

	wg.Wait()
}

func SendRequest[T any](measurer *measurers.Measurers[T]) {

	requestURL := "https://787rmeid5e.execute-api.us-east-1.amazonaws.com/prod/v1/input-data"

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(measurer)
	if err != nil {
		fmt.Print(err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, requestURL, &buf)
	if err != nil {
		ErrorLogger.Println("client: could not create request: ", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ErrorLogger.Println("client: error making http request: ", err)
	}

	InfoLogger.Println("client: got response! => status code: ", res.StatusCode)

}
