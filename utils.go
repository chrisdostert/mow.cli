package cli

import (
	"flag"
	"os"
	"strings"
)

func setFromEnv(into flag.Value, envVars string) {
	isMulti := false
	multiValued, ok := into.(multiValued)
	if ok {
		isMulti = multiValued.IsMultiValued()
	}

	if len(envVars) > 0 {
		for _, rev := range strings.Split(envVars, " ") {
			ev := strings.TrimSpace(rev)
			if len(ev) == 0 {
				continue
			}

			v := os.Getenv(ev)
			if len(v) == 0 {
				continue
			}
			if !isMulti {
				if err := into.Set(v); err == nil {
					return
				}
				continue
			}

			vs := strings.Split(v, ",")
			if err := multiValued.SetMulti(vs); err == nil {
				return
			}
		}
	}
}
