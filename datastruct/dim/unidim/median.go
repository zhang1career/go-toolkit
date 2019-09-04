package unidim

func FindMedianFromSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	start := (m + n - 1) >> 1
	stop := (m + n) >> 1 + 1
	i := 0
	j := 0
	// ready
	for ; i+j < start; {
		if (i < m) && (j < n) {
			if nums1[i] <= nums2[j] {
				i++
			} else {
				j++
			}
			continue
		}
		if i < m {
			i++
		}
		if j < n {
			j++
		}
	}
	// go
	var mid = make([]int, 0)
	for ; i+j < stop; {
		if (i < m) && (j < n) {
			if nums1[i] <= nums2[j] {
				mid = append(mid, nums1[i])
				i++
			} else {
				mid = append(mid, nums2[j])
				j++
			}
			continue
		}
		if i < m {
			mid = append(mid, nums1[i])
			i++
		}
		if j < n {
			mid = append(mid, nums2[j])
			j++
		}
	}
	
	var ret float64
	for _, v := range mid {
		ret += float64(v)
	}
	return ret / float64(len(mid))
}
