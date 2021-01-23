package main

import (
    "fmt"
    dataframe "github.com/rocketlaunchr/dataframe-go"
)

func main() {
    s1 := dataframe.NewSeriesInt64("day", nil, 1, 2, 3)
    s2 := dataframe.NewSeriesFloat64("sales", nil, 31.2, 21.2, 34.2)
    df := dataframe.NewDataFrame(s1, s2)

    iterator := df.ValuesIterator(dataframe.ValuesOptions{0, 1, true})

    df.Lock()
    for {
        row, vals, _ := iterator()
        if row == nil {
            break
        }
        fmt.Println(vals[1])
    }
    df.Unlock()
}
