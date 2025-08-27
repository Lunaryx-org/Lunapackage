package shared

import (
	"log"
	"os"

	"github.com/Lunaryx-org/LunaPackage/goimport-migrator/shared/cache"
)

func Map() {
	value, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	exetablePath := string(value)
	cache.Set("pwd", exetablePath)
}
