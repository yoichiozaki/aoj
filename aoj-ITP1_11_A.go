
import "fmt"

/*
展開図の展開の仕方を固定して、回転動作によってどのように展開図の数字が移るかを記録する。
展開図は次の通り。手前の２番を中心に据える展開の仕方。
         ___
    ____|_1_|_______
    |_4_|_2_|_3_|_5_|
        |_6_|

*/
type Dice struct {
	Nums []int
}

func (d *Dice) rotete(direction string) {
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

func main() {
	dice := Dice{
		make([]int, 6),
	}
	for i := 0; i < 6; i++ {
		var n int
		fmt.Scan(&n)
		dice.Nums[i] = n
	}
	var rotationOrder string
	fmt.Scan(&rotationOrder)
	for _, direction := range []rune(rotationOrder) {
		dice.rotete(string(direction))

		/*
		fmt.Printf("rotate direction: %s\n", string(direction))
		dice.printStatus()
		*/
	}
	fmt.Println(dice.Nums[0])
}

