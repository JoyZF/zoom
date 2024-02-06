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

func (c StoreController) Sync(ctx *gin.Context) {
	if err := store.NewStore().Sync(ctx); err != nil {
		response.WriteResponse(ctx, errors.WithCode(code.GenericServiceErrorCode, err.Error()), nil)
		return
	}
	response.WriteResponse(ctx, nil, nil)
}

func (c StoreController) Stat(ctx *gin.Context) {
	stat := store.NewStore().Stat(ctx)
	response.WriteResponse(ctx, nil, stat)
}

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
