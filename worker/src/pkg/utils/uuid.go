// Path: master/pkg/utils/uuid.go
package utils

import (
    "github.com/google/uuid"
)


func GenerateUUID() string {
    return uuid.New().String()
}

