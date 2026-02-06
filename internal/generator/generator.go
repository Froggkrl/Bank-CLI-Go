package generator

import (
	"BankCLI/pkg/models"
	"fmt"
)

func GenerateID() string {
	return fmt.Sprintf("ACC%d", len(models.Accounts)+1)
}
