package main

import (
    "github.com/kniren/gota/dataframe"
    "github.com/go-gota/gota/series"
    "fmt"
    "os"
)

func main() {
    df := dataframe.New(
        series.New([]string{"b", "a"}, series.String, "COL.1"),
        series.New([]int{1, 2}, series.Int, "COL.2"),
        series.New([]float64{3.0, 4.0}, series.Float, "COL.3"),
    )

    fil := df.Filter(
        dataframe.F{"COL.1", series.Eq, "b"},
    )

    fmt.Println("{}", fil)

    f, err := os.Open("./sample.csv")
    if err != nil {
        fmt.Println("error")
    }

    defer f.Close()

    df2 := dataframe.ReadCSV(f)

    fmt.Println("{}", df2)
}
