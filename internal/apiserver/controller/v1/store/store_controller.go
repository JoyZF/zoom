package store

import (
	"github.com/JoyZF/errors"
	"github.com/JoyZF/zoom/internal/apiserver/service/store"
	v1 "github.com/JoyZF/zoom/internal/apiserver/types/v1"
	"github.com/JoyZF/zoom/internal/pkg/code"
	"github.com/JoyZF/zoom/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

type StoreController struct {
}

func NewStoreController() StoreController {
	return StoreController{}
}

// Get
//
//	@Summary	get key
//	@Produce	json
//	@Param		key	query		string						true	"键名"
//	@Success	200	{object}	response.SuccessResponse	"成功"
//	@Failure	400	{object}	response.ErrResponse		"失败"
//	@Router		/v1/store [get]
func (c StoreController) Get(ctx *gin.Context) {
	req := v1.KeyReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.ParamsError, err.Error()), nil)
		return
	}
	val, err := store.NewStore().Get(ctx, req.Key)
	if err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.SystemError, err.Error()), nil)
		return
	}
	response.WriteResponse(ctx, nil, val)
}

// Put
//
//	@Summary	put kv
//	@Produce	json
//	@Param		key		body		string						true	"键名"
//	@Param		value	body		string						true	"键值"
//	@Success	200		{object}	response.SuccessResponse	"成功"
//	@Failure	400		{object}	response.ErrResponse		"失败"
//	@Router		/v1/store [put]
func (c StoreController) Put(ctx *gin.Context) {
	req := v1.StorePutReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.ParamsError, err.Error()), nil)
		return
	}

	err := store.NewStore().Put(ctx, &req)
	if err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.GenericServiceErrorCode, err.Error()), nil)
		return
	}
	response.WriteResponse(ctx, nil, nil)
}

// PutWithTTL
//
//	@Summary	put kv
//	@Produce	json
//	@Param		key		body		string						true	"键名"
//	@Param		value	body		string						true	"键值"
//	@Param		ttl		body		int64						true	"过期时间"
//	@Success	200		{object}	response.SuccessResponse	"成功"
//	@Failure	400		{object}	response.ErrResponse		"失败"
//	@Router		/v1/store [put]
func (c StoreController) PutWithTTL(ctx *gin.Context) {
	req := v1.StorePutWithTTLReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.ParamsError, err.Error()), nil)
		return
	}

	err := store.NewStore().PutWithTTL(ctx, &req)
	if err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.GenericServiceErrorCode, err.Error()), nil)
		return
	}
	response.WriteResponse(ctx, nil, nil)
}

// Delete
//
//	@Summary	delete key
//	@Produce	json
//	@Param		key	query		string						true	"键名"
//	@Success	200	{object}	response.SuccessResponse	"成功"
//	@Failure	400	{object}	response.ErrResponse		"失败"
//	@Router		/v1/store [delete]
func (c StoreController) Delete(ctx *gin.Context) {
	req := v1.KeyReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.ParamsError, err.Error()), nil)
		return
	}

	err := store.NewStore().Delete(ctx, &req)
	if err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.GenericServiceErrorCode, err.Error()), nil)
		return
	}
	response.WriteResponse(ctx, nil, nil)
}

// TTL
//
//	@Summary	get key ttl
//	@Produce	json
//	@Param		key	query		string						true	"键名"
//	@Success	200	{object}	response.SuccessResponse	"成功"
//	@Failure	400	{object}	response.ErrResponse		"失败"
//	@Router		/v1/store/ttl [get]
func (c StoreController) TTL(ctx *gin.Context) {
	req := v1.KeyReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.ParamsError, err.Error()), nil)
		return
	}

	ttl, err := store.NewStore().TTL(ctx, &req)
	if err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.GenericServiceErrorCode, err.Error()), nil)
		return
	}
	response.WriteResponse(ctx, nil, ttl)
}

// Sync
//
//	@Summary	sync data to file
//	@Produce	json
//	@Success	200	{object}	response.SuccessResponse	"成功"
//	@Failure	400	{object}	response.ErrResponse		"失败"
//	@Router		/v1/store/sync [get]
func (c StoreController) Sync(ctx *gin.Context) {
	if err := store.NewStore().Sync(ctx); err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.GenericServiceErrorCode, err.Error()), nil)
		return
	}
	response.WriteResponse(ctx, nil, nil)
}

// Stat
//
//	@Summary	get db stat
//	@Produce	json
//	@Success	200	{object}	response.SuccessResponse	"成功"
//	@Failure	400	{object}	response.ErrResponse		"失败"
//	@Router		/v1/store/stat [get]
func (c StoreController) Stat(ctx *gin.Context) {
	stat := store.NewStore().Stat(ctx)
	response.WriteResponse(ctx, nil, stat)
}

// Exist
//
//	@Summary	return key is exist
//	@Produce	json
//	@Param		key	query		string						true	"键名"
//	@Success	200	{object}	response.SuccessResponse	"成功"
//	@Failure	400	{object}	response.ErrResponse		"失败"
//	@Router		/v1/store/exist [get]
func (c StoreController) Exist(ctx *gin.Context) {
	req := v1.KeyReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.ParamsError, err.Error()), nil)
		return
	}
	exist, err := store.NewStore().Exist(ctx, req.Key)
	if err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.GenericServiceErrorCode, err.Error()), nil)
		return
	}
	response.WriteResponse(ctx, nil, exist)
}

// Expire
//
//	@Summary	set key expire
//	@Produce	json
//	@Param		key	query		string						true	"键名"
//	@Param		ttl	query		int64						true	"过期时间 单位秒"
//	@Success	200	{object}	response.SuccessResponse	"成功"
//	@Failure	400	{object}	response.ErrResponse		"失败"
//	@Router		/v1/store/expire [get]
func (c StoreController) Expire(ctx *gin.Context) {
	req := v1.ExpireReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.ParamsError, err.Error()), nil)
		return
	}

	err := store.NewStore().Expire(ctx, &req)
	if err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.GenericServiceErrorCode, err.Error()), nil)
		return
	}
	response.WriteResponse(ctx, nil, nil)
}
