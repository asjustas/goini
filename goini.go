package goini

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type errorString struct {
	s string
}

type IniRow struct {
	section string
	param   string
	value   string
}

type Ini struct {
	elements []IniRow
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

func getValue(ini *Ini, section string, param string) string {
	for _, element := range ini.elements {
		if element.section == section && element.param == param {
			return element.value
			break
			//fmt.Printf("Results: %v\n", element)
		}
	}

	return ""
}

func Load(conf string) (*Ini, error) {
	file, err := os.Open(conf)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error opening file: %v\n", err))
	}

	defer file.Close()

	arr := new(Ini)

	scanner := bufio.NewScanner(file)
	section := ""

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) > 0 && line[0:1] != ";" {
			if len(line) >= 3 && line[0:1] == "[" /*&& line[(len(line)-1):1] == "]"*/ {
				section = line[1 : len(line)-1]
			} else {
				if strings.Contains(line, "=") {
					KeyVal := strings.Split(line, "=")

					row := IniRow{
						section: section,
						param:   strings.TrimSpace(KeyVal[0]),
						value:   strings.TrimSpace(KeyVal[1]),
					}

					arr.elements = append(arr.elements, row)
				}
			}
		}
	}

	return arr, nil
}

func (ini *Ini) Str(section, param string) string {
	return getValue(ini, section, param)
}

func (ini *Ini) Int(section, param string) int {
	val := getValue(ini, section, param)

	if val == "" {
		return 0
	} else {
		i, err := strconv.Atoi(val)
		if err != nil {
			return 0
		} else {
			return i
		}
	}
}
