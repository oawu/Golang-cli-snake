/**
 * @author      OA Wu <oawu.tw@gmail.com>
 * @copyright   Copyright (c) 2015 - 2022
 * @license     http://opensource.org/licenses/MIT  MIT License
 * @link        https://www.ioa.tw/
 */

package window

import (
	"errors"
	"fmt"
	"github.com/eiannone/keyboard"
	"golang.org/x/term"
	"snake/lib/window/row"
	"snake/lib/window/row/column"
	"sync"
)

type Window struct {
	W       uint16
	H       uint16
	RowList *row.Row

	PressUps    []*func()
	PressDowns  []*func()
	PressLefts  []*func()
	PressRights []*func()
	PressSpaces []*func()
}

var (
	Shared *Window
)

func init() {
	Shared = New()
}

func New() *Window {
	w, h, e := term.GetSize(0)

	if e != nil {
		return nil
	}

	win := &Window{W: uint16(w), H: uint16(h - 1)}

	var prevR *row.Row
	for i := uint16(0); i < win.H; i++ {
		ro := &row.Row{ColumnList: nil, PrevRowList: prevR}

		var prevC *column.Column
		for j := uint16(0); j < win.W; j++ {
			col := &column.Column{X: j, Y: i, Str: " ", Border: column.BORDER_TBLR, Row: ro, LeftList: prevC}

			if prevC == nil {
				ro.ColumnList = col
			} else {
				prevC.RightList = col
			}
			if prevR != nil {
				col.TopList = prevR.ColumnList.Get(j)
				if col.TopList != nil {
					col.TopList.BottomList = col
				}
			}

			prevC = col
		}

		if prevR == nil {
			win.RowList = ro
		} else {
			prevR.NextRowList = ro
		}

		prevR = ro
	}

	for r := win.RowList; r != nil; r = r.NextRowList {
		for c := r.ColumnList; c != nil; c = c.RightList {
			if c.TopList == nil && c.BottomList == nil && c.LeftList == nil && c.RightList == nil {
				c.Border = column.BORDER_tblr
			}
			if c.TopList == nil && c.BottomList == nil && c.LeftList == nil && c.RightList != nil {
				c.Border = column.BORDER_tblR
			}
			if c.TopList == nil && c.BottomList == nil && c.LeftList != nil && c.RightList == nil {
				c.Border = column.BORDER_tbLr
			}
			if c.TopList == nil && c.BottomList == nil && c.LeftList != nil && c.RightList != nil {
				c.Border = column.BORDER_tbLR
			}
			if c.TopList == nil && c.BottomList != nil && c.LeftList == nil && c.RightList == nil {
				c.Border = column.BORDER_tBlr
			}
			if c.TopList == nil && c.BottomList != nil && c.LeftList == nil && c.RightList != nil {
				c.Border = column.BORDER_tBlR
			}
			if c.TopList == nil && c.BottomList != nil && c.LeftList != nil && c.RightList == nil {
				c.Border = column.BORDER_tBLr
			}
			if c.TopList == nil && c.BottomList != nil && c.LeftList != nil && c.RightList != nil {
				c.Border = column.BORDER_tBLR
			}
			if c.TopList != nil && c.BottomList == nil && c.LeftList == nil && c.RightList == nil {
				c.Border = column.BORDER_Tblr
			}
			if c.TopList != nil && c.BottomList == nil && c.LeftList == nil && c.RightList != nil {
				c.Border = column.BORDER_TblR
			}
			if c.TopList != nil && c.BottomList == nil && c.LeftList != nil && c.RightList == nil {
				c.Border = column.BORDER_TbLr
			}
			if c.TopList != nil && c.BottomList == nil && c.LeftList != nil && c.RightList != nil {
				c.Border = column.BORDER_TbLR
			}
			if c.TopList != nil && c.BottomList != nil && c.LeftList == nil && c.RightList == nil {
				c.Border = column.BORDER_TBlr
			}
			if c.TopList != nil && c.BottomList != nil && c.LeftList == nil && c.RightList != nil {
				c.Border = column.BORDER_TBlR
			}
			if c.TopList != nil && c.BottomList != nil && c.LeftList != nil && c.RightList == nil {
				c.Border = column.BORDER_TBLr
			}
			if c.TopList != nil && c.BottomList != nil && c.LeftList != nil && c.RightList != nil {
				c.Border = column.BORDER_TBLR
			}
			c.Reset()
		}
	}

	return win
}

