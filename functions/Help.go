package functions

import (
	"fmt"
	"go-telegram-bot/utils"
)

func Help() string {
	return fmt.Sprintf("Estos son los comandos que puedes utilizar: %s", utils.GetFunctions())
}
