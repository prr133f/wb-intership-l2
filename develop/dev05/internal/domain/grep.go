package domain

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func (d *Domain) Grep(files []*os.File, pattern string, flags Flags) error {
	if flags.IgnoreCase && !flags.Fixed {
		pattern = "(?i)" + pattern
	} else if flags.IgnoreCase && flags.Fixed {
		pattern = strings.ToLower(pattern)
	}

	rexp, err := regexp.Compile(pattern)
	if err != nil {
		d.Log.Error(err.Error())
		return err
	}

	for _, file := range files {
		if flags.Fixed {
			if err := d.grepFileFixed(file, pattern, flags); err != nil {
				d.Log.Error(err.Error())
				return err
			}
		} else {
			if err := d.grepFilebyRegexp(file, rexp, flags); err != nil {
				d.Log.Error(err.Error())
				return err
			}
		}
	}

	return nil
}

func (d *Domain) grepFileFixed(file *os.File, pattern string, flags Flags) error {
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	var lineCount int

	for sc.Scan() {
		if flags.LineNum {
			lineCount++
		}

		line := sc.Text()
		if flags.IgnoreCase {
			line = strings.ToLower(line)
		}

		if strings.Compare(line, pattern) == 0 {
			colored := color.RedString(line)
			if flags.LineNum {
				wr.WriteString(fmt.Sprintf("%d: %s\n\n", lineCount, colored))
			} else {
				wr.WriteString(fmt.Sprintf("%s\n\n", colored))
			}
		}
	}
	if err := sc.Err(); err != nil {
		d.Log.Error(err.Error())
		return err
	}

	fmt.Println(strings.TrimSuffix(wr.String(), "\n\n"))
	return nil
}

func (d *Domain) grepFilebyRegexp(file *os.File, rexp *regexp.Regexp, flags Flags) error {
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	var lineCount, lineNum int

	var beforeContext [][]byte
	var afterContextCount int
	var lastPrintedLineNum int

	for sc.Scan() {
		lineNum++
		line := sc.Text()
		row := rexp.Find([]byte(line))
		if row != nil {
			if flags.Count {
				lineCount++
			}
			for idx, beforeLine := range beforeContext {
				if flags.LineNum {
					wr.WriteString(fmt.Sprintf("%d: %s\n", lineNum-len(beforeContext)+idx, beforeLine))
				} else {
					wr.WriteString(fmt.Sprintf("%s\n", beforeLine))
				}
			}
			context := strings.Split(line, string(row))
			colored := context[0] + color.RedString(string(row)) + context[1]
			if flags.LineNum {
				wr.WriteString(fmt.Sprintf("%d: %s\n", lineNum, colored))
			} else {
				wr.WriteString(fmt.Sprintf("%s\n", colored))
			}

			beforeContext = beforeContext[:0]
			afterContextCount = max(flags.After, flags.Context)
		} else {
			if flags.Before > 0 || flags.Context > 0 {
				beforeContext = append(beforeContext, []byte(line))
				if len(beforeContext) > max(flags.Before, flags.Context) {
					beforeContext = beforeContext[1:]
				}
			}

			if afterContextCount > 0 && lineNum > lastPrintedLineNum {
				if flags.LineNum {
					wr.WriteString(fmt.Sprintf("%d: %s\n", lineNum, line))
				} else {
					wr.WriteString(fmt.Sprintf("%s\n", line))
				}
				afterContextCount--
			}
		}
	}
	if err := sc.Err(); err != nil {
		d.Log.Error(err.Error())
		return err
	}

	fmt.Println(strings.TrimSuffix(wr.String(), "\n"))

	if flags.Count {
		fmt.Printf("\nNumber of matches: %d\n", lineCount)
	}
	return nil
}
