# rltao

This is a go package of data from R.L. Wing's Tao of Power and a list of UTF-8 tetragrams

Tai Xuan Jing tetragrams
https://en.wiktionary.org/wiki/Taixuanjing

There are 81 permutations of the following three lines, given 4 lines/grams
------ Heaven, --  -- Earth, - - - Man
3^4 = 81

Source https://terebess.hu/english/tao/Wing.html


EXAMPLE USAGE

Create a file named test.go

    package main

    import (
        "fmt"

        "github.com/billzajac/rltao"
    )

    func main() {
        id := rltao.GenerateId()
        tetragram := rltao.GetTetragram(id)
        var passage rltao.Passage = rltao.GetPassage(id)
        fmt.Printf("%02d\n%s\n%s\n%s\n", id, tetragram, passage.Title, passage.Body)
    }

Then get the deps with

    go mod init
    go mod tidy

Then run the test with

    go run test.go