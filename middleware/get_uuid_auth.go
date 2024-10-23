package middleware

import (
	"context"
	"fmt"
)
var uuid User_uuid

func GetUserUUIDInAuth(ctx context.Context) string {
	uuid = "user_uuid"
	userUUID := ctx.Value(uuid)
	return fmt.Sprint(userUUID)
}
