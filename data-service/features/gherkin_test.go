package features

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"viska/data-service/features/steps"
)

func TestMain(m *testing.M) {
	status := godog.TestSuite{
		Name:                "bookService",
		ScenarioInitializer: steps.InitializeScenario,
		Options: &godog.Options{
			Format:    "pretty",
			Paths:     []string{"../features"},
			Randomize: 0,
		},
	}.Run()

	os.Exit(status)
}
