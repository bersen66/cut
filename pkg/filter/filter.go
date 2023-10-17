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

		for _, field := range config.Fields {
			if field < len(line) {
				fmt.Printf("%v%v", line[field], config.Separator)
			}
		}
		fmt.Println()
	}

	return nil
}
