package clusterfan

import (
	"fmt"
	"strings"
)

func prometheus(str store) string {
	out := []string{}
	for origin, temp := range str.Stats() {
		out = append(out, fmt.Sprintf(`clusterfan_temp{host="%s"} %d`, origin, temp))
	}

	return strings.Join(out, "\n")
}
