package goi18n

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
	"log"
	"path"
)

type IniHandler struct {
	cfgs       map[string]*goconfig.ConfigFile
	defaultCfg *goconfig.ConfigFile
}

func getIniConfigFile(fileName string, moreFiles ...string) *goconfig.ConfigFile {
	cfg, err := goconfig.LoadConfigFile(fileName, moreFiles...)
	if err != nil {
		log.Println(fmt.Sprintf("Please create a language folder in the root directory and create a file %s.ini in it", fileName))
		log.Panicln(err)
	}
	return cfg
}

// NewIniHandler 语言包目录，及默认语言
func NewIniHandler(dir string, lang string) Handler {

	handler := IniHandler{
		cfgs:       make(map[string]*goconfig.ConfigFile),
		defaultCfg: nil,
	}

	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(fmt.Sprintf("Please create a language folder in the root directory and create a dir %s in it", dir))
		log.Panicln(err)
	}
	for _, info := range dirs {
		if info.IsDir() {
			log.Println(fmt.Sprintf("load language file %s", info.Name()))
			if files, err := GetAllFile(path.Join(dir, info.Name()), ".ini"); err == nil {
				if len(files) == 1 {
					handler.cfgs[info.Name()] = getIniConfigFile(files[0])
				}
				if len(files) > 1 {
					handler.cfgs[info.Name()] = getIniConfigFile(files[0], files[1:]...)
				}
			}
		}
	}

	if handler.cfgs[lang] == nil {
		log.Panicln("default lang not found")
	}

	handler.defaultCfg = handler.cfgs[lang]

	return &handler
}

func (ini IniHandler) GetSign() string {
	return "ini"
}

func (ini IniHandler) GetSection(lang string, section string) interface{} {
	cfg := ini.cfgs[lang]
	if cfg == nil {
		cfg = ini.defaultCfg
	}

	if value, err := cfg.GetSection(section); err == nil {
		return value
	} else {
		return map[string]string{}
	}
}

func (ini IniHandler) GetSectionValue(lang string, section string, key string) interface{} {
	cfg := ini.cfgs[lang]
	if cfg == nil {
		cfg = ini.defaultCfg
	}

	if value, err := cfg.GetValue(section, key); err == nil {
		return value
	} else {
		return ""
	}
}
