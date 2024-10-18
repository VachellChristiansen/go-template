package module

import (
	"go-template/database"
	"go-template/helper"
)

type CoreModule struct {
	Database database.Database
	Helper   helper.Helper
}

func NewCoreModule() CoreModule {
	return CoreModule{
		Database: database.NewDatabase(),
		Helper:   helper.NewHelper(),
	}
}
