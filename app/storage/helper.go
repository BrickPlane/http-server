package storage

import (
	"fmt"
	"strings"
)

func HelperForUpdate(data map[string]interface{}) string {
	query := make([]string, 0, len(data))
	for key := range data {
		query = append(query, fmt.Sprintf("%[1]s=:%[1]s", key))
	}
	return strings.Join(query, ", ")
}
