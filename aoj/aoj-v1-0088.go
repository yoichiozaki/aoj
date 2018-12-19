
import (
	"bufio"
	"os"
	"strings"
		"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		encoded := encoder(scanner.Text())
		decoded := decoder(encoded)
		fmt.Println(decoded)
	}
}

func encoder(target string) string {
	input := strings.Split(target, "")
	tmp := make([]string, len(input))
	for i, v := range input {
		switch v {
		case " ":
			tmp[i] = "101"
		case "'":
			tmp[i] = "000000"
		case ",":
			tmp[i] = "000011"
		case "-":
			tmp[i] = "10010001"
		case ".":
			tmp[i] = "010001"
		case "?":
			tmp[i] = "000001"
		case "A":
			tmp[i] = "100101"
		case "B":
			tmp[i] = "10011010"
		case "C":
			tmp[i] = "0101"
		case "D":
			tmp[i] = "0001"
		case "E" :
			tmp[i] = "110"
		case "F":
			tmp[i] = "01001"
		case "G":
			tmp[i] = "10011011"
		case "H":
			tmp[i] = "010000"
		case "I":
			tmp[i] = "0111"
		case "J":
			tmp[i] = "10011000"
		case "K":
			tmp[i] = "0110"
		case "L":
			tmp[i] = "00100"
		case "M":
			tmp[i] = "10011001"
		case "N":
			tmp[i] = "10011110"
		case "O":
			tmp[i] = "00101"
		case "P":
			tmp[i] = "111"
		case "Q":
			tmp[i] = "10011111"
		case "R":
			tmp[i] = "1000"
		case "S":
			tmp[i] = "00110"
		case "T":
			tmp[i] = "00111"
		case "U":
			tmp[i] = "10011100"
		case "V":
			tmp[i] = "10011101"
		case "W":
			tmp[i] = "000010"
		case "X":
			tmp[i] = "10010010"
		case "Y":
			tmp[i] = "10010011"
		case "Z":
			tmp[i] = "10010000"
		}
	}
	return strings.Join(tmp, "")
}

func decoder(target string) string {
	const splitlen = 5
	tmp := make([]string, len(target)/5+1)
	for i := 0; i < len(target); i += splitlen {
		if i+splitlen < len(target) {
			tmp = append(tmp, target[i:(i+splitlen)])
		} else {
			tmp = append(tmp, target[i:])
		}
	}
	if more := 5 - len(tmp[len(tmp)-1]); more != 0 {
		tmp[len(tmp)-1] = tmp[len(tmp)-1] + strings.Repeat("0", more)
	}
	for i, v := range tmp {
		switch v {
		case "00000":
			tmp[i] = "A"
		case "00001":
			tmp[i] = "B"
		case "00010":
			tmp[i] = "C"
		case "00011":
			tmp[i] = "D"
		case "00100":
			tmp[i] = "E"
		case "00101":
			tmp[i] = "F"
		case "00110":
			tmp[i] = "G"
		case "00111":
			tmp[i] = "H"
		case "01000":
			tmp[i] = "I"
		case "01001":
			tmp[i] = "J"
		case "01010":
			tmp[i] = "K"
		case "01011":
			tmp[i] = "L"
		case "01100":
			tmp[i] = "M"
		case "01101":
			tmp[i] = "N"
		case "01110":
			tmp[i] = "O"
		case "01111":
			tmp[i] = "P"
		case "10000":
			tmp[i] = "Q"
		case "10001":
			tmp[i] = "R"
		case "10010":
			tmp[i] = "S"
		case "10011":
			tmp[i] = "T"
		case "10100":
			tmp[i] = "U"
		case "10101":
			tmp[i] = "V"
		case "10110":
			tmp[i] = "W"
		case "10111":
			tmp[i] = "X"
		case "11000":
			tmp[i] = "Y"
		case "11001":
			tmp[i] = "Z"
		case "11010":
			tmp[i] = " "
		case "11011":
			tmp[i] = "."
		case "11100":
			tmp[i] = ","
		case "11101":
			tmp[i] = "-"
		case "11110":
			tmp[i] = "'"
		case "11111":
			tmp[i] = "?"
			
		}
	}
	return strings.Join(tmp, "")
}
