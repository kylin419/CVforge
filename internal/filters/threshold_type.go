package filters

import "fmt"

func ParseThresholdType(s string) (ThresholdType, error) {
	switch s {
	case "binary":
		return Binary, nil

	case "binary-inv":
		return BinaryInv, nil

	case "truncate":
		return Truncate, nil

	case "tozero":
		return ToZero, nil

	case "tozero-inv":
		return ToZeroInv, nil

	default:
		return 0, fmt.Errorf("unknown threshold type: %s", s)
	}
}
