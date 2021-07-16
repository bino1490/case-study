package service

import (
	"github.com/bino1490/case-study/pkg/entity"
	"github.com/bino1490/case-study/pkg/logger"
	"github.com/bino1490/case-study/pkg/repository"
)

type DBService interface {
	GetDBRecords(request entity.DBRequest) ([]entity.DBRecord, error)
}

//--  ----
type Service struct {
	repo repository.DbRepository
}

//--  ----
func NewService(r repository.DbRepository) *Service {
	logger.BootstrapLogger.Debug("Entering Service.NewService() ...")
	return &Service{
		repo: r,
	}
}

func (s *Service) GetDBRecords(request entity.DBRequest) ([]entity.DBRecord, error) {
	logger.Logger.Debug("Entering Service.GetDBRecords() ...")
	return s.repo.GetDBRecords(request)
}
