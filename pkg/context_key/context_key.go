package contextkey

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CtxKey string

func URLParamCtx(param string, key CtxKey) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := chi.URLParam(r, param)
			if userID == "" {
				return
			}
			ctx := context.WithValue(r.Context(), key, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetFormCtx(r *http.Request, ctxKey CtxKey) string {
	key, ok := r.Context().Value(ctxKey).(string)
	if !ok {
		return ""
	}
	return key
}
