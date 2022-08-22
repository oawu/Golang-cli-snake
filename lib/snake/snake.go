/**
 * @author      OA Wu <oawu.tw@gmail.com>
 * @copyright   Copyright (c) 2015 - 2021
 * @license     http://opensource.org/licenses/MIT  MIT License
 * @link        https://www.ioa.tw/
 */

package snake

import (
	"snake/lib/food"
	"snake/lib/gameover"
	"snake/lib/info"
	"snake/lib/snake/node"
	"snake/lib/window/row/column"
	"time"
)

type Direction int

const (
	DIRECTION_EMPTY Direction = iota
	DIRECTION_TOP
	DIRECTION_BOTTOM
	DIRECTION_LEFT
	DIRECTION_RIGHT
)

type Snake struct {
	NodeList  *node.Node
	Direction Direction
	isLock    bool
}

var Shared *Snake

func init() {
	Shared = &Snake{
		NodeList:  nil,
		Direction: DIRECTION_RIGHT,
		isLock:    false,
	}
}

func (snake *Snake) lock() *Snake {
	if snake == nil {
		return snake
	}
	snake.isLock = true
	return snake
}
func (snake *Snake) unlock() *Snake {
	if snake == nil {
		return snake
	}
	snake.isLock = false
	return snake
}
func (snake *Snake) Run() *Snake {
	if snake == nil {
		return snake
	}

	go func(snake *Snake, t time.Duration) {
		for {
			snake.Turn(snake.Direction)
			time.Sleep(t)
		}
	}(snake, time.Second/5)

	return snake
}
func (snake *Snake) Turn(directions ...Direction) *Snake {
	// 沒有節點或沒有位置不執行
	if snake == nil || snake.NodeList == nil || snake.NodeList.Column == nil {
		return snake
	}
	if snake.isLock {
		return snake
	} else {
		snake.lock()
	}
	defer snake.unlock()

	direction := snake.Direction
	if len(directions) > 0 {
		direction = directions[0]
	}

	// 目前方向與轉動方向相衝突時
	if (snake.Direction == DIRECTION_TOP && direction == DIRECTION_BOTTOM) || (snake.Direction == DIRECTION_BOTTOM && direction == DIRECTION_TOP) || (snake.Direction == DIRECTION_LEFT && direction == DIRECTION_RIGHT) || (snake.Direction == DIRECTION_RIGHT && direction == DIRECTION_LEFT) {
		return snake
	} else {
		snake.Direction = direction
	}

	// 取得下一步
	c := snake.NodeList.Column
	switch snake.Direction {
	case DIRECTION_TOP:
		c = snake.NodeList.Column.TopList
	case DIRECTION_BOTTOM:
		c = snake.NodeList.Column.BottomList
	case DIRECTION_LEFT:
		c = snake.NodeList.Column.LeftList
	case DIRECTION_RIGHT:
		c = snake.NodeList.Column.RightList
	case DIRECTION_EMPTY:
		c = snake.NodeList.Column
	}

	// 不動
	if c == snake.NodeList.Column {
		return snake
	}

	// 下一步不存在時
	if c == nil {
		gameover.Show()
		return snake
	}

	// 下一步是食物時
	if food.Shared != nil && food.Shared.Column != nil && food.Shared.Column == c {
		snake.Eat(food.Shared.Column)
	}

	// 找最後的節點
	last := snake.NodeList
	for ; last.NextList != nil; last = last.NextList {
	}

	// 最後節點個位置清除
	if last.Column != nil {
		last.Column.Reset().Reflash()
	}

	// 講位置更新為前一個節點位置
	for ; last.PrevList != nil; last = last.PrevList {
		last.Column = last.PrevList.Column
	}
	// 第一節點接上新的位置
	last.Column = c

	return snake.Reflash()
}
func (snake *Snake) Eat(c *column.Column) *Snake {
	// 位置不存在時
	if snake == nil || c == nil {
		return snake
	}

	node := &node.Node{Column: c}

	// 食物位置改為開頭
	node.NextList = snake.NodeList
	snake.NodeList = node
	if node.NextList != nil {
		node.NextList.PrevList = node
	}

	// 取得蛇的身體
	columns := []*column.Column{}
	node = snake.NodeList
	for ; node != nil; node = node.NextList {
		if node.Column != nil {
			columns = append(columns, node.Column)
		}
	}

	// 重新產生食物，食物產生的位置要避免在蛇的身體上
	food.Shared.Random(columns)

	// 增加分數
	info.Shared.AddScore()

	return snake.Reflash()
}
func (snake *Snake) Reflash() *Snake {
	if snake == nil {
		return snake
	}

	for n := snake.NodeList; n != nil; n = n.NextList {
		if n.Column == nil {
			continue
		}
		if n == snake.NodeList {
			n.Column.Reflash("\x1b[38;5;6m█\x1b[0m")
		} else {
			n.Column.Reflash("\x1b[2m\x1b[38;5;6m▓\x1b[0m\x1b[0m")
		}
	}

	// 判斷是否碰壁
	if snake.NodeList != nil && snake.NodeList.Column != nil && snake.NodeList.Column.Border != column.BORDER_TBLR {
		gameover.Show()
	}

	return snake
}
func (snake *Snake) In(c *column.Column) *Snake {
	return snake.Eat(c).Reflash()
}
