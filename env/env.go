package env

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const SliceKey = "key"
const SliceValue = "value"

type Reader struct {
	reader io.Reader
}

type envVars []map[string]string

// Must Get the env key's value. Log error and exit if doesn't exist
func Must(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("env variable %s does not exist", key))
	}

	return value
}

func NewReader(ioReader io.Reader) Reader {
	return Reader{
		reader: ioReader,
	}
}

// Read Read and set ENV by io.Reader
func (r Reader) Read() {
	values := r.readEnvFromIo()
	r.setEnvFromSlice(values)
}

func (r Reader) readEnvFromIo() []map[string]string {
	var vars envVars

	scanner := bufio.NewScanner(r.reader)

	for scanner.Scan() {
		vars = r.readEnvRow(scanner, vars)
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	return vars
}

func (r Reader) readEnvRow(scanner *bufio.Scanner, env envVars) envVars {
	row := strings.SplitN(scanner.Text(), "=", 2)

	if _, exists := os.LookupEnv(row[0]); exists {
		return env
	}

	env = append(env, map[string]string{
		SliceKey:   row[0],
		SliceValue: row[1],
	})

	return env
}

func (r Reader) setEnvFromSlice(envSlice []map[string]string) {
	for i := 0; i < len(envSlice); i++ {
		row := envSlice[i]
		os.Setenv(row[SliceKey], row[SliceValue])
	}
}
