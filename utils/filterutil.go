package utils

import (
	"strings"
)

func IsSkipFilterRouter(path, method string, skipFilter map[string]map[string]bool) bool {
	if strings.HasSuffix(path, "/") && path != "/" {
		path = path[0:(len(path) - 1)]
	}

	if methodMap, ok := skipFilter[path]; ok {
		if _, ok := methodMap["Any"]; ok {
			return true
		}

		if exist, ok := methodMap[method]; ok && exist {
			return true
		}
	}

	return false
}
