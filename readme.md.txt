docs 命令行在cmd目录下执行命令：
go run main.go -structs Good -input ../service -inputController ../con
troller -imports github.com/jinzhu/gorm,PassServer/logs,PassServer/models,time,PassServer/utils -importsController github.com/jinzhu/gorm,PassServer
/logs,PassServer/models,PassServer/service,PassServer/mysqlManager,PassServer/utils,github.com/gin-gonic/gin  -queryPath ../models -logName logs.Log
ger  -transformErr true