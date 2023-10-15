package test1

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
)

const (
	theGroupURI = "uri:libgorm:group:module-gorm-sqlserver/src/test/test1"
)

// TableReg ...
type TableReg struct {

	//starter:component
	_as func(libgorm.GroupRegistry) //starter:as(".")

	Context application.Context //starter:inject("context")

	// agent libgorm.DatabaseAgent
}

func (inst *TableReg) _impl(a libgorm.Group, b libgorm.GroupRegistry) {
	a = inst
	b = inst
}

// // Group ...
// func (inst *TableReg) Group() *libgorm.Group {

// 	return &libgorm.Group{
// 		Prototypes: list,
// 		Name:       "default",
// 		Prefix:     "test1_",
// 		Enabled:    true,
// 	}
// }

// Groups ...
func (inst *TableReg) Groups() []*libgorm.GroupRegistration {
	//  inst.agent.Init(c.Database)

	r1 := &libgorm.GroupRegistration{
		URI:   theGroupURI,
		Group: inst,
	}

	return []*libgorm.GroupRegistration{r1}
}

// Prototypes ...
func (inst *TableReg) Prototypes() []any {
	list := make([]any, 0)
	list = append(list, &Table1{})
	list = append(list, &Table2{})
	list = append(list, &Table3{})
	return list
}
