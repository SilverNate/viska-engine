package steps

import (
	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"viska/auth-service/config"
)

type loginTestContext struct {
	router     *gin.Engine
	token      string
	resp       *httptest.ResponseRecorder
	configAuth config.Config
}

var ctx loginTestContext

type MockUserRepository struct{}

func (c *loginTestContext) aUserExists(email, password string) error {
	return nil
}

func (c *loginTestContext) iShouldReceiveJWT() error {

	return nil
}

func InitializeScenario(sc *godog.ScenarioContext) {
	ctx = loginTestContext{}

	sc.Step(`^the user "([^"]*)" exists with password "([^"]*)"$`, ctx.aUserExists)
	sc.Step(`^the response code should be (\d+)$`, ctx.iShouldReceiveJWT)
}
