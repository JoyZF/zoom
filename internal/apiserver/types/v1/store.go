package v1

type StoreGetReq struct {
	Key string `form:"key" binding:"required,max=255,min=1"`
}

type StorePutReq struct {
	Key   string `json:"key" binding:"required,max=255,min=1"`
	Value string `json:"value" binding:"required,max=255,min=1"`
}

type StorePutWithTTLReq struct {
	Key   string `json:"key" binding:"required,max=255,min=1"`
	Value string `json:"value" binding:"required,max=255,min=1"`
	TTL   int64  `json:"ttl" binding:"required"`
}
