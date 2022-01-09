package goi18n

import (
	"fmt"
	"io/ioutil"
	"path"
)

var container = map[string]Handler{}

func InitLang(handler Handler, more ...Handler) {
	container[handler.GetSign()] = handler
	for _, h := range more {
		container[h.GetSign()] = h
	}
}

type Handler interface {
	GetSign() string
	GetSection(lang string, section string) interface{}
	GetSectionValue(lang string, section string, key string) interface{}
}

func GetSectionValue(lug string, section string, key string, t string) interface{} {
	if container[t] == nil {
		return ""
	}
	return container[t].GetSectionValue(lug, section, key)
}

func GetSection(lug string, section string, t string) interface{} {
	if container[t] == nil {
		return nil
	}
	return container[t].GetSection(lug, section)
}

// GetAllFile 递归获取指定目录下的所有文件名
func GetAllFile(pathname string, ext string) ([]string, error) {
	var result []string

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := path.Join(pathname, fi.Name())
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err := GetAllFile(fullname, ext)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v", fullname, err)
				return result, err
			}
			result = append(result, temp...)
		} else {
			suffix := path.Ext(fullname)
			if suffix == ext {
				result = append(result, fullname)
			}
		}
	}

	return result, nil
}
