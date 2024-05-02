package config

import "os"

func init() {
	os.Setenv("RUNEWIDTH_EASTASIAN", "0")
}
