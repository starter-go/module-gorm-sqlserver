package modulegormsqlserver

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/modgorm"
	"github.com/starter-go/module-gorm-sqlserver/gen/gen4sqlserver"
)

const (
	theModuleName     = "github.com/starter-go/module-gorm-sqlserver"
	theModuleVersion  = "v0.9.0"
	theModuleRevision = 3
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// Module 导出模块 ['github.com/starter-go/module-gorm-sqlserver']
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)

	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen4sqlserver.ExportComForGormSqlserver)

	mb.Depend(modgorm.Module())

	return mb.Create()
}
