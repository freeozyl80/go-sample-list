package main
import "fmt"
const (
	// process success
	STATUS_OK = iota
	//process error
	STATUS_ERROR
	//local error
	STATUS_HALT
)
func main() {
 fmt.Println(STATUS_HALT)
 fmt.Println(STATUS_ERROR)
}
