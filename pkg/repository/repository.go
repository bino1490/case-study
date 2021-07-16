package repository

import "github.com/bino1490/case-study/pkg/entity"

type DbRepository interface {
	GetDBRecords(request entity.DBRequest) ([]entity.DBRecord, error)
}
