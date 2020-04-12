package task5

import "github.com/ivkalita/cti/internal/chapter3/common"

func TwoStacksSort(s common.Stack) common.Stack {
	sorted := &stack{}
	for {
		greatest, err := s.Pop()
		if err == common.StackIsEmpty {
			break
		}
		greatestCnt := 1
		for {
			el, err := s.Pop()
			if err == common.StackIsEmpty {
				break
			}
			if el > greatest {
				for i := 0; i < greatestCnt; i++ {
					sorted.Push(greatest)
				}
				greatest = el
				greatestCnt = 1
			} else if el == greatest {
				greatestCnt++
			} else {
				sorted.Push(el)
			}
		}

		for {
			el, err := sorted.Pop()
			if err == common.StackIsEmpty {
				break
			}
			if el > greatest {
				sorted.Push(el)
				break
			}
			if el == greatest {
				continue
			}
			s.Push(el)
		}

		for i := 0; i < greatestCnt; i++ {
			sorted.Push(greatest)
		}
	}

	return sorted
}
