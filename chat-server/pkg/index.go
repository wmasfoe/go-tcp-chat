package pkg

import (
	"reflect"
	"strings"
)

func JoinStr(args ...string) string {
	var builder strings.Builder

	if val := reflect.ValueOf(args); val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		for _, arg := range args {
			builder.WriteString(arg)
		}
	}

	return builder.String()
}

func PointJoinStr(targetStr *string, args ...string) {
	*targetStr = JoinStr(args...)
}
