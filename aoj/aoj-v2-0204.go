package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const (
	INF = 0x7FFFFFFF
	EPS = 1e-10
)

type Point struct {
	x, y float64
}

func (p *Point) plus(q *Point) *Point {
	return &Point{p.x + q.x, p.y + q.y}
}
func (p *Point) minus(q *Point) *Point {
	return &Point{p.x - q.x, p.y - q.y}
}
func (p *Point) multi(a float64) *Point {
	return &Point{p.x * a, p.y * a}
}
func (p *Point) div(a float64) *Point {
	return &Point{p.x / a, p.y / a}
}
func (p *Point) abs() float64 {
	return math.Sqrt(p.norm())
}
func (p *Point) norm() float64 {
	return p.x*p.x + p.y*p.y
}
func (p *Point) isLessThan(q *Point) bool {
	if p.x != q.x {
		return p.x < q.x && p.y < q.y
	}
	return true
}
func (p *Point) isEqual(q *Point) bool {
	return math.Abs(p.x-q.x) < EPS && math.Abs(p.y-q.y) < EPS
}

type Segment struct {
	p1, p2 *Point
}

func (s *Segment) norm() float64 {
	return (s.p2.x-s.p1.x)*(s.p2.x-s.p1.x) +
		(s.p2.y-s.p1.y)*(s.p2.y-s.p1.y)
}
func (s *Segment) abs() float64 {
	return math.Sqrt(s.norm())
}

func norm(a *Point) float64 {
	return a.x*a.x + a.y*a.y
}

func abs(a *Point) float64 {
	return math.Sqrt(norm(a))
}

func polar(a, r float64) *Point {
	return &Point{math.Cos(r) * a, math.Sin(r) * a}
}

func getDistance(a, b *Point) float64 {
	return abs(a.minus(b))
}

func dot(a, b *Point) float64 {
	return a.x*b.x + a.y*b.y
}

func cross(a, b *Point) float64 {
	return a.x*b.y - a.y*b.x
}

func getDistanceLP(s Segment, p *Point) float64 {
	return cross(s.p2.minus(s.p1), p.minus(s.p1)) / abs(s.p2.minus(s.p1))
}

func getDistanceSP(s Segment, p *Point) float64 {
	if dot(s.p2.minus(s.p1), p.minus(s.p1)) < 0 {
		return abs(p.minus(s.p1))
	}
	if dot(s.p1.minus(s.p2), p.minus(s.p2)) < 0 {
		return abs(p.minus(s.p2))
	}
	return getDistanceLP(s, p)
}

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		R, _ := strconv.Atoi(buf[0])
		N, _ := strconv.Atoi(buf[1])
		if R == 0 && N == 0 {
			break
		}
		p := make([]*Point, N)
		r := make([]float64, N)
		v := make([]float64, N)
		d := make([]float64, N)
		nd := make([]float64, N)
		die := make([]bool, N)
		for i := 0; i < N; i++ {
			scanner.Scan()
			buf_ := strings.Split(scanner.Text(), " ")
			_x, _ := strconv.ParseFloat(buf_[0], 64)
			_y, _ := strconv.ParseFloat(buf_[1], 64)
			_r, _ := strconv.ParseFloat(buf_[2], 64)
			_v, _ := strconv.ParseFloat(buf_[3], 64)
			tmp := &Point{_x, _y}
			p[i] = tmp
			r[i] = _r
			v[i] = _v
			nd[i] = getDistance(p[i], &Point{0, 0})
			d[i] = getDistance(p[i], &Point{0, 0})
			die[i] = false
		}
		for {
			for i := 0; i < N; i++ {
				nd[i] = math.Max(float64(R), nd[i]-v[i])
			}
			mini := float64(INF)
			id := -1
			for i := 0; i < N; i++ {
				if die[i] {
					continue
				}
				if nd[i] < EPS+float64(R) {
					continue
				}
				if nd[i] > mini {
					continue
				}
				mini = nd[i]
				id = i
			}
			if id == -1 {
				break
			}
			s := Segment{p[id], p[id]}
			s.p1 = s.p1.div(d[id])
			s.p1 = s.p1.multi(float64(R))
			s.p2 = s.p2.div(d[id])
			s.p2 = s.p2.multi(10000000.0)
			for i := 0; i < N; i++ {
				if nd[i] < EPS+float64(R) {
					continue
				}
				a := p[i]
				a = a.div(d[i])
				a = a.multi(nd[i])
				if getDistanceSP(s, a) < EPS+r[i] {
					die[i] = true
				}
			}
		}
		ans := 0
		for i := 0; i < N; i++ {
			if !die[i] {
				ans++
			}
		}
		fmt.Println(ans)
	}
}
