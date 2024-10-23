package user

import (
	"AI_Server/init/data"
	"context"
	"strings"
	"time"
)

func SetLogoutToken(ctx context.Context, token string, expiration time.Duration) error {
	var builder strings.Builder
	builder.WriteString("logout_")
	builder.WriteString(token)
	if err := data.RDB.SetNX(ctx, builder.String(), "", expiration).Err(); err != nil {
		return err
	}
	return nil
}

func GetLogoutToken(ctx context.Context, token string) (bool, error) {
	var builder strings.Builder
	builder.WriteString("logout_")
	builder.WriteString(token)
	ok, err := data.RDB.Exists(ctx, builder.String()).Result()
	if err != nil {
		return false, err
	}
	return ok == 1, nil
}
