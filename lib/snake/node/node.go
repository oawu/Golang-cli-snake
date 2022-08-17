/**
 * @author      OA Wu <oawu.tw@gmail.com>
 * @copyright   Copyright (c) 2015 - 2021
 * @license     http://opensource.org/licenses/MIT  MIT License
 * @link        https://www.ioa.tw/
 */

package node

import (
  "snake/lib/window/row/column"
)

type Node struct {
	Column *column.Column
	PrevList *Node
	NextList *Node
}