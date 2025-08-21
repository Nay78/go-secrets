package s

import (
	"os"
	"strings"
)

// type AvailableSecrets struct {
// 	SurfaceManager string
// }

type Msql struct {
	IP   string
	User string
	Pass string
	DB   string
}

func get(envs ...string) []string {
	var result []string
	for _, env := range envs {
		if r, exists := os.LookupEnv(env); !exists {
			panic("Environment variable " + env + " is not set")
		} else {
			result = append(result, r)
		}
	}
	return result
}

// Set parses a .env-style string and sets each key-value pair as an environment variable.
//
// Example:
//
//	Set(`FOO=bar
//	BAZ=qux`)
func Set(envs string) {
	lines := strings.Split(strings.TrimSpace(envs), "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			os.Setenv(key, value)
		}
	}
}

func SurfaceManagerDB(envs ...string) *Msql {
	// envs := []string{
	// 	"SURFACE_MANAGER_IP",
	// 	"SURFACE_MANAGER_USER",
	// 	"SURFACE_MANAGER_PASS",
	// 	"SURFACE_MANAGER_DB",
	// }

	result := get(
		"SURFACE_MANAGER_IP",
		"SURFACE_MANAGER_USER",
		"SURFACE_MANAGER_PASS",
		"SURFACE_MANAGER_DB",
	)

	return &Msql{
		IP:   result[0],
		User: result[1],
		Pass: result[2],
		DB:   result[3],
	}

	// (s.Smip, s.Smuser, s.Smpass, s.Smdb)
}
