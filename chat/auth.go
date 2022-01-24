package main

import "net/http"

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		// 何らかの別のエラーが発生
		panic(err.Error())
	} else {
		// 成功、ラップされたハンドラを呼び出す
		h.next.ServeHTTP(w, r)
	}
}

// 任意のhttp.Handlerをラップした*authHandlerを生成
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}
