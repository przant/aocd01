package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "unicode"
)

const (
    inputFile = "input.txt"
)

func main() {
    ptrFile, err := os.Open(inputFile)
    if err != nil {
        log.Fatalf("trying to open the file %q: %s", inputFile, err)
    }

    defer ptrFile.Close()

    sum := uint64(0)
    scnr := bufio.NewScanner(ptrFile)
    for scnr.Scan() {
        line := []byte(scnr.Text())
        sum += calibrationValue(line)
    }

    fmt.Printf("The solution for this day is: %d\n", sum)
}

func calibrationValue(line []byte) uint64 {
    bsCalVal := []byte{firstDigit(line), lastDigit(line)}
    uiCalVal, err := strconv.ParseUint(string(bsCalVal), 10, 64)
    if err != nil {
        log.Printf("couldn't convert text calibration value %q into a number: %s", string(bsCalVal), err)
        return 0
    }
    return uiCalVal
}

func firstDigit(line []byte) (fd byte) {
    for i := 0; i < len(line); i++ {
        if unicode.IsDigit(rune(line[i])) {
            fd = line[i]
            break
        }
    }
    return fd
}

func lastDigit(line []byte) (ld byte) {
    for i := len(line) - 1; i >= 0; i-- {
        if unicode.IsDigit(rune(line[i])) {
            ld = line[i]
            break
        }
    }
    return ld
}
