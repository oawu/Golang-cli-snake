/**
 * @author      OA Wu <oawu.tw@gmail.com>
 * @copyright   Copyright (c) 2015 - 2022
 * @license     http://opensource.org/licenses/MIT  MIT License
 * @link        https://www.ioa.tw/
 */

package column

import (
	"fmt"
	"sync"
)

type Border int

const (
	BORDER_TBLR Border = iota
	BORDER_TBLr
	BORDER_TBlR
	BORDER_TBlr
	BORDER_TbLR
	BORDER_TbLr
	BORDER_TblR
	BORDER_Tblr
	BORDER_tBLR
	BORDER_tBLr
	BORDER_tBlR
	BORDER_tBlr
	BORDER_tbLR
	BORDER_tbLr
	BORDER_tblR
	BORDER_tblr
)

type Column struct {
	X   uint16
	Y   uint16
	Str string

	Row    row
	Border Border

	LeftList   *Column
	RightList  *Column
	TopList    *Column
	BottomList *Column
}

type row interface {
	GetRowInterface()
}

func getD4Str(border Border) string {
	switch border {
	case BORDER_tblr:
		return ""
	case BORDER_tblR:
		return ""
	case BORDER_tbLr:
		return ""
	case BORDER_tbLR:
		return ""
	case BORDER_tBlr:
		return ""
	case BORDER_Tblr:
		return ""
	case BORDER_TBlr:
		return ""

	case BORDER_tBlR:
		return "╭"
	case BORDER_tBLR:
		return "─"
	case BORDER_tBLr:
		return "╮"

	case BORDER_TBlR:
		return "│"
	case BORDER_TBLR:
		return " "
	case BORDER_TBLr:
		return "│"

	case BORDER_TblR:
		return "╰"
	case BORDER_TbLR:
		return "─"
	case BORDER_TbLr:
		return "╯"
	}
	return ""
}

func (column *Column) print() *Column {
	if column == nil {
		return column
	}
	fmt.Printf("\x1b[s\x1b[%d;%dH%s\x1b[u", column.Y+1, column.X+1, column.Str)
	return column
}
func (column *Column) PrintAll() *Column {
	if column == nil {
		return column
	}

	column.print().RightList.PrintAll()
	return column
}

func (column *Column) Get(index uint16) *Column {
	for i := 0; index != uint16(i) && column != nil; i++ {
		column = column.RightList
	}
	return column
}

func (column *Column) Reset() *Column {
	if column == nil {
		return column
	}
	str := getD4Str(column.Border)
	if str == "" {
		return column
	}
	column.Str = str
	return column
}
func (column *Column) Set(str string) *Column {
	if column == nil {
		return column
	}
	column.Str = str
	return column
}
func (column *Column) Reflash(strs ...string) *Column {
	if column == nil {
		return column
	}

	if len(strs) > 0 {
		column.Str = strs[0]
	}
	return column.print()
}

func (column *Column) ResetRows() *Column {
	for ; column != nil; column = column.BottomList {
		column.Reset()
	}

	return column
}
func (column *Column) SetRows(str string) *Column {
	for ; column != nil; column = column.BottomList {
		column.Set(str)
	}

	return column
}
func (column *Column) ReflashRows(strs ...string) *Column {
	for ; column != nil; column = column.BottomList {
		column.Reflash(strs...)
	}

	return column
}

func (column *Column) ResetSync(wg *sync.WaitGroup) *Column {
	defer func(wg *sync.WaitGroup) {
		if wg != nil {
			wg.Done()
		}
	}(wg)

	if column == nil {
		return column
	}
	str := getD4Str(column.Border)
	if str == "" {
		return column
	}
	column.Str = str
	return column
}
func (column *Column) SetSync(wg *sync.WaitGroup, str string) *Column {
	defer func(wg *sync.WaitGroup) {
		if wg != nil {
			wg.Done()
		}
	}(wg)

	if column == nil {
		return column
	}
	column.Str = str
	return column
}
func (column *Column) ReflashSync(wg *sync.WaitGroup, strs ...string) *Column {
	defer func(wg *sync.WaitGroup) {
		if wg != nil {
			wg.Done()
		}
	}(wg)

	if column == nil {
		return column
	}

	if len(strs) > 0 {
		column.Str = strs[0]
	}
	return column.print()
}

func (column *Column) ResetRowsSync() *Column {
	wg := new(sync.WaitGroup)
	for ; column != nil; column = column.BottomList {
		wg.Add(1)
		go column.ResetSync(wg)
	}
	wg.Wait()

	return column
}
func (column *Column) SetRowsSync(str string) *Column {
	wg := new(sync.WaitGroup)
	for ; column != nil; column = column.BottomList {
		wg.Add(1)
		go column.SetSync(wg, str)
	}
	wg.Wait()

	return column
}
func (column *Column) ReflashRowsSync(strs ...string) *Column {
	wg := new(sync.WaitGroup)
	for ; column != nil; column = column.BottomList {
		wg.Add(1)
		column.ReflashSync(wg, strs...)
	}
	wg.Wait()

	return column
}
