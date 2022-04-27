import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenghtCurrent1 := len(nums1)
	lenghtCurrent2 := len(nums2)
	median1 := 0.0
	median2 := 0.0

	if lenghtCurrent1 == 0 {
		median2 = median(&nums2)
		return median2
	} else if lenghtCurrent2 == 0 {
		median1 = median(&nums1)
		return median1
	}
	median1 = median(&nums1)
	median2 = median(&nums2)

	for {
		if median1 == median2 {
			return median1
		}
		if lenghtCurrent1 == 1 && lenghtCurrent2 == 1 {
			twoNum := []int{nums1[0], nums2[0]}
			return median(&twoNum)
		}
		if lenghtCurrent1+lenghtCurrent2 <= 4 {
			fourNum := []int{}
			fourNum = append(fourNum, nums1[:]...)
			fourNum = append(fourNum, nums2[:]...)
			sort.Ints(fourNum[:])
			return median(&fourNum)
		}
		if lenghtCurrent1 <= 2 || lenghtCurrent2 <= 2 {
			// longer array is atleast 3 length
			fiveNum := []int{}
			if lenghtCurrent1 > lenghtCurrent2 {
				mid1 := mid(lenghtCurrent1)
				fiveNum = append(fiveNum, nums2[:]...)
				if lenghtCurrent1%2 == 0 {
					fiveNum = append(fiveNum, nums1[mid1-1:mid1+3]...)
				} else {
					fiveNum = append(fiveNum, nums1[mid1-1:mid1+2]...)
				}
			} else {
				mid2 := mid(lenghtCurrent2)
				fiveNum = append(fiveNum, nums1[:]...)
				if lenghtCurrent2%2 == 0 {
					fiveNum = append(fiveNum, nums2[mid2-1:mid2+3]...)
				} else {
					fiveNum = append(fiveNum, nums2[mid2-1:mid2+2]...)
				}
			}
			sort.Ints(fiveNum[:])
			return median(&fiveNum)
		}

		short := []int{}
		long := []int{}
		if lenghtCurrent1 <= lenghtCurrent2 {
			short = nums1
			long = nums2
		} else {
			short = nums2
			long = nums1
		}
		lenghtShort := len(short)
		lenghtLong := len(long)
		medianShort := median(&short)
		medianLong := median(&long)

		if medianShort < medianLong {
			midShort := mid(lenghtShort)
			nums1 = short[midShort:lenghtShort]
			nums2 = long[0 : lenghtLong-midShort]

		} else {
			midShort := mid(lenghtShort)
			nums1 = short[0 : lenghtShort-midShort]
			nums2 = long[midShort:lenghtLong]
		}
		lenghtCurrent1 = len(nums1)
		lenghtCurrent2 = len(nums2)
		median1 = median(&nums1)
		median2 = median(&nums2)
		fmt.Println(nums1, nums2)
	}
	return 0.0
}

func mid(length int) int {
	if length == 2 {
		return 1
	}
	if length > 1 {
		return (length - 1) / 2
	}
	return length
}

func median(nums *[]int) float64 {
	length := len(*nums)
	lenHalf := length / 2
	median := 0.0
	if length%2 == 0 {
		median = float64(((*nums)[lenHalf-1] + (*nums)[lenHalf])) / 2
	} else {
		median = float64((*nums)[lenHalf])
	}
	return median
}