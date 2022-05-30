package stats

import "math"

func SMA(data []*float64, window int) *float64 {
	sum := 0.0
	elements := 0
	buffer := make([]float64, window)

	in := 0
	for i := 0; i < len(data); i++ {
		if data[i] == nil {
			continue
		}

		sum += *data[i]
		sum -= buffer[in]
		buffer[in] = *data[i]
		elements++

		in++
		if in >= window {
			in = 0
		}
	}

	if elements == 0 {
		return nil
	}

	sma := sum / math.Min(float64(elements), float64(window))

	return &sma
}
