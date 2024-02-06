package v1

type KeyReq struct {
	Key string `form:"key" binding:"required,max=255,min=1"` // 键名
}

type StorePutReq struct {
	Key   string `json:"key" binding:"required,max=255,min=1"`   // 键名
	Value string `json:"value" binding:"required,max=255,min=1"` // 键值
}

type StorePutWithTTLReq struct {
	Key   string `json:"key" binding:"required,max=255,min=1"`
	Value string `json:"value" binding:"required,max=255,min=1"`
	TTL   int64  `json:"ttl" binding:"required"`
}

type ExpireReq struct {
	Key string `json:"key" binding:"required,max=255,min=1"`
	TTL int64  `json:"ttl" binding:"required"`
}
