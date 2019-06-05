package handler

import (
	"golang-filestore/common"
	"golang-filestore/util"
	"net/http"
)

func HTTPIntercepter(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			username := r.Form.Get("username")
			token := r.Form.Get("token")

			// 验证登录token是否有效
			if len(username) < 3 || !IsTokenValid(token) {
				resp := util.NewRespMsg(
					int(common.StatusInvalidToken),
					"token无效",
					nil,
				)
				w.Write(resp.JSONBytes())
				return
			}
			h(w, r)
		})
}
