package utils

import "runtime/debug"

func GetAppVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}
	return "v0.0.0"
}