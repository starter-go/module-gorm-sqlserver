package gen4sqlserver

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForGormSqlserver ...
func ExportComForGormSqlserver(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
