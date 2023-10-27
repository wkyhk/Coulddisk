package handler

import (
	"errors"
	"net/http"

	"CLOUDDISK/core/helper"
	"CLOUDDISK/core/internal/logic"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadChunkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		//参数必填判断
		if r.PostForm.Get("key") == "" {
			httpx.Error(w, errors.New("key is empty"))
		}
		if r.PostForm.Get("upload_id") == "" {
			httpx.Error(w, errors.New("uploadId is empty"))
		}
		if r.PostForm.Get("part_number") == "" {
			httpx.Error(w, errors.New("partnumber is empty"))
		}
		etag, err := helper.CosPartUpload(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileUploadChunkLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadChunk(&req)
		resp = new(types.FileUploadChunkReply)
		resp.Etag = etag
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
