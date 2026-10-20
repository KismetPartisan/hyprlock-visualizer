package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

const FILE = "/tmp/cava_file"
const LAYERS = 13

func main() {
	line := readCavaFile(FILE) // Read last line from cava_file
	
	var bars strings.Builder
	bars.WriteString("<span line_height=\"0.6\">")

	for layer := LAYERS; layer>0; layer-- {
		bars.WriteString(" ")
		for value := range strings.SplitSeq(line, ";") {
			n, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			height := n-8*(layer-1)
			height = min(height,8)
			height = max(height,0)

			switch height {
				case 0: 
					if layer == 1 {
						bars.WriteString("▁")
					} else {
						bars.WriteString("　")
					}
				case 1: bars.WriteString("▁")
				case 2: bars.WriteString("▂")
				case 3: bars.WriteString("▃")
				case 4: bars.WriteString("▄")
				case 5: bars.WriteString("▅")
				case 6: bars.WriteString("▆")
				case 7: bars.WriteString("▇")
				case 8: bars.WriteString("█")
			}
			bars.WriteString(" ")
		}
		bars.WriteString("\n")
	}
	bars.WriteString("</span>")

	result := bars.String()
	fmt.Println(result)
}

func readCavaFile(filepath string) string { // Open file and read last line into a string
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	var output = scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return strings.TrimRight(output, ";\r")
}
