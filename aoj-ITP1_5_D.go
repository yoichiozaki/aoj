
import "fmt"

func call(n int) {
	var i int
	for i = 1 ;i <= n ;i++ {
		x := i
		if x % 3 == 0 {
			fmt.Printf(" %d", i)
		} else {
			for {
				if x%10 == 3 {
					fmt.Printf(" %d", i)
					break
				}
				x /= 10
				if x == 0 {
					break
				}
			}
		}
		if (1 + i) > n {
			break
		}
	}
	fmt.Print("\n")

}
func main() {
	var n int
	fmt.Scan(&n)
	call(n)
}

// void call(int n){
//  int i = 1;
// CHECK_NUM:
//  int x = i;
//  if ( x % 3 == 0 ){
//   cout << " " << i;
//   goto END_CHECK_NUM;
//  }
// INCLUDE3:
//  if ( x % 10 == 3 ){
//   cout << " " << i;
//   goto END_CHECK_NUM;
//  }
//  x /= 10;
//  if ( x ) goto INCLUDE3;
// END_CHECK_NUM:
//  if ( ++i <= n ) goto CHECK_NUM;
//
//  cout << endl;
// }
