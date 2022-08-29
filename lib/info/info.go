/**
 * @author      OA Wu <oawu.tw@gmail.com>
 * @copyright   Copyright (c) 2015 - 2022
 * @license     http://opensource.org/licenses/MIT  MIT License
 * @link        https://www.ioa.tw/
 */

package info

import (
	"fmt"
	"snake/lib"
	"snake/lib/window"
)

type Info struct {
	title string
	score uint32
}

var (
	Shared *Info
)

func init() {
	Shared = &Info{title: " 操作：↑ ↓ ← →   離開：esc   暫停：space"}
}

func (info *Info) AddScore() *Info {
	if info == nil {
		return info
	}

	info.score += 1

	return info.Reflash()
}
func (info *Info) getScore() string {
	if info == nil {
		return ""
	}

	str := ""
	if info.score > 0 {
		str = fmt.Sprintf("目前分數：%3d", info.score-1)
	} else {
		str = fmt.Sprintf("目前分數：%3d", 0)
	}
	return str
}
func (info *Info) Reflash() *Info {
	if info == nil {
		return info
	}

	s := info.getScore()
	w := lib.StrWidth(s)

	fmt.Printf("\x1b[s\x1b[%d;%dH%s\x1b[u", window.Shared.H+1, 2, info.title)
	fmt.Printf("\x1b[s\x1b[%d;%dH%s\x1b[u", window.Shared.H+1, window.Shared.W-uint16(w)-1, s)

	return info
}
