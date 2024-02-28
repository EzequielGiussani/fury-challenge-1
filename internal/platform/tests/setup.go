package tests

import (
	"github.com/mercadolibre/fury_go-meli-toolkit-restful/rest"
)

func SetupTestCase() func() {
	rest.StartMockupServer()
	return func() {
		rest.ValidateCallCounts()
		rest.StopMockupServer()
	}
}

func CustomSetupTestCase(before func(), after func()) func() {
	rest.StartMockupServer()
	before()
	return func() {
		after()
		rest.ValidateCallCounts()
		rest.StopMockupServer()
	}
}
