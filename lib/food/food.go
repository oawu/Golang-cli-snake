/**
 * @author      OA Wu <oawu.tw@gmail.com>
 * @copyright   Copyright (c) 2015 - 2022
 * @license     http://opensource.org/licenses/MIT  MIT License
 * @link        https://www.ioa.tw/
 */

package food

import (
	xterm "github.com/oawu/Golang-cli-xterm"
	"math/rand"
	"snake/lib/window"
	"snake/lib/window/row/column"
	"time"
)

type Food struct {
	Column *column.Column
}

var Shared *Food

func init() {
	Shared = &Food{}
}

func include(excludes []*column.Column, c *column.Column) bool {
	for _, exclude := range excludes {
		if exclude == c {
			return true
		}
	}
	return false
}

func (food *Food) RowColumn(c *column.Column) *Food {
	if food == nil {
		return food
	}

	food.Column = c
	return food.Reflash()
}
func (food *Food) Random(excludess ...[]*column.Column) *Food {
	if food == nil {
		return food
	}

	excludes := []*column.Column{}
	if len(excludess) > 0 {
		excludes = excludess[0]
	}

	boxs := []*column.Column{}

	for r := window.Shared.RowList; r != nil; r = r.NextRowList {
		for c := r.ColumnList; c != nil; c = c.RightList {
			if !include(excludes, c) && c.Border == column.BORDER_TBLR {
				boxs = append(boxs, c)
			}
		}
	}

	if len(boxs) <= 0 {
		return food
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(boxs), func(i, j int) { boxs[i], boxs[j] = boxs[j], boxs[i] })
	food.Column = boxs[0]

	return food.Reflash()
}
func (food *Food) Reflash() *Food {
	if food == nil {
		return food
	}

	food.Column.Set(xterm.Yellow("※").String()).Reflash()

	return food
}
