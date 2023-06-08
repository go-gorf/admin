package admin

import (
	"github.com/go-gorf/gorf"
)

func setup() error {
	return nil
}

var App = gorf.BaseApp{
	Title:        "Gorf admin",
	Info:         "Gorf Admin Panel",
	RouteHandler: Urls,
	SetUpHandler: setup,
}
