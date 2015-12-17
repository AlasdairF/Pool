package pool

const (
	Max = 1024 // maximum memory overhead 76 MB
	Size = 76490 // determined in trials on writing to disk and writing to memory
)

var pool = make(chan []byte, Max)

func Get(l int) []byte {
	if l > Size {
		return make([]byte, l)
	}
	var c []byte
    select {
		case c = <- pool:
		default: c = make([]byte, Size)
    }
	return c[0:l]
}

func Return(c []byte) {
	if cap(c) == Size {
		select {
			case pool <- c[0:Size]:
			default:
		}
	}
}
