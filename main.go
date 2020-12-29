package main

import(
  "fmt"
  "io/ioutil"
  "strings"
  "regexp"
)

func main() {
    fmt.Println("Starting Advent of Code Day 5")

    // Manual unit test (automatic in repl.it not easy to run)
    //testCalculateRowNr()
    //return

    input := readInput()
    boardingPasses := strings.Split(input, "\n")
    seatIdHighest := 0

    for index, boardingPass := range boardingPasses  {
      fmt.Println("idx: ", index+1, " boardingPass: ", boardingPass)

      // check if valid format
      if(len(boardingPass) != 10) {
        fmt.Println("invalid: wrong length")
        continue;
      }

      if(!regexp.MustCompile("[F|B]{7}[L|R]{3}").MatchString(boardingPass)) {
        fmt.Println("invalid: wrong content")
        continue;
      }

      // extract row and column part
      boardingPassRow := boardingPass[0:7]
      boardingPassCol := boardingPass[7:10]

      // calculate row and col number
      rowNr := calculateRowNr(boardingPassRow)
      colNr := calculateColNr(boardingPassCol)

      fmt.Println("rowNr:", rowNr, " colNr: ", colNr)

      // calculate seat id and memory highest
      seatId := rowNr * 8 + colNr
      fmt.Println("seatId: ", seatId)

      if(seatId > seatIdHighest) {
        seatIdHighest = seatId
      }
    }

    fmt.Println("seatIdHighest: ", seatIdHighest)
}

func readInput() string {
    content, err := ioutil.ReadFile("puzzleInput.txt")

    if err != nil {
        return ""
    }

    // Convert []byte to string and print to screen
    text := string(content)
    return text
}

func calculateRowNr(boardingPassRow string) int{
  rowNr := 0
  rowRangeMin := 0
  rowRangeMax := 127
  lastCharIndex := len(boardingPassRow)-1

  for index, rowChar := range boardingPassRow  {
    seatCountBefore := rowRangeMax - rowRangeMin
    seatCountBefore++
    
    if(string(rowChar) == "B") {
      // B means to take the upper half
      if(index == lastCharIndex) {
        rowNr = rowRangeMax
      }

      rowRangeMin = rowRangeMin + seatCountBefore / 2
    } else {
      // F means take the lower half
      if(index == lastCharIndex) {
        rowNr = rowRangeMin
      }

      rowRangeMax = seatCountBefore / 2 + rowRangeMin
      rowRangeMax--
    }

    fmt.Println("calculateRowNr: index=", index, " rowRangeMin=", rowRangeMin, " rowRangeMax=", rowRangeMax, " seatCountBefore=", seatCountBefore)
  }

  return rowNr
}

func calculateColNr(boardingPassCol string) int{
  colNr := 0
  colRangeMin := 0
  colRangeMax := 7
  lastCharIndex := len(boardingPassCol)-1

  for index, colChar := range boardingPassCol  {
    if(string(colChar) == "R") {
      // R means to take the upper half
      if(index == lastCharIndex) {
        colNr = colRangeMax
      }

      colRangeMin = colRangeMin + ( colRangeMax - colRangeMin ) / 2 + 1 // add one
    } else {
      // L means take the lower half
      if(index == lastCharIndex) {
        colNr = colRangeMin
      }

      colRangeMax = ( colRangeMax - colRangeMin ) / 2 + colRangeMin - 1 // minus one
    }

    fmt.Println("calculateColNr: index=", index, " colRangeMin=", colRangeMin, " colRangeMax=", colRangeMax)
  }

  return colNr
}

func testCalculateRowNr() {
	testDataTable := []struct {
		boardingPassCol string
		correctRowNr int
	}{
		{"FBFBBFF", 44},
		{"BFFFBBF", 70},
		{"FFFBBBF", 14},
		{"BBFFBBF", 102},
	}

	for _, testData := range testDataTable {
		actualRowNr := calculateRowNr(testData.boardingPassCol)
		if actualRowNr == testData.correctRowNr {
      fmt.Println("Correct RowNr of ", testData.boardingPassCol, " got: ", actualRowNr, " , want: " , testData.correctRowNr)
    } else {
			fmt.Println("!Fail RowNr of ", testData.boardingPassCol, " was incorrect, got: ", actualRowNr, " , want: " , testData.correctRowNr)
		}
	}
}