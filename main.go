package s

import (
	"os"
	"reflect"
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

// set the environment variables from a struct for testing
func Set(st any) {
	val := reflect.ValueOf(st)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()
	for i := range val.NumField() {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		if field.Kind() == reflect.String && field.CanSet() {
			fieldValue := field.String()
			envName := strings.ToUpper(fieldName)
			os.Setenv(envName, fieldValue)
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
