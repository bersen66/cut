package filter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run(config *Config) error {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), config.Delimeter)
		if len(line) > config.Field {
			fmt.Println(line[config.Field-1])
		}
	}

	return nil
}
