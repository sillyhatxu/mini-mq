package version

import (
	"fmt"
)

var Version = Conf{Version: 1.0, Description: "mini message queue"}

type Conf struct {
	Version     float64
	Description string
}

func GetCurrentVersion() string {
	return fmt.Sprintf("Version : %b", Version.Version)
}
