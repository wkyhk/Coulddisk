package handler

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"path"

	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/helper"
	"CLOUDDISK/core/internal/logic"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		err = svcCtx.DB.Where("hash=?", hash).First(rp).Error
		//如果文件不存在往COS中存储
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cosPath, err := helper.CosUpload(r)
			if err != nil {
				fmt.Println(err)
				return
			}
			//往logic中传递req
			req.Name = fileHeader.Filename
			req.Ext = path.Ext(fileHeader.Filename)
			req.Size = fileHeader.Size
			req.Hash = hash
			req.Path = cosPath

		} else if err != nil {
			return
		} else {
			httpx.OkJson(w, &types.FileUploadReply{Identity: rp.Identity, Ext: rp.Ext, Name: rp.Name})
			return
		}

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
