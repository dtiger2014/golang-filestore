package handler

import (
	"golang-filestore/common"
	"golang-filestore/util"
	"net/http"
)

// HTTPIntercepter ： http请求拦截器
func HTTPIntercepter(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			username := r.Form.Get("username")
			token := r.Form.Get("token")

			// 验证登录token是否有效
			if len(username) < 3 || !IsTokenValid(token) {
				// token娇艳失败则跳转到直接返回失败提示
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
