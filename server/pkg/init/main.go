package initFunc

import (
	"fmt"
	"os"
)

func CreateEnvFile() error {
	envContent := `DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=username
DB_PASSWORD=secretpassword
DB_NAME=mydatabase
DB_SSLMODE=disable
	`

	file, err := os.Create(".env")
	if err != nil {
		return fmt.Errorf("failed to create .env file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(envContent)
	if err != nil {
		return fmt.Errorf("failed to write to .env file: %w", err)
	}

	fmt.Println(".env file created successfully")
	return nil
}
