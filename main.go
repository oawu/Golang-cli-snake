/**
 * @author      OA Wu <oawu.tw@gmail.com>
 * @copyright   Copyright (c) 2015 - 2021
 * @license     http://opensource.org/licenses/MIT  MIT License
 * @link        https://www.ioa.tw/
 */

package main

import (
	"snake/lib/gameover"
	"snake/lib/info"
	"snake/lib/snake"
	"snake/lib/window"
)

func main() {
	// 清除螢幕，並且印出預設
	window.Shared.Clean().PrintAll()

	// 設置底下資訊
	info.Shared.Reflash()

	// 設定蛇的起始位置
	snake.Shared.In(window.Shared.GetRowColumn(1, 1))
	snake.Shared.Run()

	// 設定上下左右時
	window.Shared.PressUp(func() { // 上
		snake.Shared.Turn(snake.DIRECTION_TOP)
	})
	window.Shared.PressDown(func() { // 下
		snake.Shared.Turn(snake.DIRECTION_BOTTOM)
	})
	window.Shared.PressLeft(func() { // 左
		snake.Shared.Turn(snake.DIRECTION_LEFT)
	})
	window.Shared.PressRight(func() { // 右
		snake.Shared.Turn(snake.DIRECTION_RIGHT)
	})
	window.Shared.PressSpace(func() { // 空白鍵，暫停
		snake.Shared.Turn(snake.DIRECTION_EMPTY)
	})

	// 開始監聽鍵盤
	window.Shared.ListenKeyboard(func(err *error) { // 錯誤時
		gameover.Show((*err).Error())
	})
}
