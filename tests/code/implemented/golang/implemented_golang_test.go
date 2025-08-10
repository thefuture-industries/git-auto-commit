package golang

import (
	"git-auto-commit/tests"
	"testing"
)

func TestImplementedGolang(t *testing.T) {
	mocks := tests.SaveMocks()
	defer mocks.Apply()
}
