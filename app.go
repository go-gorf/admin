package admin

import (
	"github.com/go-gorf/gorf"
)

func setup() error {
	return nil
}

var AdminApp = gorf.GorfBaseApp{
	Name:         "gorf-admin",
	RouteHandler: Urls,
	SetUpHandler: setup,
}
