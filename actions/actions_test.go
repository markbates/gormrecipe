package actions

import (
	"testing"

	"github.com/gobuffalo/packr"
	"github.com/markbates/gormrecipe/suite"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(App(), packr.NewBox("../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
	// suite.Run(t, as)
}
