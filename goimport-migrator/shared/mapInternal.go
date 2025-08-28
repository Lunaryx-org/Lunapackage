package shared

import (
	"fmt"
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
	fmt.Println(value) // DEBUG MESSAGE
	cache.Set("pwd", exetablePath)
	fmt.Println("Sucessfully cached working directory") // DEBUG MESSAGE

	//After running the process both debug messages work
}
