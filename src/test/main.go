package main

import (
	"embed"
	"os"

	"github.com/starter-go/application"
	modulegormsqlserver "github.com/starter-go/module-gorm-sqlserver"
	"github.com/starter-go/module-gorm-sqlserver/gen/gen4test"
	"github.com/starter-go/starter"
)

func main() {

	m := module()

	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}

////////////////////////////////////////////////////////////////////////////////

//go:embed "resources"
var theResFS embed.FS

func module() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name("module-gorm-sqlserver/src/test")
	mb.Version("v1")
	mb.Revision(1)

	mb.EmbedResources(theResFS, "resources")
	mb.Components(gen4test.ExportComForGormSqlserverTest)
	mb.Depend(modulegormsqlserver.Module())

	return mb.Create()
}
