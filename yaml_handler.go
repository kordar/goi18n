package goi18n

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"reflect"
)

type YamlHandler struct {
	defaultLang string
	dir         string
	configs     map[string][]byte
	types       map[string]reflect.Type
	cache       map[string]interface{}
}

func (handler *YamlHandler) GetObj(lang string, section string) interface{} {
	cacheKey := lang + ":" + section
	if handler.cache[cacheKey] != nil {
		return handler.cache[cacheKey]
	}

	typ := handler.types[section]
	if typ == nil {
		return nil
	}

	cfg := reflect.New(typ)
	var params []reflect.Value
	call := cfg.MethodByName("Path").Call(params)
	c := path.Join(handler.dir, lang, call[0].String())
	config := handler.configs[c]
	// 不存在选取默认语言
	if config == nil {
		c = path.Join(handler.dir, handler.defaultLang, call[0].String())
		config = handler.configs[c]
	}

	obj := cfg.Interface()
	if err := yaml.Unmarshal(config, obj); err == nil {
		handler.cache[cacheKey] = obj
		return obj
	}

	return nil
}

func (handler *YamlHandler) GetSign() string {
	return "yaml"
}

func (handler *YamlHandler) GetSection(lang string, section string) interface{} {
	return handler.GetObj(lang, section)
}

func (handler *YamlHandler) GetSectionValue(lang string, section string, key string) interface{} {
	obj := handler.GetObj(lang, section)
	if obj == nil {
		return ""
	}

	value := reflect.ValueOf(obj)
	return value.Elem().FieldByName(key).Interface()
}

func (handler *YamlHandler) AddCfg(cfg YamlPath) {
	key := cfg.Section()
	handler.types[key] = reflect.TypeOf(cfg)
}

func (handler *YamlHandler) AddConfig(fileName string, configs []byte) {
	handler.configs[fileName] = configs
}

func NewYamlHandler(dir string, lang string, cfgs ...YamlPath) *YamlHandler {
	handler := YamlHandler{
		dir:         dir,
		defaultLang: lang,
		configs:     make(map[string][]byte),
		types:       make(map[string]reflect.Type),
		cache:       make(map[string]interface{}),
	}

	for _, cfg := range cfgs {
		handler.AddCfg(cfg)
	}

	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(fmt.Sprintf("Please create a language folder in the root directory and create a dir %s in it", dir))
		log.Panicln(err)
	}

	for _, info := range dirs {
		if info.IsDir() {
			log.Println(fmt.Sprintf("load language file %s", info.Name()))
			if files, err := GetAllFile(path.Join(dir, info.Name()), ".yaml"); err == nil {
				for _, filename := range files {
					// 读取yaml文件到缓存中
					configs, err := ioutil.ReadFile(filename)
					if err != nil {
						log.Panicln(err)
					}
					handler.AddConfig(filename, configs)
				}
			}
		}
	}

	return &handler
}

type YamlPath interface {
	Path() string
	Section() string
}

type YamlDemo struct {
	Name  string `yaml:"Name"`
	Age   int    `yaml:"Age"`
	Sex   string `yaml:"Sex"`
	Class string `yaml:"class"`
}

func (y YamlDemo) Path() string {
	return "demo.yaml"
}

func (y YamlDemo) Section() string {
	return "demo"
}
