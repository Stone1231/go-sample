package envex

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test_envStr(t *testing.T) {

	var envKeyValues = os.Environ()

	for _, env := range envKeyValues {

		ss := strings.SplitN(env, "=", 2)
		if len(ss) < 2 {
			//act as sentinel
			continue
		}

		fmt.Printf("%v:%v\n", ss[0], ss[1])
	}

}
