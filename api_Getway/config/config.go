package config

import (
    "os"
)

var (
    AuthServiceURL       = getEnv("AUTH_SERVICE_URL", "http://localhost:2121")
    LibraryServiceURL = getEnv("LIBRARY_SERVICE_URL", "http://localhost:50020")
)

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
