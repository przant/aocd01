package main

import (
    "bufio"
    "bytes"
    "fmt"
    "log"
    "os"
    "strconv"
    "unicode"
)

const (
    inputFile = "input.txt"
)

var (
    digits = map[string]byte{
        "one":   byte(49),
        "two":   byte(50),
        "three": byte(51),
        "four":  byte(52),
        "five":  byte(53),
        "six":   byte(54),
        "seven": byte(55),
        "eight": byte(56),
        "nine":  byte(57),
    }
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
        value, err := calibrationValue(line)
        if err != nil {
            log.Printf("getting the calibration value from %q: %s", string(line), err)
        }
        sum += value
    }

    fmt.Printf("The solution for this day is: %d\n", sum)
}

func calibrationValue(line []byte) (uint64, error) {
    bsCalVal := []byte{firstDigit(line), lastDigit(line)}
    uiCalVal, err := strconv.ParseUint(string(bsCalVal), 10, 64)
    if err != nil {
        log.Printf("couldn't convert text calibration value %q into a number: %s", string(bsCalVal), err)
        return 0, err
    }
    return uiCalVal, nil
}

func firstDigit(line []byte) (fd byte) {
    numIdx := 0
    numByteVal := byte(0)
    strIdx := len(line)
    strByteVal := byte(0)

    for i := 0; i < len(line); i++ {
        if unicode.IsDigit(rune(line[i])) {
            numByteVal, numIdx = line[i], i
            break
        }
    }

    for key, value := range digits {
        bsKey := []byte(key)
        if bytes.Contains(line[0:numIdx], bsKey) {
            idx := bytes.Index(line[0:numIdx], bsKey)
            if idx != -1 && idx < strIdx {
                strIdx = idx
                strByteVal = value
            }
        }
    }

    if strByteVal != 0 && strIdx < numIdx {
        fd = strByteVal
    } else {
        fd = numByteVal
    }

    return fd
}

func lastDigit(line []byte) (ld byte) {
    numIdx := 0
    numByteVal := byte(0)
    strIdx := 0
    strByteVal := byte(0)

    for i := len(line) - 1; i >= 0; i-- {
        if unicode.IsDigit(rune(line[i])) {
            numByteVal, numIdx = line[i], i
            break
        }
    }

    for key, value := range digits {
        bsKey := []byte(key)
        if bytes.Contains(line[numIdx+1:], bsKey) {
            idx := bytes.LastIndex(line, bsKey)
            if idx != -1 && strIdx < idx {
                strByteVal = value
                strIdx = idx
            }
        }
    }

    if strByteVal != 0 && strIdx > numIdx {
        ld = strByteVal
    } else {
        ld = numByteVal
    }

    return ld
}
