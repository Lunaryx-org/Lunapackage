/*
* Package to handle the storing of log files
* Based on the name of the function it handles different file formats
	-> ex: Localerrjs works with json files
*/

package err

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Lunaryx-org/Lunapackage/model"
)

func Localerrjs(path string, e error) error {
	if e == nil {
		return nil // Nothing to log
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open the file:	%w", err)
	}
	defer file.Close()

	entry := model.LogObject{
		StructTimeStamp: time.Now(),
		StructError:     e.Error(),
	}

	errorObject, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("Failed to marshal json object:	%w", err)
	}

	_, err = file.Write(append(errorObject, '\n'))
	if err != nil {
		return fmt.Errorf("Failed to write the json object into the file:	%w", err)
	}

	return nil
}
