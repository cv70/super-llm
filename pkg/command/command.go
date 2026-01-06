package command

import (
	"fmt"
)

func AttachExecuteCommand(session, cmd string) string {
	return fmt.Sprintf(`screen -S %s -X stuff "%s\n"`, session, cmd)
}
