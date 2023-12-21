package main

import (
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(nil)
	i.MainModule(sqlserver.Module())
	i.WithPanic(true).Run()
}
