package maps

import (
	"sort"

	"github.com/rwxrob/bonzai/fn"
)

// Note to maintainers: This file contains maps that require additional
// arguments and are therefore not able to call simple map functions
// from the mapf package. Please keep simple mapf-able maps in
// simple.go instead.

// Prefix adds a prefix to the string.
func Prefix(in []string, pre string) []string {
	return fn.Map(in, func(i string) string { return pre + i })
}

// Keys returns the keys in lexicographically sorted order.
func Keys[T any](m map[string]T) []string {
	keys := []string{}
	for k, _ := range m {
		keys = append(keys, k)
		sort.Strings(keys)
	}
	return keys
}
