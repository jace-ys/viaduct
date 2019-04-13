package cmd

import (
	"encoding/json"
	"os"

	"github.com/jace-ys/viaduct/pkg/utils/format"
	"github.com/jace-ys/viaduct/pkg/utils/log"
)

func setupEnv(f *Flags) error {
	flags, err := mapFlags(f)
	if err != nil {
		return err
	}

	return replaceWithEnv(flags)
}

func mapFlags(f *Flags) (m map[string]interface{}, e error) {
	out, err := json.Marshal(f)
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(out, &m)
	if err != nil {
		return m, err
	}

	return m, nil
}

func replaceWithEnv(flags map[string]interface{}) error {
	for key, value := range flags {
		key = format.CamelToSnakeUnderscore(key)
		env, set := os.LookupEnv(key)
		if set {
			log.Warn().Printf("Environmental variable for %s detected. Overriding specified flag: %s -> %s", key, value, env)
		} else {
			os.Setenv(key, value.(string))
		}
	}

	return nil
}
