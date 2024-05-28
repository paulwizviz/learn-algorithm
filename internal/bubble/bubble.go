package bubble

type NumericType interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Array[V NumericType] []V

func Sort[V NumericType](data Array[V]) Array[V] {
	datalen := len(data)
	swapped := make(Array[V], datalen)
	copy(swapped, data)
	for range datalen - 1 {
		for i := 0; i < datalen; i++ {
			if i+1 == datalen-1 {
				if swapped[i] > swapped[i+1] {
					swapped[i], swapped[i+1] = swapped[i+1], swapped[i]
				}
				break
			} else {
				if swapped[i] > swapped[i+1] {
					swapped[i], swapped[i+1] = swapped[i+1], swapped[i]
				}
			}
		}
	}
	return swapped
}
