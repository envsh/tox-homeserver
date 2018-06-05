package gomain2c

/*
 */
import "C"
import (
	"log"
	"sync"
)

// usage:
/*
go scope:
import "gomain2c"
func init(){ gomain2c.Set(main) }
func init(){gomain2c.Setx(somefunc, 42)}

c scope:
extern void go_main_wrapper();
go_main_wrapper();
extern void go_func_wrapper(int);
go_func_wrapper(42);

*/

func Set(f func()) { _gomain = f }

var _gomain func()

//export go_main_wrapper
func go_main_wrapper() { _gomain() }

/////
var _gofuncs = make(map[int]func())
var _gofuncsMu sync.RWMutex

func Setx(f func(), x int) {
	_gofuncsMu.Lock()
	defer _gofuncsMu.Unlock()
	if oldfn, ok := _gofuncs[x]; ok {
		log.Printf("Warning, override exist: %v => %v\n", oldfn, f)
	}
	_gofuncs[x] = f
}

//export go_func_wrapper
func go_func_wrapper(x C.int) {
	_gofuncsMu.RLock()
	f, ok := _gofuncs[int(x)]
	_gofuncsMu.RUnlock()
	if ok && f != nil {
		f()
	}
}
