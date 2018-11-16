package movingaverage
import "math"

// @author Robin Verlangen
// Moving average implementation for Go

type MovingAverage struct {
	Window int
	values []float64
	valPos int
	slotsFilled bool
}

func (ma *MovingAverage) Avg() float64 {
	var sum = float64(0)
	var c = ma.Window-1

	// Are all slots filled? If not, ignore unused
	if !ma.slotsFilled {
		c = ma.valPos-1
		if c < 0 {
			// Empty register
			return 0
		}
	}

	// Sum values
	var ic = 0
	for i := 0; i <= c; i++ {
		sum += ma.values[i]
		ic++
	}

	// Finalize average and return
	avg := sum / float64(ic)
	return avg
}

func (ma *MovingAverage) Max() float64 {
	var max = float64(0)
	var c = ma.Window-1

	// Are all slots filled? If not, ignore unused
	if !ma.slotsFilled {
		c = ma.valPos-1
		if c < 0 {
			// Empty register
			return 0
		}
	}

	// Calculate max value
	for i := 0; i <= c; i++ {
		if ma.values[i] > max || i == 0 && ma.values[0] != 0{
			max = ma.values[i]
		}
	}

	return max
}

func (ma *MovingAverage) Min() float64 {
	var min = float64(0)
	var c = ma.Window-1

	// Are all slots filled? If not, ignore unused
	if !ma.slotsFilled {
		c = ma.valPos-1
		if c < 0 {
			// Empty register
			return 0
		}
	}

	// Calculate min value
	for i := 0; i <= c; i++ {
		if ma.values[i] < min  || i == 0 && ma.values[0] != 0{
			min = ma.values[i]
		}
	}

	return min
}

func (ma *MovingAverage) Values() []float64{
	// return all values
	return ma.values
}

func (ma *MovingAverage) SlotsFilled() bool{
	// return all values
	return ma.slotsFilled
}

func (ma *MovingAverage) Add(val float64) {
	if math.IsNaN(val) {
		panic("Value to add is NaN.")
	}

	// Put into values array
	ma.values[ma.valPos] = val

	// Increment value position
	ma.valPos = (ma.valPos + 1) % ma.Window

	// Did we just go back to 0, effectively meaning we filled all registers?
	if !ma.slotsFilled && ma.valPos == 0 {
		ma.slotsFilled = true
	}
}

func New(window int) *MovingAverage {
	return &MovingAverage{
		Window : window,
		values : make([]float64, window),
		valPos : 0,
		slotsFilled : false,
	}
}
