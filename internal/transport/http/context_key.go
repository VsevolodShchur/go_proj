package http

import contextkey "proj/pkg/context_key"

const (
	userIDParam = "userId"
	noteIDParam = "noteId"
)

var (
	userIDCtxKey = contextkey.CtxKey(userIDParam)
	noteIDCtxKey = contextkey.CtxKey(noteIDParam)
)
