// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build k210

package bus

const (
	Core Bus = iota
	TileLink
	AXI
	AHB
	APB0
	APB1
	APB2

	AHBLast = AHB
	APBLast = APB2
)

type Bus uint8

var str = [7]string{
	"Core",
	"TileLink",
	"AXI",
	"AHB",
	"APB0",
	"APB1",
	"APB2",
}

func (b Bus) String() string { return str[b] }

var buses [7]struct{ clockHz int64 }

func (b Bus) Clock() int64      { return buses[b].clockHz }
func (b Bus) SetClock(Hz int64) { buses[b].clockHz = Hz }