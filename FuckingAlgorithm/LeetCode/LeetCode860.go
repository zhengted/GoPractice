package main

func lemonadeChange(bills []int) bool {
	if len(bills) <= 0 {
		return true
	}
	mHasMoney := map[int]int{
		5:  0,
		10: 0,
		20: 0,
	}
	for _, item := range bills {
		mHasMoney[item]++
		if item == 5 {
			continue
		}
		if item == 10 {
			if mHasMoney[5] <= 0 {
				return false
			}
			mHasMoney[5]--
		}
		if item == 20 {
			if mHasMoney[10] > 0 && mHasMoney[5] > 0 {
				mHasMoney[10]--
				mHasMoney[5]--
				continue
			}
			if mHasMoney[5] >= 3 {
				mHasMoney[5] = mHasMoney[5] - 3
				continue
			}
			return false
		}
	}
	return true
}
