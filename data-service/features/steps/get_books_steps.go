package steps

import (
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"io/ioutil"
	"net/http"
	"viska/data-service/internal/library"
)

var response *http.Response
var body []byte

func iSendAGetRequestWithToken() error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8081/api/books", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQHZpc2thLmlvIiwiZXhwIjoxNzQ0NzIxMDY1LCJ1c2VyX2lkIjoxfQ.ftJHVpqRathyLaKpkKFYWAM3GoyjP35-kFIp0nBV6OY")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	response = resp
	body, _ = ioutil.ReadAll(resp.Body)
	return nil
}

func theResponseCodeShouldBe200() error {
	if response.StatusCode != 200 {
		return fmt.Errorf("expected 200 but got %d", response.StatusCode)
	}
	return nil
}

func theResponseShouldContainBooks() error {
	var books []library.Book
	err := json.Unmarshal(body, &books)
	if err != nil {
		return err
	}

	if len(books) == 0 {
		return fmt.Errorf("expected at least one book, got 0")
	}
	return nil
}

func theDataserviceIsRunning() error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the data-service is running$`, theDataserviceIsRunning)
	ctx.Step(`^I send a GET request to "/api/books" with a valid token$`, iSendAGetRequestWithToken)
	ctx.Step(`^the response code should be 200$`, theResponseCodeShouldBe200)
	ctx.Step(`^the response should contain a list of books$`, theResponseShouldContainBooks)
}
