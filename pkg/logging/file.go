package logging

import (
	"fmt"
	"os"
	"time"

	"xjosiah.com/go-gin/file"
	"xjosiah.com/go-gin/pkg/setting"
)

/*
var (
	LogSavePath = setting.AppSetting.RuntimeRootPath + setting.AppSetting.LogSavePath
	LogSaveName = setting.AppSetting.LogSaveName
	LogFileExt  = setting.AppSetting.LogFileExt

	//	Format根据layout指定的格式返回t代表的时间点的格式化文本表示。
	//	layout定义了参考时间：
	//	Mon Jan 2 15:04:05 -0700 MST 2006

	TimeFormat = setting.AppSetting.TimeFormat
)
*/
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName, time.Now().Format(setting.AppSetting.TimeFormat), setting.AppSetting.LogFileExt)

}

func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err :%v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMKDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile %v", err)
	}

	return f, nil
}

func mkDir() {
	//	Getwd返回一个对应当前工作目录的根路径。如果当前目录可以经过多条路径抵达（因为硬链接），Getwd会返回其中一个。
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
	_, err = os.Create(dir + "/" + getLogFilePath() + getLogFileName())
	if err != nil {
		panic(err)
	}
}
