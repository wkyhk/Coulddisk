package handler

import (
	"net/http"

	"CLOUDDISK/core/internal/logic"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadPrepareHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadPrepareRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFileUploadPrepareLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadPrepare(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
