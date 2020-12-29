package service

import (
	"github.com/jinzhu/gorm"

	"PassServer/logs"

	"PassServer/models"

	"time"

	"PassServer/utils"
)

type GoodService struct {
}

// Add add one record
func (goodService *GoodService) Add(db *gorm.DB, entity *models.Good) (err error) {
	entity.CreateTime = time.Now()
	entity.UpdateTime = time.Now()
	if err = db.Create(t).Error; err != nil {
		logs.Logger.Errorln(err)

		return err
	}
	return nil
}

// Delete delete record
func (goodService *GoodService) Delete(db *gorm.DB, entity *models.Good) (err error) {
	if err = db.Delete(t).Error; err != nil {
		logs.Logger.Errorln(err)

		return err
	}
	return nil
}

// Updates update record
func (goodService *GoodService) Updates(db *gorm.DB, m map[string]interface{}, t *models.Good) (err error) {
	m["updateTime"] = time.Now()
	if err = db.Model(&models.Good{}).Where("id = ?", t.ID).Updates(m).Error; err != nil {
		logs.Logger.Errorln(err)

		return err
	}
	return nil
}

// GetGoodAll get all record
func GetGoodAll(db *gorm.DB) (ret []*models.Good, err error) {
	if err = db.Find(&ret).Error; err != nil {
		logs.Logger.Errorln(err)

		return ret, err
	}
	return ret, nil
}

// GetGoodCount get count
func GetGoodCount(db *gorm.DB) (ret int64) {
	db.Model(&models.Good{}).Count(&ret)
	return ret
}

// GetByID get one record by ID
func (goodService *GoodService) GetByID(db *gorm.DB, entiy *models.Good) (err error) {
	if err = db.First(entiy, "id = ?", entiy.ID).Error; err != nil {
		logs.Logger.Errorln(err)

		return err
	}
	return nil
}

// DeleteByID delete record by ID
func (goodService *GoodService) DeleteByID(db *gorm.DB, entiy *models.Good) (err error) {
	if err = db.Delete(entiy, "id = ?", entiy.ID).Error; err != nil {
		logs.Logger.Errorln(err)

		return err
	}
	return nil
}

func FindOneGood(db *gorm.DB, queryMap map[string]interface{}, entity *models.Good) (string, error) {

	if ID, isExist := queryMap["ID"]; isExist != nil {
		db = db.Where("id  = ?", ID)
	}

	if GoodsName, isExist := queryMap["GoodsName"]; isExist != nil {
		db = db.Where("goods_name  = ?", GoodsName)
	}

	if Price, isExist := queryMap["Price"]; isExist != nil {
		db = db.Where("price  = ?", Price)
	}

	if Picture, isExist := queryMap["Picture"]; isExist != nil {
		db = db.Where("picture  = ?", Picture)
	}

	if CreateTime, isExist := queryMap["CreateTime"]; isExist != nil {
		db = db.Where("create_time  = ?", CreateTime)
	}

	if UpdateTime, isExist := queryMap["UpdateTime"]; isExist != nil {
		db = db.Where("update_time  = ?", UpdateTime)
	}

	result := db.Find(&entity)
	if result.Error != nil {
		if result.RecordNotFound() {
			return utils.ORGANIZATIONACCOUNTNOTEXIST, nil
		}
		return utils.DATABASEQUERYEXCEPTION, result.Error
	}
	return utils.SUCCESS, nil

}
func FindGoodList(db *gorm.DB, queryMap map[string]interface{}) ([]models.Good, error) {

	goods := make([]models.Good, 0)

	if ID, isExist := queryMap["ID"]; isExist != nil {
		db = db.Where("id  = ?", ID)
	}

	if GoodsName, isExist := queryMap["GoodsName"]; isExist != nil {
		db = db.Where("goods_name  = ?", GoodsName)
	}

	if Price, isExist := queryMap["Price"]; isExist != nil {
		db = db.Where("price  = ?", Price)
	}

	if Picture, isExist := queryMap["Picture"]; isExist != nil {
		db = db.Where("picture  = ?", Picture)
	}

	if CreateTime, isExist := queryMap["CreateTime"]; isExist != nil {
		db = db.Where("create_time  = ?", CreateTime)
	}

	if UpdateTime, isExist := queryMap["UpdateTime"]; isExist != nil {
		db = db.Where("update_time  = ?", UpdateTime)
	}

	result := Db.Find(&goods)

	return goods, result.Error

}
