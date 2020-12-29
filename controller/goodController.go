package controller

import (
	"github.com/jinzhu/gorm"

	"PassServer/logs"

	"PassServer/models"

	"PassServer/service"

	"PassServer/mysqlManager"

	"PassServer/utils"

	"github.com/gin-gonic/gin"
)

type GoodController struct {
}

var goodService service.GoodService

//  add one record
func (goodController *GoodController) AddGood(context *gin.Context) {
	resultMap := make(map[string]interface{})
	var good models.Good
	err := context.ShouldBind(&good)
	if err != nil {
		logs.Logger.Errorln(err)
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return
	}
	if good == (models.Good) {
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return

	}
	db := mysqlManager.SqlDb
	err = goodService.Add(db, &good)
	if err != nil {
		logs.Logger.Errorln(err)
		resultMap["errorCode"] = utils.DATABASEDATAEXCEPTION
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.DATABASEDATAEXCEPTION)
		utils.ResultJson(context, resultMap)
		return

	}
	resultMap["errorCode"] = utils.SUCCESS
	resultMap["errorInfo"] = utils.GetResultMessage(context, utils.SUCCESS)
	utils.ResultJson(context, resultMap)
	return

}

//DeleteGood
func (goodController *GoodController) DeleteGood(context *gin.Context) {

	resultMap := make(map[string]interface{})
	var good models.Good
	db := mysqlManager.SqlDb
	err := context.ShouldBind(&good)
	if err != nil {
		logs.Logger.Errorln(err)
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return
	}

	err := context.ShouldBind(&good)
	if good.ID == "" {
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return
	}

	err = goodService.GetByID(db, &good)
	if err != nil {
		resultMap["errorCode"] = utils.DATABASEDATAEXCEPTION
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.DATABASEDATAEXCEPTION)
		utils.ResultJson(context, resultMap)
		return
	}
	err = goodService.DeleteByID(db, &good)
	if err != nil {
		resultMap["errorCode"] = utils.DATABASEDATAEXCEPTION
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.DATABASEDATAEXCEPTION)
		utils.ResultJson(context, resultMap)
		return
	}
	resultMap["errorCode"] = utils.SUCCESS
	resultMap["errorInfo"] = utils.GetResultMessage(context, utils.SUCCESS)
	utils.ResultJson(context, resultMap)
	return

}

//update
func (goodController *GoodController) UpdateGood(context *gin.Context) {
	resultMap := make(map[string]interface{})
	db := mysqlManager.SqlDb
	updateMap := make(map[string]interface{})

	err := context.BindJSON(&updateMap)
	if err != nil {
		logs.Logger.Errorln(err)
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return
	}
	if len(updateMap) == 0 {
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return
	}
	var good models.Good
	good.ID = updateMap["id"]
	err = goodService.GetByID(db, &good)
	if err != nil {
		logs.Logger.Errorln(err)
		resultMap["errorCode"] = utils.DATABASEDATAEXCEPTION
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.DATABASEDATAEXCEPTION)
		utils.ResultJson(context, resultMap)
		return
	}
	err = goodService.Updates(db, updateMap, &good)
	if err != nil {
		logs.Logger.Errorln(err)
		resultMap["errorCode"] = utils.DATABASEDATAEXCEPTION
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.DATABASEDATAEXCEPTION)
		utils.ResultJson(context, resultMap)
		return
	}
	resultMap["errorCode"] = utils.SUCCESS
	resultMap["errorInfo"] = utils.GetResultMessage(context, utils.SUCCESS)
	utils.ResultJson(context, resultMap)
	return
}

//Find one
func (goodController *GoodController) FindOne(context *gin.Context) {
	resultMap := make(map[string]interface{})
	db := mysqlManager.SqlDb
	queryMap := make(map[string]interface{})

	err := context.BindJSON(&queryMap)
	if err != nil {
		logs.Logger.Errorln(err)
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return
	}
	if len(queryMap) == 0 {
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return
	}
	var good models.Good

	err = service.FindOneGood(db, queryMap, &good)
	if err != nil {
		logs.Logger.Errorln(err)
		resultMap["errorCode"] = utils.DATABASEDATAEXCEPTION
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.DATABASEDATAEXCEPTION)
		utils.ResultJson(context, resultMap)
		return
	}
	resultMap["good_info"] = good
	resultMap["errorCode"] = utils.SUCCESS
	resultMap["errorInfo"] = utils.GetResultMessage(context, utils.SUCCESS)
	utils.ResultJson(context, resultMap)
	return
}

//find list
func (goodController *GoodController) FindGoodList(context *gin.Context) {
	resultMap := make(map[string]interface{})
	db := mysqlManager.SqlDb
	queryMap := make(map[string]interface{})

	err := context.BindJSON(&queryMap)
	if err != nil {
		logs.Logger.Errorln(err)
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return
	}
	if len(queryMap) == 0 {
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return
	}
	var good models.Good

	goodList, errPrm := service.FindGoodList(db, queryMap)
	if errPrm != nil {
		logs.Logger.Errorln(errPrm)
		resultMap["errorCode"] = utils.DATABASEDATAEXCEPTION
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.DATABASEDATAEXCEPTION)
		utils.ResultJson(context, resultMap)
		return
	}
	resultMap["good_list"] = goodList
	resultMap["errorCode"] = utils.SUCCESS
	resultMap["errorInfo"] = utils.GetResultMessage(context, utils.SUCCESS)
	utils.ResultJson(context, resultMap)
	return
}
