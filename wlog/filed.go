package wlog

import (
	"context"
	"fmt"
	"github.com/cnWangFei/go-Wlog/utils"
	"reflect"
	"sort"
)

type field struct {
	group string
}

func ctxToMap(ctx context.Context) (m map[string]interface{}) {
	if ctx != nil {
		keys := getKeys(ctx) // 获取context中的key
		m = make(map[string]interface{}, len(keys)+4)
		// 将context中的内容拷贝到map
		for _, key := range keys {
			if key != Group {
				m[key] = ctx.Value(key)
			}
		}
		// 获取分组名
		m[Group] = getCtxGroup(ctx)
	} else {
		m = make(map[string]interface{}, 4)
	}

	// 打印调用者
	if reportCaller || utils.IsErrorLogCaller() {
		m["caller"] = utils.FileWithLineNum(callerProjectFolder)
		// 如果是gormlog.go日志
		if utils.IsGormLogCaller() && ctx != nil {
			elapsed := ctx.Value("elapsed")
			m["elapsed"] = elapsed
			rows := ctx.Value("rows")
			m["rows"] = rows
		}
	}
	return
}

// getCtxGroup 从ctx获取分组名
func getCtxGroup(ctx context.Context) (name string) {
	n := ctx.Value(Group)
	if n == nil {
		return
	}
	if reflect.TypeOf(n).Name() != "string" {
		return
	}
	name = n.(string)
	return
}

// getMapGroup 从map获取分组名
func getMapGroup(m map[string]interface{}) (name string) {
	if value, ok := m[Group]; ok {
		if value == nil {
			return
		}
		if reflect.TypeOf(value).Name() != "string" {
			return
		}
		name = value.(string)
	}
	return
}

// mapToStr 将map转换为字符串
func mapToStr(m map[string]interface{}) (str string) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		str = fmt.Sprintf("%v%v : %v\n", str, k, m[k])
	}
	return str
}

// getKeys 递归获取context中的key
func getKeys(ctx context.Context) (keys []string) {
	v := reflect.ValueOf(ctx)
	keys = getCtxKeys(v)
	return
}

// getCtxKeys 递归获取context中的key
func getCtxKeys(v reflect.Value) (keys []string) {
	tp := v.Type().String()
	if tp != "context.valueCtx" && tp != "*context.valueCtx" {
		return
	}
	if tp == "*context.valueCtx" {
		v = v.Elem()
	}
	key := v.FieldByName("key")
	//val := v.FieldByName("val")
	keys = append(keys, fmt.Sprintf("%v", key))
	keys = append(keys, getCtxKeys(v.Field(0).Elem())...)
	return
}
