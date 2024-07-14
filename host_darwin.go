//go:build darwin

package uuid

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// hostID returns the uuid returned by `ioreg -rd1 -c IOPlatformExpertDevice`.
// If there is an error running the commad an empty string is returned.
func hostID() (string, error) {
	buf := &bytes.Buffer{}
	err := run(buf, os.Stderr, "ioreg", "-rd1", "-c", "IOPlatformExpertDevice")
	if err != nil {
		return "", err
	}
	id, err := extractID(buf.String())
	if err != nil {
		return "", err
	}
	return trim(id), nil
}

func extractID(lines string) (string, error) {
	for _, line := range strings.Split(lines, "\n") {
		if strings.Contains(line, "IOPlatformUUID") {
			parts := strings.SplitAfter(line, `" = "`)
			if len(parts) == 2 {
				return strings.TrimRight(parts[1], `"`), nil
			}
		}
	}
	return "", fmt.Errorf("Failed to extract 'IOPlatformUUID' value from `ioreg` output.\n%s", lines)
}
