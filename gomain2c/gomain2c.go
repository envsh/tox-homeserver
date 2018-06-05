package gomain2c

/*
 */
import "C"

// usage:
/*
import "gomain2c"
func init(){ gomain2c.Set(main) }
*/

func Set(f func()) { _gomain = f }

var _gomain func()

//export go_main_wrapper
func go_main_wrapper() { _gomain() }
