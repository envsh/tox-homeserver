package avhlp

/*
 */
import "C"
import "os/exec"

type suvcapd struct {
	suvcap  *exec.Cmd
	rawctxp *C.char
	cretp   C.int
	cstopp  C.int
}
