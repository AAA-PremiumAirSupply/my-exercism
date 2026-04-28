package raindrops

import "strconv"

func Convert(number int) string {
	var mask uint8
	if number%3 == 0 {
		mask |= 0b001
	}
	if number%5 == 0 {
		mask |= 0b010
	}
	if number%7 == 0 {
		mask |= 0b100
	}

	switch mask {
	case 0b000:
		return strconv.Itoa(number)
	case 0b001:
		return "Pling"
	case 0b010:
		return "Plang"
	case 0b011:
		return "PlingPlang"
	case 0b100:
		return "Plong"
	case 0b101:
		return "PlingPlong"
	case 0b110:
		return "PlangPlong"
	case 0b111:
		return "PlingPlangPlong"
	default:
		return strconv.Itoa(number)
	}
}
