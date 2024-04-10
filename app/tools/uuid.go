package tools

import (
	"fmt"

	"github.com/google/uuid"
)

func GetUUID() string {
	id := uuid.New() // 默认v4
	fmt.Printf("uuid:%s,version:%s", id.String(), id.Version().String())
	return id.String()
}
