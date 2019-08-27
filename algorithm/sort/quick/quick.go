package quick

func Run(nums []int) {
	length := len(nums)
	if length < 2 {
		return
	}

	if length == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return
	}

	p := sort(nums, nums[length >> 1])

	Run(nums[:p-1])
	Run(nums[p-1:])

	return
}

func sort(nums []int, pivotValue int) int {
	small := 0
	pivot := 0
	great := 0

	for cursor := 0; cursor < len(nums); cursor++ {
		if nums[cursor] > pivotValue {
			great++
		} else if nums[cursor] == pivotValue {
			if pivot != great {
				nums[pivot], nums[great] = nums[great], nums[pivot]
			}
			pivot++
			great++
		} else if nums[cursor] < pivotValue {
			if pivot != great {
				nums[pivot], nums[great] = nums[great], nums[pivot]
			}
			if small != pivot {
				nums[small], nums[pivot] = nums[pivot], nums[small]
			}
			small++
			pivot++
			great++
		}
	}

	return pivot
}