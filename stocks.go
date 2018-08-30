package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "sort"
    "bufio"
)

type company_investment struct {
    Company string
    Percentage int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func atoi_wrapper(str string) int {
    invest, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    return invest
}

func read_stocks() []company_investment {
    file, err := os.Open("stocks.txt")
    check(err)

    defer file.Close()

    stocks := []company_investment{}

    scanner := bufio.NewScanner(file)
    scanner.Scan() // skip first line

    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " ")
        stocks = append(stocks,
            company_investment{
                Company: line[0],
                Percentage: atoi_wrapper(line[1]),
            })
    }

    return stocks
}

func sort_stocks(stocks []company_investment) []company_investment {
    sort.Slice(stocks, func(i, j int) bool {
        return stocks[i].Percentage > stocks[j].Percentage
    })

    return stocks
}

func print_stock_invest(invest int) {
    stocks := read_stocks()
    sorted_stocks := sort_stocks(stocks)
    for _, company_investment := range sorted_stocks {
        fmt.Println(
            strings.Title(company_investment.Company), "-", invest*company_investment.Percentage/100)
    }
}

func main() {
    invest_s := os.Args[1]
    invest := atoi_wrapper(invest_s)
    print_stock_invest(invest)
}