func (win *Window) Clean() *Window {
	if win == nil {
		return win
	}

	fmt.Print("\x1b[2J\x1b[0f")
	for i := uint16(0); i < win.H; i++ {
		for j := uint16(0); j < win.W; j++ {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
	fmt.Print("\x1b[?25l")

	return win
}
func (win *Window) PrintAll() *Window {
	if win == nil {
		return win
	}
	win.RowList.PrintAll()
	return win
}

func (win *Window) GetRow(y uint16) *row.Row {
	if win == nil {
		return nil
	}
	return win.RowList.Get(y)
}
func (win *Window) GetColumn(x uint16) *column.Column {
	if win == nil {
		return nil
	}
	return win.RowList.GetColumn(x)
}
func (win *Window) GetRowColumn(y uint16, x uint16) *column.Column {
	return win.GetRow(y).GetColumn(x)
}

func (win *Window) Reset() *Window {
	if win == nil {
		return win
	}

	for r := win.RowList; r != nil; r = r.NextRowList {
		r.ResetColumns()
	}

	return win
}
func (win *Window) Set(str string) *Window {
	if win == nil {
		return win
	}
	for r := win.RowList; r != nil; r = r.NextRowList {
		r.SetColumns(str)
	}
	return win
}
func (win *Window) Reflash(strs ...string) *Window {
	if win == nil {
		return win
	}
	for r := win.RowList; r != nil; r = r.NextRowList {
		r.ReflashColumns(strs...)
	}
	return win
}

func (win *Window) ResetSync() *Window {
	if win == nil {
		return win
	}

	wg := new(sync.WaitGroup)
	for r := win.RowList; r != nil; r = r.NextRowList {
		wg.Add(1)
		go r.ResetColumnsSync(wg)
	}
	wg.Wait()

	return win
}
func (win *Window) SetSync(str string) *Window {
	if win == nil {
		return win
	}
	wg := new(sync.WaitGroup)
	for r := win.RowList; r != nil; r = r.NextRowList {
		wg.Add(1)
		go r.SetColumnsSync(wg, str)
	}
	wg.Wait()
	return win
}
func (win *Window) ReflashSync(strs ...string) *Window {
	if win == nil {
		return win
	}
	wg := new(sync.WaitGroup)
	for r := win.RowList; r != nil; r = r.NextRowList {
		wg.Add(1)
		go r.ReflashColumnsSync(wg, strs...)
	}
	wg.Wait()
	return win
}

func (win *Window) PressUp(closure func()) *Window {
	if win == nil {
		return win
	}
	win.PressUps = append(win.PressUps, &closure)
	return win
}
func (win *Window) PressDown(closure func()) *Window {
	if win == nil {
		return win
	}
	win.PressDowns = append(win.PressDowns, &closure)
	return win
}
func (win *Window) PressLeft(closure func()) *Window {
	if win == nil {
		return win
	}
	win.PressLefts = append(win.PressLefts, &closure)
	return win
}
func (win *Window) PressRight(closure func()) *Window {
	if win == nil {
		return win
	}
	win.PressRights = append(win.PressRights, &closure)
	return win
}
func (win *Window) PressSpace(closure func()) *Window {
	if win == nil {
		return win
	}
	win.PressSpaces = append(win.PressSpaces, &closure)
	return win
}
func (win *Window) ListenKeyboard(closure func(*error)) *Window {
	if win == nil {
		return win
	}

	events, err := keyboard.GetKeys(10)

	if err != nil {
		closure(&err)
		return win
	}

	defer keyboard.Close()

	for {
		event := <-events

		if event.Err != nil {
			continue
		}
		if event.Key == keyboard.KeyEsc {
			err := errors.New("結束遊戲！")
			closure(&err)
			return win
		}
		if event.Key == keyboard.KeyArrowUp {
			for _, pressUp := range win.PressUps {
				go (*pressUp)()
			}
		}
		if event.Key == keyboard.KeyArrowDown {
			for _, pressDown := range win.PressDowns {
				go (*pressDown)()
			}
		}
		if event.Key == keyboard.KeyArrowLeft {
			for _, pressLeft := range win.PressLefts {
				go (*pressLeft)()
			}
		}
		if event.Key == keyboard.KeyArrowRight {
			for _, pressRight := range win.PressRights {
				go (*pressRight)()
			}
		}
		if event.Key == keyboard.KeySpace {
			for _, PressSpace := range win.PressSpaces {
				go (*PressSpace)()
			}
		}
	}
}
