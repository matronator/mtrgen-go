// ******************************************************************************
// Matronator © 2024.                                                          *
// ******************************************************************************

package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Store struct {
	Templates map[string]string `json:"templates"`
}

type Storage struct {
	HomeDir     string
	TempDir     string
	TemplateDir string
	CacheDir    string
	Store       *Store
	storePath   string
}

func New() *Storage {
	userHomeDir, _ := os.UserHomeDir()
	userTempDir := os.TempDir()
	userCacheDir, _ := os.UserCacheDir()

	homeDir := filepath.Join(userHomeDir, ".mtrgen")
	tempDir := filepath.Join(userTempDir, "mtrgen")
	templateDir := filepath.Join(homeDir, "templates")
	cacheDir := filepath.Join(userCacheDir, "mtrgen")
	storePath := filepath.Join(homeDir, "templates.json")

	_ = os.MkdirAll(tempDir, os.ModePerm)
	_ = os.MkdirAll(templateDir, os.ModePerm)
	_ = os.MkdirAll(cacheDir, os.ModePerm)

	store := createStoreFile(storePath)

	return &Storage{
		HomeDir:     homeDir,
		TempDir:     tempDir,
		TemplateDir: templateDir,
		CacheDir:    cacheDir,
		storePath:   storePath,
		Store:       &store,
	}
}

func (s *Storage) SaveTemplate(name string, path string) error {
	templateFile, err := os.ReadFile(templatePath(path))

	if errors.Is(err, os.ErrNotExist) {
		return err
	}

	originalFilename := filepath.Base(templatePath(path))
	newLocation := filepath.Join(s.TemplateDir, originalFilename)

	file, err := os.Open(newLocation)

	if err == nil {
		_ = file.Close()
		newLocation = filepath.Join(s.TemplateDir, name+"_"+originalFilename)
	}

	err = os.WriteFile(newLocation, templateFile, os.ModePerm)
	if err != nil {
		return err
	}

	err = s.AddEntry(name, newLocation)

	return err
}

func (s *Storage) RemoveTemplate(name string) error {
	path := s.Store.Templates[name]
	_ = os.Remove(path)

	delete(s.Store.Templates, name)
	b, err := json.Marshal(s.Store)

	if err != nil {
		return err
	}

	err = os.WriteFile(s.storePath, b, os.ModePerm)

	return err
}

func (s *Storage) ListEntries() map[string]string {
	return s.Store.Templates
}

func (s *Storage) GetEntry(name string) string {
	return s.Store.Templates[name]
}

func (s *Storage) AddEntry(name string, templatePath string) error {
	s.Store.Templates[name] = templatePath
	b, err := json.Marshal(s.Store)

	if err != nil {
		return err
	}

	err = os.WriteFile(s.storePath, b, os.ModePerm)

	return err
}

func FileToObject[T any](path string, object T, defaultContent []byte) T {
	file, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var (
		b   []byte
		err error
	)

	contents, _ := os.ReadFile(path)

	if contents != nil {
		err = json.Unmarshal(contents, &object)
	} else {
		b = make([]byte, 1024)
		_, _ = file.Read(b)
		err = json.Unmarshal(b, &object)
	}

	if err != nil {
		_ = file.Truncate(0)
		_, _ = file.Write(defaultContent)
		_, _ = file.Seek(0, 0)
		_ = json.Unmarshal(defaultContent, &object)
	}

	return object
}

func templatePath(path string) string {
	cwd, _ := GetCwd()
	return filepath.Join(cwd, path)
}

func createStoreFile(storePath string) Store {
	var store Store

	store = FileToObject[Store](storePath, store, []byte(`{"templates": []}`))

	if store.Templates == nil {
		store.Templates = make(map[string]string)
	}

	return store
}
