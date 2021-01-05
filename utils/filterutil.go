package utils

import (
	"regexp"
	"strings"

	"github.com/astaxie/beego/context"
)

func IsSkipFilterRouter(ctx *context.Context, method string, skipFilter map[string]map[string]bool) bool {
	path := ctx.Input.URL()
	if strings.HasSuffix(path, "/") && path != "/" {
		path = path[0:(len(path) - 1)]
	}

	//匹配及静态路由
	if methodMap, ok := skipFilter[path]; ok {
		if _, ok := methodMap["Any"]; ok {
			return true
		}

		if exist, ok := methodMap[method]; ok && exist {
			return true
		}
	}

	//匹配及正则路由
	for route, methodMap := range skipFilter {
		if canMatch(ctx, path, route, method, methodMap) {
			return true
		}
	}

	return false
}

func canMatch(ctx *context.Context, webPath, routePath, method string, methodMap map[string]bool) bool {
	routePaths := splitPath(strings.ToLower(routePath))
	webPaths := splitPath(strings.ToLower(webPath))
	if len(routePaths) != len(webPaths) {
		return false
	}

	for i, path := range routePaths {
		isRexRoute := strings.HasPrefix(path, ":")
		if !isRexRoute && path != webPaths[i] {
			return false
		}

		if isRexRoute && ctx.Input.Param(getKeyName(path)) != webPaths[i] {
			return false
		}

	}

	if _, ok := methodMap["Any"]; ok {
		return true
	}

	if exist, ok := methodMap[method]; ok && exist {
		return true
	}

	return false
}

func splitPath(key string) []string {
	index := strings.IndexAny(key, "?")
	if index != -1 {
		key = key[:index]
	}

	key = strings.Trim(key, "/ ")
	if key == "" {
		return []string{}
	}
	return strings.Split(key, "/")
}

func getKeyName(val string) string {
	rex := `^(?P<key>:[a-z]+)[:(].*`
	return regexp.MustCompile(rex).FindStringSubmatch(val)[1]
}
