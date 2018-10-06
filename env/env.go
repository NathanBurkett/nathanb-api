package env

import (
	"bufio"
	"fmt"
	"github.com/nathanburkett/nathanb-api/app"
	"log"
	"os"
	"strings"
)

const SliceKey = "key"
const SliceValue = "value"

// Must Get the env key's value. Log error and exit if doesn't exist
func Must(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("env variable %s does not exist", key))
	}

	return value
}

func ReadEnv(app *app.Instance) {
	envVars := readEnvFromFile(app.RootDir())
	setEnvFromSlice(envVars)
}

func readEnvFromFile(dir string) []map[string]string {
	var envVars []map[string]string

	file, err := os.Open(fmt.Sprintf("%s/.env", dir))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := strings.SplitN(scanner.Text(), "=", 2)
		count := len(row)

		if count > 2 || count < 2 {
			log.Fatalf("Could not read ENV")
		}

		if _, exists := os.LookupEnv(row[0]); exists {
			continue
		}

		envVars = append(envVars, map[string]string{
			SliceKey:   row[0],
			SliceValue: row[1],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return envVars
}

func setEnvFromSlice(envRowSlice []map[string]string) {
	for i := 0; i < len(envRowSlice); i++ {
		row := envRowSlice[i]
		os.Setenv(row[SliceKey], row[SliceKey])
	}
}
