// Pair of elements.
package pair

import "fmt"

type Pair struct {
	First  interface{}
	Second interface{}
}

func (p *Pair) String() string {
	return fmt.Sprintf("('%v', '%v')", p.First, p.Second)
}
