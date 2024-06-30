package domain

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func (d *Domain) Sort(in *os.File, out *os.File, flags Flags) error {
	sc := bufio.NewScanner(in)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		d.Log.Error(err.Error())
		return err
	}

	if flags.U {
		lines = revomeDublicates(lines)
	}

	sort.Slice(lines, func(i, j int) bool {
		line1 := lines[i]
		line2 := lines[j]

		if flags.K > 0 {
			words1 := strings.Fields(line1)
			words2 := strings.Fields(line2)

			if len(words1) >= flags.K && len(words2) >= flags.K {
				line1 = words1[flags.K-1]
				line2 = words2[flags.K-1]
			}
		}

		return compareStrings(line1, line2, flags)
	})

	for _, line := range lines {
		_, err := out.WriteString(line)
		if err != nil {
			d.Log.Error(err.Error())
			return err
		}
	}

	return nil
}

func removeNonNumeric(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func revomeDublicates(s []string) []string {
	set := make(map[string]struct{})
	for _, line := range s {
		set[line] = struct{}{}
	}
	s = make([]string, 0, len(set))
	for line := range set {
		s = append(s, line)
	}
	return s
}

func compareStrings(str1, str2 string, flags Flags) bool {
	if flags.N {
		stripped1 := strings.Trim(removeNonNumeric(str1), " ")
		stripped2 := strings.Trim(removeNonNumeric(str2), " ")

		if stripped1 == "" || stripped2 == "" {
			if flags.R {
				return str1 > str2
			}
			return str1 < str2
		}
		num1, err1 := strconv.Atoi(strings.Fields(stripped1)[0])
		num2, err2 := strconv.Atoi(strings.Fields(stripped2)[0])

		if err1 == nil && err2 == nil {
			if flags.R {
				return num1 > num2
			}
			return num1 < num2
		}
	}

	if flags.R {
		return str1 > str2
	}
	return str1 < str2
}
