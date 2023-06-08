package path

import (
	"runtime"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/errors"
	"github.com/dop251/goja_nodejs/require"
)

const ModuleName = "node:path"

func IsWin32Platform() bool {
	return runtime.GOOS == "windows"
}

type Path struct {
	runtime   *goja.Runtime
	Win32     PlatformPath
	Posix     PlatformPath
	Sep       string
	Delimiter string
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

func CreatePathCalled(name string, fn PlatformPath) func(goja.FunctionCall) goja.Value {
	return func(call goja.FunctionCall) goja.Value {
		arguments := call.Arguments
		switch name {
		case "basename":
			// fn()
			if len(arguments) == 0 {
				panic(errors.NewTypeError(r *goja.Runtime, code string, params ...interface{}))
			}
		}
	}
}

func Require(runtime *goja.Runtime, module *goja.Object) {
	win32Path := NewWin32Path(runtime)
	proto := runtime.NewObject()
	if IsWin32Platform() {
		proto.Set("basename", win32Path.Basename)
		proto.Set("posix", nil)
	}
}

func Enable(runtime *goja.Runtime) {
	runtime.Set("path", require.Require(runtime, ModuleName))
}

func init() {
	require.RegisterNativeModule(ModuleName, Require)
}
