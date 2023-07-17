package gen4test
import (
    p0ef6f2938 "github.com/starter-go/application"
    p6df7683da "github.com/starter-go/module-gorm-sqlserver/src/test/test1"
     "github.com/starter-go/application"
)

// type p6df7683da.TableReg in package:github.com/starter-go/module-gorm-sqlserver/src/test/test1
//
// id:com-6df7683daab96742-test1-TableReg
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-TableRegistry
// alias:
// scope:singleton
//
type p6df7683daa_test1_TableReg struct {
}

func (inst* p6df7683daa_test1_TableReg) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6df7683daab96742-test1-TableReg"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-TableRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6df7683daa_test1_TableReg) new() any {
    return &p6df7683da.TableReg{}
}

func (inst* p6df7683daa_test1_TableReg) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6df7683da.TableReg)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)


    return nil
}


func (inst*p6df7683daa_test1_TableReg) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


