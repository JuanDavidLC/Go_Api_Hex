package validators

import "strings"

func IsFieldEmpty(field string) bool {

	return strings.TrimSpace(field) == ""

}
