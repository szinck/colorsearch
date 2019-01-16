/*
Copyright (c) 2019 Shaun Zinck

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main


import (
"fmt"
"math/rand"
)

var height = 20
var width = 20

var numColors = 16
var minLen = 4
var maxLen = 7
var numAnswers = 10


// See this excellent blog post where I got this list of colors:
// https://sashat.me/2017/01/11/list-of-20-simple-distinct-colors/
var colors = []string { "#e6194b", "#3cb44b", "#ffe119", "#4363d8", "#f58231", "#911eb4", "#42d4f4", "#f032e6", "#bfef45", "#469990", "#9A6324", "#800000", "#808000", "#000075", "#fabebe", "#a9a9a9" }

func genAnswer(grid [][]int) []int {
    // get a suitable answer
    size := minLen + rand.Intn(maxLen - minLen + 1)
    answer := make([]int,size)

    for {
        dirX := 1 - rand.Intn(3)
        dirY := 1 - rand.Intn(3)

        if (dirX == 0 && dirY == 0) {
            continue
        }

        x := rand.Intn(len(grid))
        y := rand.Intn(len(grid[0]))


        endX := x + (dirX * size)
        endY := y + (dirY * size)


        if (endX < 0 || endX >= len(grid)) {
            continue
        }
        if (endY < 0 || endY >= len(grid[0])) {
            continue
        }

        // found a viable answer
        for i := 0 ; i < size; i++ {
            answer[i] = grid[x][y]
            x += dirX
            y += dirY
        }
        break
    }

    return answer

}



func printPuzzle() {

    var rows = make([][]int, height)

    for i, _ := range rows {
        cols := make([]int, width)
        for j, _ := range cols {
            cols[j] = rand.Intn(numColors)
        }
        rows[i] = cols
    }


    // random row
    // random col
    // random direction 
    // random len 4-7

    answers := make([][]int, numAnswers)
    for i, _ := range answers {
        answers[i] = genAnswer(rows)
    }





fmt.Println(`
<table> 
  <tr>
    <td>
<h1>Color Search</h1>
      <table>
`)



    for _, r := range rows {
       fmt.Println("<tr>")
        for _, c := range r {
            fmt.Printf("<td style='background:%s;width:18px; height: 10px'>&nbsp; </td> ", colors[c] )
        }
        fmt.Println()
       fmt.Println("</tr>")
    }

    fmt.Println(`
      </table>
    </td>
  </tr>
  <tr>
    <td style='padding-top:3em'>
    <div style='text-align:center; width: 100%'>Find these patterns, forward, backward, or diagonal</div>
   `)


    fmt.Println(`
    <div style='align: center; text-align:center; width:100%'>
      <table><tr><td>`)
    for i, answer := range answers {
        if (i == len(answers) / 2) {
            fmt.Println("</td><td style='padding-left:4em'>")
        }

        fmt.Println("<table><tr>")
        for _, c := range answer {
            fmt.Printf("<td style='background:%s;width:18px;height:10px'>&nbsp; </td>", colors[c])
        }
        fmt.Println("</tr></table>")
    }
    fmt.Println("</td></tr></table></div>")

    fmt.Println(`
    </td>
  </tr>
 </table>`)


}

func main() {


    fmt.Println(` <html> 
<head>
<style>
@media print{@page {size: landscape}}

.h1 {
    text-align: center
    width: 800px
    align: center
}

.pagebreak { page-break-before: always; }

</style>
</head>

<body style='-webkit-print-color-adjust: exact;'>
`)

    for i := 0 ; i < 20 ;  i++ {
        printPuzzle();
        fmt.Println(`<div class="pagebreak"> </div>`)
    }
    fmt.Println("</body></html>")
}
