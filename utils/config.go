package utils

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

var envs = make(map[string]string)

func LoadConfigFromFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		//skip empty line
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		params := strings.SplitN(line, "=", 2)

		if len(params) < 2 {
			continue
		}

		key, val := strings.TrimSpace(params[0]), strings.TrimSpace(params[1])
		envs[key] = val
	}

	envStr, err := json.MarshalIndent(envs, "", "  ")
	LogTrace.Printf("Load config: %s\n", envStr)

	return nil
}

func GetEnv(key string) (bool, string) {
	if _, ok := envs[key]; ok {
		return true, envs[key]
	} else {
		return false, ""
	}
}
