
package numlib

import (
    "log"
    "os"
    "strconv"
    "strings"
    //n1d "github.com/ikuo0/gotest1/num1d"
    n2d "github.com/ikuo0/gotest1/num2d"
)

func SaveTxt(fileName string, m n2d.Mat) {
    if f, err := os.Create(fileName); err != nil {
        log.Fatal(err)
    } else {
        defer func() {
            f.Close()
        }()
        //file.Write(([]byte)(output))
        rowSize, colSize := n2d.Size(m)
        for r := 0; r < rowSize; r++ {
            sarr := []string{}
            for c := 0; c < colSize; c++ {
                sarr = append(sarr, strconv.FormatFloat(m[r][c], 'e', -1, 64))
            }
            f.Write([]byte(strings.Join(sarr, "\t") + "\n"))
        }
    }
}
