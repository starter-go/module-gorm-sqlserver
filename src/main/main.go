package main

import (
	modulegormsqlserver "github.com/starter-go/module-gorm-sqlserver"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(nil)
	i.MainModule(modulegormsqlserver.Module())
	i.WithPanic(true).Run()
}
