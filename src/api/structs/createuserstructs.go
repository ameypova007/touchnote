package structs

type Createuser struct {
	Uuid uint `json:"uuid" binding:"required`
	Name  string `json:"name" binding:"required"`
	EmailId string `json:"emailid" binding:"required"`
	Balance  uint `json:"balance" binding:"required"`
}

type Payload struct{
	Uuid uint `json:"uuid" binding:"required"`
	Priority uint `json:"priority" binding:"required"`
	Amount  uint `json:"amount" binding:"required"`
	Type string `json:"type" binding:"required"`
	Expiry string `json:"expiry" binding:"required"`
}
type Credits struct {
	Activity string `json:"name" binding:"required"`
	Key Payload `json:"payload" binding:"required"`
}

type Payload1 struct{
	Uuid uint `json:"uuid" binding:"required"`
	Amount  uint `json:"amount" binding:"required"`
}

type Debits struct {
	Activity string `json:"name" binding:"required"`
	Key Payload1 `json:"payload" binding:"required"`
}

