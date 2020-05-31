package utils

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// ErrExit ...
func ErrExit(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// CopyString ...
func CopyString(s string) string {
	if len(s) <= 0 {
		return ""
	}
	copiedString := s
	return copiedString
}

// Max ...
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min ...
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Releaser ...
type Releaser struct {
	released bool
}

// Released ...
func (r *Releaser) Released() bool {
	return r.released
}

// Release ...
func (r *Releaser) Release() {
	if !r.released {
		r.released = true
	}
}

// Time2String ...
func Time2String(t *time.Time, delim string) string {
	ts := []int{t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second()}
	format := []string{"%d", "%02d", "%02d", "%02d", "%02d", "%02d"}
	tss := make([]string, 6)
	for i, v := range ts {
		tss[i] = fmt.Sprintf(format[i], v)
	}
	return strings.Join(tss, delim)
}

// CloneTimePtr ...
func CloneTimePtr(t *time.Time) (*time.Time, error) {
	clone := time.Time{}

	b, err := t.MarshalBinary()
	if err != nil {
		return nil, err
	}

	if err := clone.UnmarshalBinary(b); err != nil {
		return nil, err
	}

	return &clone, nil
}
