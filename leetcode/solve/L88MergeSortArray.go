package main

func merge1(nums1 []int, m int, nums2 []int, n int) {
	mi, ni := m-1, n-1
	for i := m + n - 1; mi >= 0 || ni >= 0; i-- {
		if mi < 0 {
			nums1[i] = nums2[ni]
			ni--
		} else if ni < 0 {
			nums1[i] = nums1[mi]
			mi--
		} else {
			if nums1[mi] >= nums2[ni] {
				nums1[i] = nums1[mi]
				mi--
			} else {
				nums1[i] = nums2[ni]
				ni--
			}
		}
	}
}
