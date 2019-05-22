// run the main function
package main

// import modules, you can ("module1","module2") but the below format is typically better
import (
	"fmt"
	"math"

	"github.com/arpbadger/go_crash_course2/03_packages/strutil"
)

// create and run the main function
func main() {
	// print to the screen
	fmt.Println(strutil.Reverse("olleh"))
	// run some modules from the math package
	fmt.Println(math.Floor(2.7))
	//	fmt.Println(math.Ceil(2.7))
	//	fmt.Println(math.Sqrt(4))
}
