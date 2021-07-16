package repository

import (
	"context"

	"github.com/bino1490/case-study/pkg/config"
	"github.com/bino1490/case-study/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	cbConnStr = config.SrvConfig.GetString(
		"database.mongodb.connectionstring")
	cbTable = config.SrvConfig.GetString(
		"database.mongodb.collection")
	cbDatabase = config.SrvConfig.GetString(
		"database.mongodb.db")
)

type CbRepository struct {
	Collection *mongo.Collection
}

//NewCbRepository to initialize the MonoDB connection
//Conntects with respective Table
func NewCbRepository() *CbRepository {
	logger.BootstrapLogger.Debug("Entering Repository.NewCbRepository() ...")

	clientOptions := options.Client().ApplyURI(cbConnStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.BootstrapLogger.Error(err)
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.BootstrapLogger.Error(err)
		panic(err)
	}
	logger.BootstrapLogger.Debug("Connected to MongoDB!")
	collection := client.Database(cbDatabase).Collection(cbTable)
	return &CbRepository{
		Collection: collection,
	}
}

func (r *CbRepository) GetScheduleByChannelID(channelId string, epochtime string,
	logFields map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

// func (r *CbRepository) GetScheduleByChannelID(channelId string, epochtime string,
// 	logFields map[string]interface{}) (map[string]interface{}, error) {
// 	logger.LogDebug("Entering CbRepository.GetScheduleByChannelID() ...", logFields)
// 	//var schedule entity.ScheduleRequest
// 	queryStr := fmt.Sprintf("SELECT * from `%s` use index (`%s`) where channel_id = '%s'"+
// 		" and start_epoch <= '%s' ORDER BY start_epoch DESC limit 1", cbBucket, ScheduleIndex, channelId, epochtime)

// 	logger.LogDebug("GetScheduleByChannelID formatted query: "+queryStr, logFields)
// 	var params []interface{}

// 	rows, err := r.Cluster.Query(queryStr, &gocb.QueryOptions{
// 		PositionalParameters: params,
// 	})
// 	if err != nil {
// 		logger.LogError("Error occurred when retrieving list of "+
// 			"Schedules records from database "+
// 			channelId, logFields)
// 		logger.LogError(err, logFields)
// 		return nil, entity.ErrDatabaseFailure
// 	}
// 	// Returns list of content entities.
// 	var item map[string]interface{}
// 	var resultMap map[string]interface{}
// 	for rows.Next() {
// 		err = rows.Row(&item)
// 		if err != nil {
// 			logger.LogError("Error occurred when parsing list of "+
// 				"Schedule records from database, Channel Id: "+channelId, logFields)
// 			logger.LogError(err, logFields)
// 			return nil, entity.ErrDatabaseFailure
// 		}

// 		resultMap = item[cbBucket].(map[string]interface{})
// 		delete(resultMap, "channel_id")
// 		delete(resultMap, "source")
// 		delete(resultMap, "source_channel")
// 		//items = append(items, resultMap)
// 	}

// 	//Finding the Next Schedule Epoch Time
// 	queryStr2 := fmt.Sprintf("SELECT * from `%s` use index (`%s`) where channel_id = '%s'"+
// 		" and start_epoch > '%s' limit 1", cbBucket, ScheduleIndex, channelId, epochtime)

// 	logger.LogDebug("GetScheduleByChannelID formatted query: "+queryStr2, logFields)

// 	rows2, err := r.Cluster.Query(queryStr2, &gocb.QueryOptions{
// 		PositionalParameters: params,
// 	})
// 	if err != nil {
// 		logger.LogError("Error occurred when retrieving Epoch Time "+
// 			"Schedules records from database "+
// 			channelId, logFields)
// 		logger.LogError(err, logFields)
// 		return nil, entity.ErrDatabaseFailure
// 	}
// 	// Returns list of content entities.
// 	var item2 map[string]interface{}
// 	var resultMap2 map[string]interface{}
// 	for rows2.Next() {
// 		err = rows2.Row(&item2)
// 		if err != nil {
// 			logger.LogError("Error occurred when parsing list of "+
// 				"Schedule records from database, Channel Id: "+channelId, logFields)
// 			logger.LogError(err, logFields)
// 			return nil, entity.ErrDatabaseFailure
// 		}

// 		resultMap2 = item2[cbBucket].(map[string]interface{})

// 	}
// 	if len(resultMap) != 0 {
// 		resultMap["next_program_start_epoch"] = resultMap2["start_epoch"]
// 	}
// 	return resultMap, nil
// }
