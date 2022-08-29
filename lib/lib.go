/**
 * @author      OA Wu <oawu.tw@gmail.com>
 * @copyright   Copyright (c) 2015 - 2022
 * @license     http://opensource.org/licenses/MIT  MIT License
 * @link        https://www.ioa.tw/
 */

package lib

import (
	"fmt"
	"sync"
)

func StrWidth(str string) int {
	w := 0
	for _, c := range []rune(str) {
		s := fmt.Sprintf("%c", c)
		if len(s) == 1 {
			w += 1
		} else {
			w += 2
		}
	}
	return w
}
func GoGroups(funcs ...*func()) {
	wg := new(sync.WaitGroup)
	wg.Add(len(funcs))

	for _, tmp := range funcs {
		go func(wg *sync.WaitGroup, closure *func()) {
			defer wg.Done()
			tmp := *closure
			tmp()
		}(wg, tmp)
	}
	wg.Wait()
}
func GoGroup(funcs ...func()) {
	wg := new(sync.WaitGroup)
	wg.Add(len(funcs))
	for _, tmp := range funcs {
		go func(wg *sync.WaitGroup, closure func()) {
			defer wg.Done()
			closure()
		}(wg, tmp)
	}
	wg.Wait()
}
func GoFunc(wg *sync.WaitGroup, callback func()) {
	defer wg.Done()
	callback()
}
