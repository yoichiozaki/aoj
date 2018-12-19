
import "fmt"

/*
展開図の展開の仕方を固定して、回転動作によってどのように展開図の数字が移るかを記録する。
展開図は次の通り。手前の目2番を中心に据える展開の仕方。
         ___
    ____|_1_|_______
    |_4_|_2_|_3_|_5_|
        |_6_|

サイコロの状態は上面と前面を指定することで一意に決定される。それぞれを指定された目に揃えて、
展開図の3番を記録。
入力された二つのサイコロについて上面前面を揃えた上で右側面の値を比較することで判定。
で、できるかと思ったが、サイコロの各面の数字がすべて異なるという制約は今回の問題にはないので、
この判定方法は不完全。

よって、力ずくで全通り調べる。
*/

type Dice struct {
	Nums []int
}

func (d *Dice) rotate(direction string, times int) {
	times %= 4
	switch direction {
	case "W":
		for i := 0; i < times; i++ {
			d.Nums[0], d.Nums[2], d.Nums[5], d.Nums[3] = d.Nums[2], d.Nums[5], d.Nums[3], d.Nums[0]
		}
	case "E":
		for i := 0; i < times; i++ {
			d.Nums[0], d.Nums[2], d.Nums[5], d.Nums[3] = d.Nums[3], d.Nums[0], d.Nums[2], d.Nums[5]
		}
	case "N":
		for i := 0; i < times; i++ {
			d.Nums[0], d.Nums[4], d.Nums[5], d.Nums[1] = d.Nums[1], d.Nums[0], d.Nums[4], d.Nums[5]
		}
	case "S":
		for i := 0; i < times; i++ {
			d.Nums[5], d.Nums[1], d.Nums[0], d.Nums[4] = d.Nums[1], d.Nums[0], d.Nums[4], d.Nums[5]
		}
	}
}

func (d *Dice) roll(direction string, times int) {
	times %= 4
	switch direction {
	case "R":
		for i := 0; i < times; i++ {
			d.Nums[3], d.Nums[1], d.Nums[2], d.Nums[4] = d.Nums[1], d.Nums[2], d.Nums[4], d.Nums[3]
		}
	case "L":
		for i := 0; i < times; i++ {
			d.Nums[3], d.Nums[1], d.Nums[2], d.Nums[4] = d.Nums[4], d.Nums[3], d.Nums[1], d.Nums[2]
		}
	}
}

func (d *Dice) printStatus() {
	status :=
`	　　　 _ _ _ _
     _ _ _|_%3d_|_ _ _ _ _ _
    |_%3d_|_%3d_|_%3d_|_%3d_|
 　       |_%3d_|
`
	fmt.Printf(status, d.Nums[0], d.Nums[3], d.Nums[1], d.Nums[2], d.Nums[4], d.Nums[5])
	fmt.Println()
}

func getIndex(nums []int, target int) (index int) {
	for i, v := range nums {
		if target == v {
			index = i
			break
		} else {
			index = -1
		}
	}
	return index
}

func (d *Dice) isEqual(other Dice) bool {
	for i, v := range d.Nums {
		if v != other.Nums[i] {
			return false
		}
	}
	return true
}

func main() {
	dice1 := Dice{
		make([]int, 6),
	}
	dice2 := Dice{
		make([]int, 6),
	}
	for i := 0; i < 6; i++ {
		var n int
		fmt.Scan(&n)
		dice1.Nums[i] = n
	}
	for i := 0; i < 6; i++ {
		var n int
		fmt.Scan(&n)
		dice2.Nums[i] = n
	}

	directionOrder := []string{"N", "S", "W", "E"}
	rollOrder := []string{"R", "L"}
	flag := false

	CHECK:
		for i := 0; i <= 2; i++ {
			for j := 0; j <= 2; j++ {
				for _, direction := range directionOrder {
					for _, roll := range rollOrder {
						tmp := new(Dice)
						tmp.Nums = append([]int{}, dice2.Nums...)
						tmp.rotate(direction, i)
						tmp.roll(roll, j)
						if tmp.isEqual(dice1) {
							flag = true
							break CHECK
						}
					}
				}
			}
		}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

