package policy

// 排序算法接口
type ISortPolicy interface {
	Name() string
	Sort(data []int) []int
}

// 冒泡排序
type bubbleSortPolicy struct {
}

func newBubbleSortPolicy() ISortPolicy {
	return &bubbleSortPolicy{}
}

func (b *bubbleSortPolicy) Name() string {
	return "BubbleSort"
}

func (b *bubbleSortPolicy) Sort(data []int) []int {
	if data == nil {
		return nil
	}

	size := len(data)
	if size <= 1 {
		return data
	}

	for {
		i := size - 1
		changed := false

		for {
			if i <= 0 {
				break
			}
			j := i - 1

			if data[j] > data[i] {
				data[i],data[j] = data[j],data[i]
				changed = true
			}

			i--
		}

		if !changed {
			break
		}
	}

	return data
}

// 选择排序
type selectSortPolicy struct {
}

func newSelectSortPolicy() ISortPolicy {
	return &selectSortPolicy{}
}

func (s *selectSortPolicy) Name() string {
	return "SelectSort"
}

func (s *selectSortPolicy) Sort(data []int) []int {
	if data == nil {
		return nil
	}

	size := len(data)
	if size <= 1 {
		return data
	}

	i := 0
	for {
		if i >= size - 1 {
			break
		}

		p, m := s.min(data, i + 1, size - 1)
		if m < data[i] {
			data[p], data[i] = data[i], data[p]
		}

		i++
	}

	return data
}

func (s *selectSortPolicy) min(data []int, from int, to int) (p int, v int) {
	i := from
	p = from
	v = data[from]

	if to <= from {
		return p, v
	}

	for {
		i++
		if i > to {
			break
		}

		if data[i] < v {
			p = i
			v = data[i]
		}
	}

	return p, v
}