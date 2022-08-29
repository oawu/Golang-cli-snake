/**
 * @author      OA Wu <oawu.tw@gmail.com>
 * @copyright   Copyright (c) 2015 - 2022
 * @license     http://opensource.org/licenses/MIT  MIT License
 * @link        https://www.ioa.tw/
 */

package row

import (
	"snake/lib/window/row/column"
	"sync"
)

type Row struct {
	ColumnList  *column.Column
	PrevRowList *Row
	NextRowList *Row
}

func (row *Row) GetRowInterface() {}

func (row *Row) PrintAll() *Row {
	if row == nil {
		return row
	}
	row.ColumnList.PrintAll()
	row.NextRowList.PrintAll()
	return row
}

func (row *Row) Get(index uint16) *Row {
	for i := 0; index != uint16(i) && row != nil; i++ {
		row = row.NextRowList
	}
	return row
}
func (row *Row) GetColumn(index uint16) *column.Column {
	if row == nil {
		return nil
	}
	return row.ColumnList.Get(index)
}

func (row *Row) ResetColumns() *Row {
	if row == nil {
		return row
	}
	for c := row.ColumnList; c != nil; c = c.RightList {
		c.Reset()
	}

	return row
}
func (row *Row) SetColumns(str string) *Row {
	if row == nil {
		return row
	}
	for c := row.ColumnList; c != nil; c = c.RightList {
		c.Set(str)
	}

	return row
}
func (row *Row) ReflashColumns(strs ...string) *Row {
	if row == nil {
		return row
	}

	for c := row.ColumnList; c != nil; c = c.RightList {
		c.Reflash(strs...)
	}
	return row
}

func (row *Row) ResetColumnsSync(wg1 *sync.WaitGroup) *Row {
	defer func(wg *sync.WaitGroup) {
		if wg != nil {
			wg1.Done()
		}
	}(wg1)

	if row == nil {
		return row
	}

	wg2 := new(sync.WaitGroup)
	for c := row.ColumnList; c != nil; c = c.RightList {
		wg2.Add(1)
		go c.ResetSync(wg2)
	}
	wg2.Wait()

	return row
}
func (row *Row) SetColumnsSync(wg1 *sync.WaitGroup, str string) *Row {
	defer func(wg *sync.WaitGroup) {
		if wg != nil {
			wg1.Done()
		}
	}(wg1)

	if row == nil {
		return row
	}
	wg2 := new(sync.WaitGroup)
	for c := row.ColumnList; c != nil; c = c.RightList {
		wg2.Add(1)
		go c.SetSync(wg2, str)
	}
	wg2.Wait()

	return row
}
func (row *Row) ReflashColumnsSync(wg1 *sync.WaitGroup, strs ...string) *Row {
	defer func(wg *sync.WaitGroup) {
		if wg != nil {
			wg1.Done()
		}
	}(wg1)

	if row == nil {
		return row
	}

	wg2 := new(sync.WaitGroup)
	for c := row.ColumnList; c != nil; c = c.RightList {
		wg2.Add(1)
		go c.ReflashSync(wg2, strs...)
	}
	wg2.Wait()
	return row
}
