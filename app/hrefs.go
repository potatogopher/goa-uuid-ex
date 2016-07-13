//************************************************************************//
// API "UUID-Example": Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/potatogopher/goa-uuid-ex/design
// --notest=true
// --out=$(GOPATH)/src/github.com/potatogopher/goa-uuid-ex
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	return fmt.Sprintf("/users/%v", userID)
}
