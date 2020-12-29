package gormgen

import (
	"fmt"
	"text/template"
)

func parseTemplateOrPanic(t string) *template.Template {
	tpl, err := template.New("output_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}

var commonTemplate = parseTemplateOrPanic(fmt.Sprintf(`
package {{.PkgName}}
type FieldData struct {
		Value interface{} %sjson:"value" form:"value"%s
		Symbol string %sjson:"symbol" form:"symbol"%s
	
}
`, "`", "`", "`", "`"))

var outputTemplate = parseTemplateOrPanic(fmt.Sprintf(`
package {{.PkgName}}

import (
	{{range .ImportPkgs}}
		"{{.Pkg}}"
	{{end}}
)

type {{.StructName}}Service struct {	
}

{{$LogName := .LogName}}


	
	// Add add one record
	func ({{.StructNameLitle}}Service *{{.StructName}}Service) Add(db *gorm.DB, entity *models.{{.StructName}})(err error) {
		entity.CreateTime = time.Now()
		entity.UpdateTime = time.Now()
		if err = db.Create(t).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err){{end}}
			
			return err
		}
		return nil
	}

	// Delete delete record
	func ({{.StructNameLitle}}Service *{{.StructName}}Service) Delete(db *gorm.DB, entity *models.{{.StructName}})(err error) {
		if err =  db.Delete(t).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			
			return err
		}
		return nil
	}
	
	// Updates update record
	func ({{.StructNameLitle}}Service *{{.StructName}}Service) Updates(db *gorm.DB, m map[string]interface{}, t *models.{{.StructName}})(err error) {
		m["updateTime"] = time.Now()
		if err = db.Model(&models.{{.StructName}}{}).Where("id = ?",t.ID).Updates(m).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			
			return err
		}
		return nil
	}

	// Get{{.StructName}}All get all record
	func Get{{.StructName}}All(db *gorm.DB)(ret []*models.{{.StructName}},err error){
		if err = db.Find(&ret).Error;err!=nil{
			{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
			
			return ret, err
		}
		return ret, nil
	}
	
	// Get{{.StructName}}Count get count
	func Get{{.StructName}}Count(db *gorm.DB)(ret int64){
		db.Model(&models.{{.StructName}}{}).Count(&ret)
		return ret
	}

	{{$StructName := .StructName}}
	{{$StructNameLitle := .StructNameLitle}}
	
	{{range .OnlyFields}}
		
		// GetBy{{.FieldName}} get one record by {{.FieldName}}
		func ({{$StructNameLitle}}Service *{{$StructName}}Service)GetBy{{.FieldName}}(db *gorm.DB, entiy *models.{{$StructName}})(err error){
			if err = db.First(entiy,"{{.ColumnName}} = ?",entiy.{{.FieldName}}).Error;err!=nil{
				{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
				
				return err
			}
			return nil
		}
		// DeleteBy{{.FieldName}} delete record by {{.FieldName}}
		func ({{$StructNameLitle}}Service *{{$StructName}}Service) DeleteBy{{.FieldName}}(db *gorm.DB, entiy *models.{{$StructName}})(err error) {
			if err= db.Delete(entiy,"{{.ColumnName}} = ?",entiy.{{.FieldName}}).Error;err!=nil{
				{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
				
				return err
				}
			return nil
		}
	{{end}}
	func FindOne{{$StructName}}(db *gorm.DB, queryMap map[string]interface{}, entity *models.{{$StructName}}) (string, error) {
	
	{{range .OptionFields}} 
		if {{.FieldName}}, isExist := queryMap["{{.FieldName}}"]; isExist != nil{
			db = db.Where("{{.ColumnName}}  = ?",{{.FieldName}})
		} 
		
	{{end}}
	result:= db.Find(&entity)
	if result.Error != nil {
		if result.RecordNotFound() {
			return utils.ORGANIZATIONACCOUNTNOTEXIST, nil
		}
		return utils.DATABASEQUERYEXCEPTION, result.Error
	}
	return utils.SUCCESS, nil

	}
	func Find{{.StructName}}List(db *gorm.DB, queryMap map[string]interface{}) ([]models.{{.StructName}}, error) {

		{{.StructNameLitle}}s := make([]models.{{.StructName}}, 0)
	
		{{range .OptionFields}} 
		if {{.FieldName}}, isExist := queryMap["{{.FieldName}}"]; isExist != nil{
			db = db.Where("{{.ColumnName}}  = ?",{{.FieldName}})
		} 
		
		{{end}}
		result := Db.Find(&{{.StructNameLitle}}s)

		return {{.StructNameLitle}}s, result.Error

	}

`))

var outputControllerTemplate = parseTemplateOrPanic(fmt.Sprintf(`
package {{.PkgName}}


import (

	{{range .ImportPkgs}}
	"{{.Pkg}}"
	{{end}}
)
type {{.StructName}}Controller struct {	
}
var {{.StructNameLitle}}Service service.{{.StructName}}Service
{{$LogName := .LogName}}

	
	//  add one record
	func ({{.StructNameLitle}}Controller *GoodController) Add{{.StructName}}(context *gin.Context) {
	resultMap := make(map[string]interface{})
	var {{.StructNameLitle}} models.{{.StructName}}
	err := context.ShouldBind(&{{.StructNameLitle}})
	if err != nil {
		{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return
	}
	if {{.StructNameLitle}} == (models.{{.StructName}}) {
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return

	}
	db := mysqlManager.SqlDb
	err = {{.StructNameLitle}}Service.Add(db, &{{.StructNameLitle}})
	if err != nil {
		{{if $LogName}} {{ $LogName}}.Errorln(err) {{end}}
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
func ({{.StructNameLitle}}Controller *{{.StructName}}Controller) Delete{{.StructName}}(context *gin.Context) {

	resultMap := make(map[string]interface{})
	var {{.StructNameLitle}} models.{{.StructName}}
	db := mysqlManager.SqlDb
	err := context.ShouldBind(&{{.StructNameLitle}})
	if err != nil {
		logs.Logger.Errorln(err)
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return
	}
	
	err := context.ShouldBind(&{{.StructNameLitle}})
	if {{.StructNameLitle}}.ID == "" {
		resultMap["errorCode"] = utils.ILLEGALLOGINPARAMETERS
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.ILLEGALLOGINPARAMETERS)
		utils.ResultJson(context, resultMap)
		return
	}
	
	err = {{.StructNameLitle}}Service.GetByID(db, &{{.StructNameLitle}})
	if err != nil {
		resultMap["errorCode"] = utils.DATABASEDATAEXCEPTION
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.DATABASEDATAEXCEPTION)
		utils.ResultJson(context, resultMap)
		return
	}
	err = {{.StructNameLitle}}Service.DeleteByID(db, &{{.StructNameLitle}})
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
func ({{.StructNameLitle}}Controller *{{.StructName}}Controller) Update{{.StructName}}(context *gin.Context) {
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
	var {{.StructNameLitle}} models.{{.StructName}}
	{{.StructNameLitle}}.ID = updateMap["id"]
	err = {{.StructNameLitle}}Service.GetByID(db, &{{.StructNameLitle}})
	if err != nil {
		logs.Logger.Errorln(err)
		resultMap["errorCode"] = utils.DATABASEDATAEXCEPTION
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.DATABASEDATAEXCEPTION)
		utils.ResultJson(context, resultMap)
		return
	}
	err = {{.StructNameLitle}}Service.Updates(db, updateMap, &{{.StructNameLitle}})
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
func ({{.StructNameLitle}}Controller *{{.StructName}}Controller) FindOne(context *gin.Context) {
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
	var {{.StructNameLitle}} models.{{.StructName}}

	err = service.FindOne{{.StructName}}(db, queryMap, &{{.StructNameLitle}})
	if err != nil {
		logs.Logger.Errorln(err)
		resultMap["errorCode"] = utils.DATABASEDATAEXCEPTION
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.DATABASEDATAEXCEPTION)
		utils.ResultJson(context, resultMap)
		return
	}
	resultMap["{{.StructNameLitle}}_info"] ={{.StructNameLitle}}
	resultMap["errorCode"] = utils.SUCCESS
	resultMap["errorInfo"] = utils.GetResultMessage(context, utils.SUCCESS)
	utils.ResultJson(context, resultMap)
	return
}

//find list
func ({{.StructNameLitle}}Controller *{{.StructName}}Controller) Find{{.StructName}}List(context *gin.Context) {
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
	var {{.StructNameLitle}} models.{{.StructName}}

	{{.StructNameLitle}}List, errPrm := service.Find{{.StructName}}List(db, queryMap)
	if errPrm != nil {
		logs.Logger.Errorln(errPrm)
		resultMap["errorCode"] = utils.DATABASEDATAEXCEPTION
		resultMap["errorInfo"] = utils.GetResultMessage(context, utils.DATABASEDATAEXCEPTION)
		utils.ResultJson(context, resultMap)
		return
	}
	resultMap["{{.StructNameLitle}}_list"] = {{.StructNameLitle}}List
	resultMap["errorCode"] = utils.SUCCESS
	resultMap["errorInfo"] = utils.GetResultMessage(context, utils.SUCCESS)
	utils.ResultJson(context, resultMap)
	return
}


`))
