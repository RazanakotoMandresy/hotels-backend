package middleware

import (
	"context"
	"fmt"
)

func GetUserUUIDInAuth(ctx context.Context) string {
	var uuid User_uuid
	uuid = "user_uuid"
	userUUID := ctx.Value(uuid)
	return fmt.Sprint(userUUID)
}
