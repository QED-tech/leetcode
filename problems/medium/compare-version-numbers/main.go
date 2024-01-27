package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	r := compareVersion("1.0.0.0.0.0.1", "1.0.0.0.2")
	fmt.Println(r)
}

func compareVersion(version1 string, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")

	if len(v2) > len(v1) {
		v1 = normalizeVersion(v1, len(v2)-len(v1))
	}

	if len(v1) > len(v2) {
		v2 = normalizeVersion(v2, len(v1)-len(v2))
	}

	for i := 0; i < len(v1); i++ {
		ver1, _ := strconv.Atoi(v1[i])
		ver2, _ := strconv.Atoi(v2[i])

		if ver1 > ver2 {
			return 1
		}

		if ver2 > ver1 {
			return -1
		}
	}

	return 0
}

func normalizeVersion(version []string, k int) []string {
	for i := 0; i < k; i++ {
		version = append(version, "0")
	}

	return version
}
