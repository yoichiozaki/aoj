
import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	flag := false
	for scanner.Scan() {
		board := strings.Split(scanner.Text(), "")
		flag = true
		for i := 0; i <= 6; i += 3 {
			if (board[i] == "o" || board[i] == "x") && board[i] == board[i+1] && board[i+1] == board[i+2] {
				fmt.Println(board[i])
				flag = false
				break
			}
		}
		if flag {
			for i := 0; i <= 2; i++ {
				if (board[i] == "o" || board[i] == "x") && board[i] == board[i+3] && board[i] == board[i+6] {
					fmt.Println(board[i])
					flag = false
					break
				}
			}
		}
		if flag {
			if (board[0] == "o" || board[0] == "x") && board[0] == board[4] && board[4] == board[8] {
				fmt.Println(board[0])
				flag = false
			}
		}
		if flag {
			if (board[2] == "o" || board[2] == "x") && board[2] == board[4] && board[4] == board[6] {
				fmt.Println(board[2])
				flag = false
			}
		}
		if flag {
			fmt.Println("d")
		}
		board = board[:0]
	}
}
