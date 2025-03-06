package utils

import "context"

func GetUserIdFromCtx(ctx context.Context) int32 {
	userId := ctx.Value(SessionUserId)
	if userId == nil {
		return 0
	}
	return userId.(int32)
}

func GetTokenFromCtx(ctx context.Context) string {
	token := ctx.Value(SessionToken)
	if token == nil {
		return ""
	}
	return token.(string)
}
