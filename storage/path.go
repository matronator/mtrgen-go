// ******************************************************************************
// Matronator © 2024.                                                          *
// ******************************************************************************

package storage

import (
	"os"
	"path/filepath"
)

func GetCwd() (string, error) {
	cwd, err := os.Executable()

	return filepath.Dir(cwd), err
}
