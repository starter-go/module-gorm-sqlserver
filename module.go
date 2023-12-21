package modulegormsqlserver

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm/modules/libgorm"
)

const (
	theModuleName     = "github.com/starter-go/module-gorm-sqlserver"
	theModuleVersion  = "v0.9.12"
	theModuleRevision = 7
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// NewDriverModule 导出模块 ['github.com/starter-go/module-gorm-sqlserver']
func NewDriverModule() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Depend(libgorm.Module())

	return mb
}
