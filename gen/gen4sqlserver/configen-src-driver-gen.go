package gen4sqlserver
import (
    p3ecabb6de "github.com/starter-go/module-gorm-sqlserver/driver"
     "github.com/starter-go/application"
)

// type p3ecabb6de.Driver in package:github.com/starter-go/module-gorm-sqlserver/driver
//
// id:com-3ecabb6dee1f6295-driver-Driver
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-Driver
// alias:
// scope:singleton
//
type p3ecabb6dee_driver_Driver struct {
}

func (inst* p3ecabb6dee_driver_Driver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3ecabb6dee1f6295-driver-Driver"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-Driver"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3ecabb6dee_driver_Driver) new() any {
    return &p3ecabb6de.Driver{}
}

func (inst* p3ecabb6dee_driver_Driver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3ecabb6de.Driver)
	nop(ie, com)

	


    return nil
}


