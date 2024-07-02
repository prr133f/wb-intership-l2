package domain

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func (d *Domain) Cut(in io.Reader, flags Flags) ([]string, error) {
	var lines []string
	var ans []string
	sc := bufio.NewScanner(in)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	for _, line := range lines {
		if flags.B != "" {
			line, err := cutRange(flags.B, []byte(line))
			if err != nil {
				d.Log.Error(err.Error())
				return nil, err
			}
			ans = append(ans, string(line))
		} else if flags.C != "" {
			line, err := cutRange(flags.C, []rune(line))
			if err != nil {
				d.Log.Error(err.Error())
				return nil, err
			}
			ans = append(ans, string(line))
		} else if flags.F != "" && flags.D != "" {
			if flags.S && !strings.Contains(line, flags.D) {
				continue
			}
			line, err := cutRange(flags.F, strings.Split(line, flags.D))
			if err != nil {
				d.Log.Error(err.Error())
				return nil, err
			}
			ans = append(ans, strings.Join(line, flags.D))
		}
	}

	return ans, nil
}

func cutRange[T any](r string, slice []T) ([]T, error) {
	var ans []T
	rng, err := parseRange(r)
	if err != nil {
		return nil, err
	}

	if len(rng) == 0 {
		return slice, nil
	}

	for _, r := range rng {
		if r.isSolo {
			ans = append(ans, slice[r.first-1])
		} else {
			if len(ans) >= 0 && r.first == 0 {
				ans = ans[:0]
				ans = append(ans, slice[r.first:r.last]...)
			} else if len(ans) >= 0 && r.last == 0 {
				ans = append(ans, slice[r.first-1:]...)
			} else {
				ans = append(ans, slice[r.first-1:r.last]...)
			}
		}
	}
	return ans, nil
}

func parseRange(r string) ([]parsedRange, error) {
	if r == "" {
		return nil, errors.New("empty range")
	}

	solo := strings.Split(r, ",")
	rangeStruct := make([]parsedRange, 0, len(solo))
	var cursor int = -1
	for _, s := range solo {
		var rng parsedRange
		var err error
		if strings.Contains(s, "-") {
			rng.isSolo = false
			pair := strings.Split(s, "-")
			if len(pair) > 2 {
				return nil, errors.New("too many numbers in range")
			} else if pair[0] == "" {
				rng.first = cursor + 1
				rng.last, err = strconv.Atoi(pair[1])
				if err != nil {
					return nil, err
				}
			} else if pair[1] == "" {
				rng.first, err = strconv.Atoi(pair[0])
				if err != nil {
					return nil, err
				}
			} else {
				rng.first, err = strconv.Atoi(pair[0])
				if err != nil {
					return nil, err
				}
				rng.last, err = strconv.Atoi(pair[1])
				if err != nil {
					return nil, err
				}
			}
		} else {
			rng.isSolo = true
			rng.first, err = strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
		}

		rangeStruct = append(rangeStruct, rng)
	}

	return rangeStruct, nil
}
