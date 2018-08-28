
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

const END = 2100000000

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	scanner.Split(bufio.ScanWords)
	q := nextInt()
	var new_ele *element
	var front *element
	front = gene_element(END)
	var p *element = front
	for i := 0; i < q; i++ {
		switch nextInt() {
		case 0:
			x := nextInt()
			new_ele = gene_element(x)
			if front.val == END {
				front = new_ele
				new_ele.next = p
				p.prev = new_ele
				p = new_ele
			} else {
				if front == p {
					front = new_ele
					new_ele.next = p
					p.prev = new_ele
					p = new_ele
				} else {
					p.prev.next = new_ele
					new_ele.prev = p.prev
					new_ele.next = p
					p.prev = new_ele
					p = new_ele
				}
			}
		case 1:
			d := nextInt()
			if d > 0 {
				for j := 0; j < d; j++ {
					p = p.next
				}
			} else {
				for j := 0; j < -d; j++ {
					p = p.prev
				}
			}
		case 2:
			if front == p {
				p.next.prev = nil
				front = front.next
				p = front
			} else {
				p.prev.next = p.next
				p.next.prev = p.prev
				p = p.next
			}
		default:
			panic("no such operation!")
		}
	}
	p = front
	for p != nil {
		if p.val != END {
			fmt.Println(p.val)
		}
		p = p.next
	}
}

type element struct {
	val int
	next *element
	prev *element
}

func gene_element(val int) *element {
	return &element{val:val, next:nil, prev:nil}
}
