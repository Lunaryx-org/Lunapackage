package err

import (
	"fmt"
)

func ReturnERRmessage(err string) string {

	return fmt.Sprintf("err %s", err)
}
