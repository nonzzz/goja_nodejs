package path

import (
	"github.com/dop251/goja"
)

type Win32Path struct {
	Path
}

func (wp *Win32Path) Basename(p string, suffix ...string) string {
	return ""
}

func (wp *Win32Path) Dirname(p string) string {
	return ""
}

func (wp *Win32Path) Extname(p string) string {
	return ""
}

func (wp *Win32Path) Format(p ForamtInputPathObject) string {
	return ""
}

func (wp *Win32Path) IsAbsolute(p string) bool {
	return true
}

func (wp *Win32Path) Join(p ...string) string {
	return ""
}

func (wp *Win32Path) Normalize(p string) string { return "" }

func (wp *Win32Path) Parse(p string) ForamtInputPathObject {
	return nil
}

func (wp *Win32Path) Relative(from, to string) string {
	return ""
}

func (wp *Win32Path) Resolve(p ...string) string {
	return ""
}

func (wp *Win32Path) ToNamespacedPath(p string) string {
	return ""
}

func NewWin32Path(runtime *goja.Runtime) PlatformPath {
	wp := &Win32Path{}
	wp.runtime = runtime
	return wp
}
