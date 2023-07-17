package test1

import (
	"sort"

	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/vlog"
)

type Table1 struct {
	ID    int
	Value int
}

type Table2 struct {
	ID    int
	Value string
}

type Table3 struct {
	ID    int
	Value bool
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamePrefix string

// TableName ...
func (t Table1) TableName() string {
	return theTableNamePrefix + "table1"
}

// TableName ...
func (t Table2) TableName() string {
	return theTableNamePrefix + "table2"
}

// TableName ...
func (t Table3) TableName() string {
	return theTableNamePrefix + "table3"
}

////////////////////////////////////////////////////////////////////////////////

// TableReg ...
type TableReg struct {
	//starter:component
	_as func(libgorm.TableRegistry) //starter:as(".")

	Context application.Context //starter:inject("context")

}

func (inst *TableReg) _impl() {
	inst._as(inst)
}

// Life ...
func (inst *TableReg) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
	}
}

func (inst *TableReg) init() error {
	vlog.Info("properties:")
	table := inst.Context.GetProperties().Export(nil)
	keys := make([]string, 0)
	for k := range table {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		val := table[key]
		vlog.Info("  %s: %s", key, val)
	}
	return nil
}

// ListTableRegistrations ...
func (inst *TableReg) ListTableRegistrations() []*libgorm.TableRegistration {

	tgb := &libgorm.TableGroupBuilder{}
	tgb.Group(&libgorm.TableGroup{
		Name:         "test1",
		PrefixSetter: inst.setPrefix,
	})

	tgb.Entity(&Table1{})
	tgb.Entity(&Table2{})
	tgb.Entity(&Table3{})

	return tgb.Create()
}

func (inst *TableReg) setPrefix(prefix string) {
	theTableNamePrefix = prefix
}
