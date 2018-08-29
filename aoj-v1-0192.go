// UNFINISHED

package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type Car struct {
	ID, Remaining int
}

type Part struct {
	PID, Status, Remaining int
	Upper, Lower           *Car
}

func (p *Part) progress(time int) {
	if p.Upper != nil {
		p.Upper.Remaining -= time
	}
	if p.Lower != nil {
		p.Lower.Remaining -= time
	}
	if p.Status != 0 {
		p.Remaining -= time
	}
}
func (p *Part) goOut() []int {
	switch p.Status {
	case 2:
		if p.Lower.Remaining <= 0 && p.Upper.Remaining <= 0 {
			out := []int{p.Lower.ID, p.Upper.ID}
			p.Lower = nil
			p.Upper = nil
			p.Status = 0
			p.Remaining = -1
			return out
		}
		if p.Lower.Remaining <= 0 {
			out := []int{p.Lower.ID}
			p.Lower = nil
			p.Status = 1
			p.Remaining = p.Upper.Remaining
			return out
		}
	case 1:
		if p.Upper.Remaining <= 0 {
			out := []int{p.Upper.ID}
			p.Upper = nil
			p.Status = 0
			p.Remaining = -1
			return out
		}
	}
	return []int{}
}
func (p *Part) comeIn(id, remaining int) {
	switch p.Status {
	case 0:
		p.Upper = &Car{id, remaining}
		p.Status = 1
		p.Remaining = remaining
	case 1:
		p.Lower = &Car{id, remaining}
		p.Status = 2
		p.Remaining = remaining
	}
}

type ParkingLot struct {
	Size, Capacity, CurrentSpace int
	Body                         []Part
}

func (pl *ParkingLot) progress(time int) {
	for _, part := range pl.Body {
		part.progress(time)
	}
}
func (pl *ParkingLot) goOut() []int {
	outs := make([][]int, 0)
	for _, part := range pl.Body {
		if part.Status >= 1 && part.Remaining <= 0 {
			outs = append(outs, part.goOut())
		}
	}
	ret := make([]int, 0)
	for _, out := range outs {
		ret = append(ret, out...)
	}
	pl.CurrentSpace += len(ret)
	return ret
}
func (pl *ParkingLot) comeIn(id, remaining int) {
	pl.CurrentSpace--
	for _, part := range pl.Body {
		if part.Status == 0 {
			part.comeIn(id, remaining)
			return
		}
	}
	remaining_list := make([]Pair1, 0)
	for _, part := range pl.Body {
		if part.Status == 1 {
			remaining_list = append(remaining_list, Pair1{part.PID, part.Remaining})
		}
	}
	sort.Sort(ByRemaining(remaining_list))
	r, i := 0, 0
	for _, p := range remaining_list {
		r, i = p.Remaining, p.ID
		if r >= remaining {
			pl.Body[i].comeIn(id, remaining)
			return
		}
	}
	max_r := r
	for _, p := range remaining_list {
		if p.Remaining == max_r {
			pl.Body[p.ID].comeIn(id, remaining)
			return
		}
	}
}

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		m, _ := strconv.Atoi(buf[0])
		n, _ := strconv.Atoi(buf[1])
		if m == 0 && n == 0 {
			break
		}
		tmp := make([]Part, m)
		for i := range tmp {
			tmp[i] = Part{i, 0, -1, nil, nil}
		}
		parkinglot := ParkingLot{m, 2 * m, 2 * m, tmp}
		que := make([]Pair2, 0)
		ans := make([]int, 0)
		log.Printf("parkinglot: %+v\n", parkinglot)
		for clock := 0; clock < n*120-1; clock++ {
			parkinglot.progress(1)
			ans = append(ans, parkinglot.goOut()...)
			if clock <= (n-1)*10 && clock%10 == 0 {
				scanner.Scan()
				r, _ := strconv.Atoi(scanner.Text())
				que = append(que, Pair2{r, clock/10 + 1})
			}
			for i := 0; i < min(parkinglot.CurrentSpace, len(que)); i++ {
				x := que[0]
				que = que[1:]
				rem, ind := x.x, x.y
				parkinglot.comeIn(rem, ind)
			}
		}
		log.Println(ans)
	}
}

type Pair1 struct {
	ID, Remaining int
}
type ByRemaining []Pair1

func (a ByRemaining) Len() int           { return len(a) }
func (a ByRemaining) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRemaining) Less(i, j int) bool { return a[i].Remaining < a[j].Remaining }

type Pair2 struct {
	x, y int
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
