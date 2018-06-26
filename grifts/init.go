package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/gormrecipe/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
