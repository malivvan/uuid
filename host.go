package uuid

import (
	"fmt"
)

// HostID returns the platform specific machine id of the current host OS.
// Regard the returned id as "confidential" and consider using ProtectedID() instead.
func HostID() (string, error) {
	id, err := hostID()
	if err != nil {
		return "", fmt.Errorf("hostID: %v", err)
	}
	return id, nil
}

// ProtectedHostID returns a hashed version of the machine ID in a cryptographically secure way,
// using a fixed, application-specific key.
// Internally, this function calculates HMAC-SHA256 of the application ID, keyed by the machine ID.
func ProtectedHostID(appID string) (string, error) {
	id, err := HostID()
	if err != nil {
		return "", fmt.Errorf("machineid: %v", err)
	}
	return protect(appID, id), nil
}
