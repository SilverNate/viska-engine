package features

import (
	"testing"
	"viska/auth-service/features/steps"

	"github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{"../features"},
	}

	status := godog.TestSuite{
		Name:                "auth",
		ScenarioInitializer: steps.InitializeScenario,
		Options:             &opts,
	}.Run()

	if status != 0 {
		t.Fail()
	}
}
