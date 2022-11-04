package graph

import "github.com/mgeale/homeserver/cmd/web/app"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	App *app.Application
}
