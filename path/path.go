package path

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const ModuleName = "node:path"

func IsWin32Platform() {
}

type Path struct {
	runtime *goja.Runtime
	win32   PlatformPath
	posix   PlatformPath
}

type ForamtInputPathObject struct {
	Root string
	Dir  string
	Base string
	Ext  string
	Name string
}

type PlatformPath interface {
	Basename(p string, suffix ...string) string
	Dirname(p string) string
	Extname(p string) string
	Format(p ForamtInputPathObject) string
	IsAbsolute(p string) bool
	Join(p ...string) string
	Normalize(p string) string
	Parse(p string) ForamtInputPathObject
	Relative(from, to string) string
	Resolve(p ...string) string
	ToNamespacedPath(p string) string
}

func Require(runtime *goja.Runtime, module *goja.Object) {
	p := &Path{runtime: runtime}

	exports := module.Get("exports").(*goja.Object)
	// exports.Set("Path", "")
}

func Enable(runtime *goja.Runtime) {
	runtime.Set("path", require.Require(runtime, ModuleName))
}

func init() {
	require.RegisterNativeModule(ModuleName, Require)
}
