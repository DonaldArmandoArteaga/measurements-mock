package main

import (
	"bytes"
	"encoding/json"
	"errors"
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
	API_URL     = os.Getenv("API_URL")
)

func init() {
	currentTime := time.Now()
	date := fmt.Sprintf("%d-%d-%d %d-%d-%d", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second())

	err := os.Mkdir("logs", os.ModePerm)
	if !errors.Is(err, os.ErrExist) && err != nil {
		panic("can't create logs directory: " + err.Error())
	}
	file, err := os.Create("./logs/" + date + "_logs")
	if err != nil {
		panic("can't create log file" + err.Error())
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	InfoLogger.Println("API_URL:", API_URL)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {

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

		for {
			time.Sleep(2 * time.Minute)

			measurers.TemperatureMeasurer3.Date = time.Now()
			measurers.TemperatureMeasurer3.Values.Temperature = 15 + rand.Intn(34-15)
			SendRequest(measurers.TemperatureMeasurer3)

			measurers.TemperatureMeasurer4.Date = time.Now()
			measurers.TemperatureMeasurer4.Values.Temperature = 25 + rand.Intn(40-25)
			SendRequest(measurers.TemperatureMeasurer4)

		}

	}()

	go func() {

		for {
			time.Sleep(4 * time.Minute)

			measurers.TemperatureMeasurer5.Date = time.Now()
			measurers.TemperatureMeasurer5.Values.Temperature = 15 + rand.Intn(34-15)
			SendRequest(measurers.TemperatureMeasurer5)
		}

	}()

	wg.Wait()
}

func SendRequest[T measurers.Temperature | measurers.Energy](measurer *measurers.Measurers[T]) {

	requestURL := fmt.Sprintf("%s%s", API_URL, "v1/input-data")

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(measurer)
	if err != nil {
		ErrorLogger.Println("Encoder: could not encoding the measurer", err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, requestURL, &buf)
	if err != nil {
		ErrorLogger.Println("Client: could not create request: ", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ErrorLogger.Println("Client: error making http request: ", err)
	}

	InfoLogger.Println("Client: got response! => status code: ", res.StatusCode)

}
