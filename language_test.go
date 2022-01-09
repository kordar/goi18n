package goi18n

import (
	"log"
	"testing"
)

func TestInitLang(t *testing.T) {
	handler := NewIniHandler("language", "en")
	InitLang(handler)
	value := GetSectionValue("en", "system", "a")
	log.Println("value = ", value)
	log.Println("value = ", GetSectionValue("en", "system", "bb"))
}

