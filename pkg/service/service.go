package service

import (
	"github.com/bino1490/case-study/pkg/logger"
	"github.com/bino1490/case-study/pkg/repository"
)

type ScheduleService interface {
	GetScheduleByID(channelId string, epochtime string,
		logFields map[string]interface{}) (map[string]interface{}, error)
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

func (s *Service) GetScheduleByID(channelId string, epochtime string,
	logFields map[string]interface{}) (map[string]interface{}, error) {
	logger.LogDebug("Entering Service.GetScheduleByID() ...", logFields)
	return s.repo.GetScheduleByChannelID(channelId, epochtime, logFields)
}
