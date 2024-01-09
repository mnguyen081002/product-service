package utils

import "github.com/hashicorp/go-uuid"

func IsValidUUID(u string) bool {
	_, err := uuid.ParseUUID(u)
	if err != nil {
		return false
	}
	return true
}
