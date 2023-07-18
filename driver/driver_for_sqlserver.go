package driver

import (
	"errors"
	"strconv"
	"strings"

	"gorm.io/driver/sqlserver"
	driver_pkg "gorm.io/driver/sqlserver"

	"github.com/starter-go/libgorm"

	"gorm.io/gorm"
)

// Driver SQLServer 驱动
type Driver struct {

	//starter:component
	_as func(libgorm.Driver) //starter:as(".")

}

func (inst *Driver) _impl() {
	inst._as(inst)
}

// Registration ...
func (inst *Driver) Registration() *libgorm.DriverRegistration {
	return &libgorm.DriverRegistration{
		Name:   "sqlserver",
		Driver: inst,
	}
}

// Open ...
func (inst *Driver) Open(cfg *libgorm.Configuration) (libgorm.Database, error) {
	db, err := inst.openDB(cfg)
	if err != nil {
		return nil, err
	}
	builder := &libgorm.DatabaseBuilder{DB: db}
	return builder.Create(), nil
}

func (inst *Driver) openDB(cfg *libgorm.Configuration) (*gorm.DB, error) {

	//	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930/?database=gorm"
	//	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if cfg == nil {
		return nil, errors.New("libgorm config is nil")
	}

	inst.prepareForDefaultPort(cfg)

	dsnbuilder := &strings.Builder{}
	dsnbuilder.WriteString("sqlserver://")
	dsnbuilder.WriteString(cfg.User)
	dsnbuilder.WriteString(":")
	dsnbuilder.WriteString(cfg.Password)
	dsnbuilder.WriteString("@")
	dsnbuilder.WriteString(cfg.Host)
	dsnbuilder.WriteString(":")
	dsnbuilder.WriteString(strconv.Itoa(cfg.Port))
	dsnbuilder.WriteString("?database=")
	dsnbuilder.WriteString(cfg.Database)
	dsn := dsnbuilder.String()

	dialector := driver_pkg.Open(dsn)
	if dialector == nil {
		return nil, errors.New("driver_sqlserver.Open() return nil")
	}

	return gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
}

func (inst *Driver) prepareForDefaultPort(cfg *libgorm.Configuration) {
	const defport = 1433
	port := cfg.Port
	if port < 1 {
		port = defport
	}
	cfg.Port = port
}
