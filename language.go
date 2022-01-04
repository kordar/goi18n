package goi18n

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
)

var lang *language

type language struct {
	cfgs       map[string]*goconfig.ConfigFile
	defaultCfg *goconfig.ConfigFile
}

func InitLang(lug string) {
	lang = &language{
		cfgs: make(map[string]*goconfig.ConfigFile),
	}

	if cfg, err := goconfig.LoadConfigFile("language/" + lug + ".ini"); err != nil {
		log.Println(fmt.Sprintf("Please create a language folder in the root directory and create a file %s.ini in it", lang))
		log.Fatalln(err)
	} else {
		lang.defaultCfg = cfg
	}
}

func (l *language) getConfigFile(lug string) *goconfig.ConfigFile {
	if l.cfgs[lug] == nil {
		c, err := goconfig.LoadConfigFile("language/" + lug + ".ini")
		if err != nil {
			return l.defaultCfg
		}
		l.cfgs[lug] = c
	}
	return l.cfgs[lug]
}

func GetSectionValue(lug string, section string, key string) string {
	cfg := lang.getConfigFile(lug)
	if value, err := cfg.GetValue(section, key); err == nil {
		return value
	} else {
		return ""
	}
}

func GetSection(lug string, section string) map[string]string {
	cfg := lang.getConfigFile(lug)
	if value, err := cfg.GetSection(section); err == nil {
		return value
	} else {
		return map[string]string{}
	}
}
