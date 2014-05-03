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

// IniRow contains parsed line section, param and value.
type IniRow struct {
	section string
	param   string
	value   string
}

// Ini struct contains parsed ini file lines.
type Ini struct {
	Elements []IniRow
}

func (e *errorString) Error() string {
	return e.s
}

// New returns new error.
func New(text string) error {
	return &errorString{text}
}

// getValue returns value as string.
func getValue(ini *Ini, section string, param string) string {
	for _, element := range ini.Elements {
		if element.section == section && element.param == param {
			return element.value
			break
			//fmt.Printf("Results: %v\n", element)
		}
	}

	return ""
}

// Load function read file to memory and return Ini struct.
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

					arr.Elements = append(arr.Elements, row)
				}
			}
		}
	}

	return arr, nil
}

// Str returns string type value.
func (ini *Ini) Str(section, param string) string {
	return getValue(ini, section, param)
}

// Int returns integer type value.
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
