package parser

import (
	"go/ast"
	"path/filepath"
	"strings"
)

func (p *ClassParser) getDisplayPackageName(pack *ast.Package) string {
	var path string
	for filename := range pack.Files {
		path = filepath.Dir(filename)
		break
	}

	path = strings.TrimPrefix(path, getLastSecondDirectory(p.dirCommonPrefix))
	path = toDisplayPath(path)
	// path name, package name equals, simplify name
	if path == pack.Name {
		return pack.Name
	}

	return path + ":" + pack.Name
}

func getRealPackageName(displayPackageName string) string {
	parts := strings.Split(displayPackageName, ":")
	return parts[len(parts)-1]
}

func getCommonPrefix(dirs []string) string {
	if len(dirs) == 0 {
		return ""
	}

	if len(dirs) == 1 {
		return dirs[0]
	}

	l := stringsMinLength(dirs)
	i := 0
	for ; i < l; i++ {
		c := dirs[0][i]
		for _, s := range dirs[1:] {
			if s[i] != c {
				return dirs[0][:i]
			}
		}
	}

	return dirs[0][:i]
}

// returned dir ends with /
func getLastSecondDirectory(dir string) string {
	if dir == "" { // invalid dir
		return "/"
	}

	// remove trailing /
	if dir[len(dir)-1] == '/' {
		dir = dir[:len(dir)-1]
	}

	parts := strings.Split(dir, "/")
	parts = parts[:len(parts)-1]
	if len(parts) == 1 && parts[0] == "" {
		return "/"
	}
	return strings.Join(parts, "/") + "/"
}

func stringsMinLength(dirs []string) int {
	if len(dirs) == 0 {
		return 0
	}

	minLength := len(dirs[0])
	for _, s := range dirs[1:] {
		if minLength > len(s) {
			minLength = len(s)
		}
	}
	return minLength
}

// plantuml does not support '/' in name
// use '.', Java package syntax
const newSep = "."

func toDisplayPath(path string) string {
	if path == "" {
		return ""
	}
	p := strings.Replace(path, "/", newSep, -1)
	return p
}
