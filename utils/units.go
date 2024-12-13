package utils

import (
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var mylogSourceDir string

func init() {
	_, file, _, _ := runtime.Caller(0)
	// compatible solution to get gorm source directory with various operating systems
	mylogSourceDir = sourceDir(file)
}

func sourceDir(file string) string {
	dir := filepath.Dir(file)
	dir = filepath.Dir(dir)

	s := filepath.Dir(dir)
	if filepath.Base(s) != "go-wlog" {
		s = dir
	}
	return filepath.ToSlash(s) + "/"
}

// FileWithLineNum 返回当前文件的文件名和行号
// projectFolder 项目文件夹名称，用来去除敏感路径的
func FileWithLineNum(projectFolder string) string {
	pcs := [13]uintptr{}
	// the third caller usually from gorm internal
	len := runtime.Callers(0, pcs[:])
	frames := runtime.CallersFrames(pcs[:len])
	for i := 0; i < len; i++ {
		// second return value is "more", not "ok"
		frame, _ := frames.Next()
		if !strings.HasPrefix(frame.File, mylogSourceDir) &&
			strings.Contains(frame.File, projectFolder) &&
			!strings.Contains(frame.File, "go-wlog") {
			path := string(strconv.AppendInt(append([]byte(frame.File), ':'), int64(frame.Line), 10))
			return RemoveSensitiveInfo(path, projectFolder)
		}
	}

	return ""
}

// IsGormLogCaller 如果是gormlog.go日志，返回true
func IsGormLogCaller() bool {
	pcs := [13]uintptr{}
	// the third caller usually from gorm internal
	len := runtime.Callers(1, pcs[:])
	frames := runtime.CallersFrames(pcs[:len])
	for i := 0; i < len; i++ {
		// second return value is "more", not "ok"
		frame, _ := frames.Next()
		if strings.HasSuffix(frame.File, "gormlog.go") {
			return true
		}
	}

	return false
}

// IsErrorLogCaller 如果是error日志，返回true
func IsErrorLogCaller() bool {
	pcs := [13]uintptr{}
	// the third caller usually from gorm internal
	len := runtime.Callers(1, pcs[:])
	frames := runtime.CallersFrames(pcs[:len])
	for i := 0; i < len; i++ {
		// second return value is "more", not "ok"
		frame, _ := frames.Next()
		if strings.Contains(frame.Func.Name(), "ErrorContext") {
			return true
		}
	}

	return false
}

// RemoveSensitiveInfo 去除敏感信息
// projectFolder 项目文件夹名称，用来去除敏感路径
func RemoveSensitiveInfo(str, projectFolder string) string {
	re := regexp.MustCompile(`.*?` + projectFolder)
	str = re.ReplaceAllString(str, "***")
	re = regexp.MustCompile(`.*?src`)
	str = re.ReplaceAllString(str, "GoSdk")
	return str
}
