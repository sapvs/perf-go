package perf

func MapWithSwitch(x int) bool {
	switch x1 := x; {
	case x1 < 0:
		return false
	case x < 9:
		return true
	case x < 99:
		return false
	case x < 999:
		return true
	case x < 9999:
		return false
	case x < 99999:
		return true
	default:
		return false
	}
}

func MapWithIf(x int) bool {
	if x < 0 {
		return false
	}
	if x < 9 {
		return true
	}
	if x < 99 {
		return false
	}
	if x < 999 {
		return true
	}
	if x < 9999 {
		return false
	}
	if x < 99999 {
		return true
	}
	return false

}

func MapWithElseif(x int) bool {

	if x < 0 {
		return false
	} else if x < 9 {
		return true
	} else if x < 99 {
		return false
	} else if x < 999 {
		return true
	} else if x < 9999 {
		return false
	} else if x < 99999 {
		return true
	} else {
		return false
	}
}
