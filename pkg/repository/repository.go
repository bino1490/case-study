package repository

type DbRepository interface {
	GetScheduleByChannelID(channelId string, epochtime string,
		logFields map[string]interface{}) (map[string]interface{}, error)
}
