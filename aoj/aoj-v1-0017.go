
import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func cipher(text string, direction int, int_shift int) string {
	shift, offset := rune(int_shift), rune(26)
	runes := []rune(text)
	for index, char := range runes {
		switch direction {
		case 1:
			if char >= 'a' && char <= 'z'-shift || char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' || char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		case -1:
			if char >= 'a'+shift && char <= 'z' || char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift || char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		}
		runes[index] = char
	}
	return string(runes)
}

func main() {
	for scanner.Scan() {
		infos := strings.Split(scanner.Text(), " ")
		plain := make([]string, len(infos))
		// fmt.Println(infos)
		Decryption:
			for i := 0; i < 26; i++ {
				for j := range infos {
					plain[j] = cipher(infos[j], 1, i)
				}
				for j := range plain {
					if plain[j] == "this" || plain[j] == "that" || plain[j] == "the" {
						break Decryption
					}
				}
			}
			fmt.Println(strings.Join(plain, " "))
	}
}

