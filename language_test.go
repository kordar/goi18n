package goi18n

import (
	"log"
	"testing"
)

func TestInitLang(t *testing.T) {
	handler := NewIniHandler("language", "en")

	demo := YamlDemo{}
	handler2 := NewYamlHandler("language", "en", demo)

	InitLang(handler, handler2)
	section := GetSection("en", "system", "ini")
	log.Println("section = ", section.(map[string]string))
	value := GetSectionValue("en", "system", "a", "ini")
	log.Println("value = ", value)
	log.Println("value = ", GetSectionValue("en", "system", "bb", "ini"))

	// **************************
	log.Println(GetSection("en", "demo", "yaml"))
	v := GetSectionValue("zh", "demo", "Age", "yaml")
	log.Println(v.(int))
}

func TestYamlDemo_Path(t *testing.T) {
	demo := YamlDemo{}
	handler := NewYamlHandler("language", "en", demo)
	InitLang(handler)
	log.Println(GetSection("en", "demo", "yaml"))
	value := GetSectionValue("zh", "demo", "Age", "yaml")
	log.Println(value.(int))
}
