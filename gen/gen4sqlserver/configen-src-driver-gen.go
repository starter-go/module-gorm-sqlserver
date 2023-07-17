package gen4sqlserver
import (
    pc249b5436 "github.com/starter-go/module-gorm-sqlserver/dsqlserver"
     "github.com/starter-go/application"
)

// type pc249b5436.Driver in package:github.com/starter-go/module-gorm-sqlserver/dsqlserver
//
// id:com-c249b543600968ad-dsqlserver-Driver
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-Driver
// alias:
// scope:singleton
//
type pc249b54360_dsqlserver_Driver struct {
}

func (inst* pc249b54360_dsqlserver_Driver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c249b543600968ad-dsqlserver-Driver"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-Driver"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc249b54360_dsqlserver_Driver) new() any {
    return &pc249b5436.Driver{}
}

func (inst* pc249b54360_dsqlserver_Driver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc249b5436.Driver)
	nop(ie, com)

	


    return nil
}


