// ******************************************************************************
// Matronator © 2024.                                                          *
// ******************************************************************************

package registry

import (
	"path/filepath"

	"mtrgen/storage"
)

type Profile struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Registry struct {
	Profile  Profile
	Filepath string
}

func New() *Registry {
	s := storage.New()

	profilePath := filepath.Join(s.HomeDir, "profile.json")

	var p Profile

	storage.FileToObject[Profile](profilePath, p, []byte(`{"username": null, "token": null}`))

	return &Registry{
		Profile:  p,
		Filepath: profilePath,
	}
}
