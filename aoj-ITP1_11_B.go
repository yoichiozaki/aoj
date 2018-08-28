
import "fmt"

/*
展開図の展開の仕方を固定して、回転動作によってどのように展開図の数字が移るかを記録する。
展開図は次の通り。手前の目2番を中心に据える展開の仕方。
         ___
    ____|_1_|_______
    |_4_|_2_|_3_|_5_|
        |_6_|

サイコロの状態は上面と前面を指定することで一意に決定される。それぞれを指定された目に揃えて、
展開図の3番を出力すれば良い。
*/
type Dice struct {
	Nums []int
}

func (d *Dice) rotate(direction string) {
	switch direction {
	case "W":
		d.Nums[0], d.Nums[2], d.Nums[5], d.Nums[3] = d.Nums[2], d.Nums[5], d.Nums[3], d.Nums[0]
	case "E":
		d.Nums[0], d.Nums[2], d.Nums[5], d.Nums[3] = d.Nums[3], d.Nums[0], d.Nums[2], d.Nums[5]
	case "N":
		d.Nums[0], d.Nums[4], d.Nums[5], d.Nums[1] = d.Nums[1], d.Nums[0], d.Nums[4], d.Nums[5]
	case "S":
		d.Nums[5], d.Nums[1], d.Nums[0], d.Nums[4] = d.Nums[1], d.Nums[0], d.Nums[4], d.Nums[5]
	}
}

func (d *Dice) roll(direction string) {
	switch direction {
	case "R":
		d.Nums[3], d.Nums[1], d.Nums[2], d.Nums[4] = d.Nums[1], d.Nums[2], d.Nums[4], d.Nums[3]
	case "L":
		d.Nums[3], d.Nums[1], d.Nums[2], d.Nums[4] = d.Nums[4], d.Nums[3], d.Nums[1], d.Nums[2]
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

func main() {
	dice := Dice{
		make([]int, 6),
	}
	for i := 0; i < 6; i++ {
		var n int
		fmt.Scan(&n)
		dice.Nums[i] = n
	}
	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		var up, front int
		fmt.Scan(&up, &front)

		// 上面を揃える
		upIndex := getIndex(dice.Nums, up)
		switch upIndex {
		case 0: // もう揃ってる
			break
		case 1: // N方向に一回回転させて揃う
			dice.rotate("N")
		case 2: // W方向に一回回転させて揃う
			dice.rotate("W")
		case 3: // E方向に一回回転させて揃う
			dice.rotate("E")
		case 4: // S方向に一回回転させて揃う
			dice.rotate("S")
		default: // 初期状態の上面の対面を上面に指定する→適当な同一方向に二回転
			dice.rotate("N")
			dice.rotate("N")
		}
		// dice.printStatus()

		// 前面を揃える
		frontIndex := getIndex(dice.Nums, front)
		switch frontIndex {
		case 1: // もう揃ってる
			break
		case 2: // 右方向に一回回転させて揃う
			dice.roll("R")
		case 3: // 左方向に一回回転させて揃う
			dice.roll("L")
		case 4: // 上下面指定された後の状態の前面の対面を前面に指定する→適当な同一方向に二回転
			dice.roll("R")
			dice.roll("R")
		default: // あり得ない組み合わせの指定。たとえばup == 1 && front == 1。
			fmt.Println("Impossible order.")
		}
		// dice.printStatus()

		fmt.Println(dice.Nums[2])
	}
}

