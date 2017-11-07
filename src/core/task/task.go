package task

type Data struct {
	DataId         int    `json:"DataId"`         // 基本数据id
	DataName       string `json:"DataName"`       // 基本数据名称
	DataStatus     int    `json:"DataStatus"`     // 基本数据状态
	CategoryId     int    `json:"CategoryId"`     // 基本数据分类
	UserId         int    `json:"UserId"`         // 基本数据所属用户
	DataCharSet    string `json:"DataCharSet"`    // 基本数据编码
	DataType       string `json:"DataType"`       // 基本数据类型
	DataAddress    string `json:"DataAddress"`    // 基本数据路径
	DataTitle      string `json:"DataTitle"`      // 基本数据标题
	DataAbstract   string `json:"DataAbstract"`   // 基本数据摘要
	DataDesc       string `json:"DataDesc"`       // 基本数据描述
	DataModifyTime string `json:"DataModifyTime"` // 基本数据修改时间
	DataCreateTime string `json:"DataCreateTime"` // 基本数据创建时间
}

type Task struct {
	Tpye               string `json:"Type"` //the type of the task
	AgentId            string `json:"AgentId"`
	BuyerId            string `json:"BuyerId"`
	OrderId            string `json:"OrderId"`
	DataId             string `json:"DataId"`
	Data               Data   `json:"Data"`
	EncryptDataKey     string `json:"EncryptDataKey"`
	EncryptEnvelopeKey string `json:"EncryptEnvelopeKey"`
}
