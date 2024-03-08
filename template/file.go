// ******************************************************************************
// Matronator © 2024.                                                          *
// ******************************************************************************

package template

type GenericFileObject struct {
	Contents string
	Filename string
	Path     string
}

func NewFileObject(name string, dir string, contents string) *GenericFileObject {
	return &GenericFileObject{
		Contents: contents,
		Filename: name,
		Path:     dir,
	}
}

type Header struct {
	Name     string
	Filename string
	Path     string
}

func FromMap(m map[string]string) *Header {
	return &Header{
		Name:     m["name"],
		Filename: m["filename"],
		Path:     m["path"],
	}
}
