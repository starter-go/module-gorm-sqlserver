package sqlserver

import (
	"github.com/starter-go/application"
	modulegormsqlserver "github.com/starter-go/module-gorm-sqlserver"
	"github.com/starter-go/module-gorm-sqlserver/gen/gen4sqlserver"
)

// Module 函数返回 sqlserver 的 libgorm 驱动模块
func Module() application.Module {
	mb := modulegormsqlserver.NewDriverModule()
	mb.Components(gen4sqlserver.ExportComForGormSqlserver)
	return mb.Create()
}
