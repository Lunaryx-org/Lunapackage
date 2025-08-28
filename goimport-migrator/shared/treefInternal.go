package shared

import (
	"fmt"
	"io/ioutil"

	"github.com/Lunaryx-org/LunaPackage/goimport-migrator/shared/cache"
)

func TreeFprint() {
	workindWD, ok := cache.Get("pwd")
	if ok != true {
		fmt.Println("Stated of cahcing system: ", ok) // DEBUG MESSAGE
		Map()
		// Realized that I was caching it but not retrieving it after
		workindWD, ok = cache.Get("pwd") // Failsafe to retrieve from the cache system
		// after saving it
	}
	fmt.Println(workindWD) // DEBUG MESSAGE
	files, err := ioutil.ReadDir(workindWD)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	for _, f := range files {
		fmt.Printf(f.Name(), "\n")
	}
}
