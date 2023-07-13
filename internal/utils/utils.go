package utils

import (
	"bufio"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func ReadPassDB(pass *string) {
	var line string
	file, err := os.Open("../config.txt")
	if err != nil {
		logrus.Warn("Error while reading DB password from config.txt")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if strings.Contains(line, "cmActDb_pass") {
			*pass = strings.Split(line, "=")[1]
		}
	}
}
