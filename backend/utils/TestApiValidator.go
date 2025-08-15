package utils

import (
	"errors"
	"net/url"
)

// ValidateAPIEndpoint checks if the provided API endpoint is valid.
func ValidateAPIEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return errors.New("invalid URL format")
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("only http and https are allowed")
	}

	if u.Host == "" {
		return errors.New("host cannot be empty")
	}

	if u.Port() == "8080" && u.Hostname() == "localhost" {
		return errors.New("localhost:8080 is not allowed")
	}

	// ips, err := net.LookupIP(u.Hostname())
	// if err != nil {
	// 	return errors.New("failed to resolve hostname")
	// }

	// for _, ip := range ips {
	// 	if ip.IsLoopback() || ip.IsPrivate() {
	// 		return errors.New("private or loopback addresses are not allowed")
	// 	}
	// }

	return nil
}
