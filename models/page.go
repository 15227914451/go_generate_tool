package models

type Page struct {
	PageNumber int `json:"pageNumber" ` //当前页
	PageSize   int `json:"pageSize" `   //每页显示数量

	Total      int  `json:"total" `   //数据总量
	NextPage   bool `json:"nextPage"` //是否有下一页 true-有 false-无
	NotifyType int  `json:"notifyType"`

	DataList []map[string]interface{} `json:"dataList"` //数据
}
