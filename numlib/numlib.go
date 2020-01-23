
package numlib

import (
    //"fmt"
    "io/ioutil"
    "log"
    "os"
    "regexp"
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

func LoadTxt(fileName string) (bool, n2d.Mat){
    if f, err := os.Open(fileName); err != nil {
        log.Fatal(err)
        return false, nil
    } else {
        defer func() {
            f.Close()
        }()
        if b, err := ioutil.ReadAll(f); err != nil {
            log.Fatal(err)
            return false, nil
        } else {
            s := string(b)
            repChomp := regexp.MustCompile("[\r\n]+$")
            s = repChomp.ReplaceAllString(s, "")
            rep := regexp.MustCompile("[\r\n]+")
            lines := rep.Split(s, -1)
            rSize := len(lines)
            if rSize < 1 {
                log.Fatal("row length 0")
                return false, nil
            } else {
                firstCol := strings.Split(lines[0], "\t")
                cSize := len(firstCol)
                m := n2d.Create(rSize, cSize)
                for ri, l := range(lines) {
                    for ci, c := range(strings.Split(l, "\t")) {
                        if f, err := strconv.ParseFloat(c, 64); err != nil {
                            log.Fatal(err)
                            return false, nil
                        } else {
                            m[ri][ci] = f
                        }
                    }
                }
                return true, m
            }
        }
    }
}