/**
 * @author      OA Wu <oawu.tw@gmail.com>
 * @copyright   Copyright (c) 2015 - 2022
 * @license     http://opensource.org/licenses/MIT  MIT License
 * @link        https://www.ioa.tw/
 */

package gameover

import (
	"fmt"
	"math"
	"os"
	"snake/lib"
	"snake/lib/window"
)

func repeat(str string, c int) string {
	tmp := ""
	for i := 0; i < c; i++ {
		tmp = tmp + str
	}
	return tmp
}
func Show(strs ...string) {
	str := "Game Over"
	if len(strs) > 0 && len(strs[0]) > 0 {
		str = strs[0]
	}
	padding := 4
	w := lib.StrWidth(str) + padding*2

	y := uint16(math.Floor(float64(float64(window.Shared.H))/2-1) - 1)
	x := uint16(math.Floor((float64(float64(window.Shared.W))-16)/2 - 1))

	strs = []string{
		repeat(" ", w),
		fmt.Sprintf("%s%s%s", repeat(" ", padding), str, repeat(" ", padding)),
		repeat(" ", w),
	}

	for i, str := range strs {
		window.Shared.GetRowColumn(y+uint16(i), x).Set(fmt.Sprintf("\x1b[48;5;1m%s\x1b[0m", str)).Reflash()
	}
	fmt.Print("\x1b[?25h")
	os.Exit(0)
}
