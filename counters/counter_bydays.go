package counters

import "time"

const (
	DAY_DURATION = 24 * time.Hour
)

type CounterByDays struct {
	startDate time.Time
	value     []int
	total     int
}

func NewCounterByDays(days int, starDate time.Time) *CounterByDays {
	return &CounterByDays{
		startDate: zeroTime(starDate),
		value:     make([]int, days),
	}
}

func (c *CounterByDays) Add(value int, now time.Time) {
	dayIndex := int(now.Sub(c.startDate) / DAY_DURATION)
	if dayIndex < 0 {
		return
	}
	if int(dayIndex) >= len(c.value) {
		// slice 值往前移动, 使用 copy 直接 shift
		shift := int(dayIndex) - len(c.value) + 1
		if shift >= len(c.value) {
			for i := 0; i < len(c.value); i++ {
				c.value[i] = 0
			}
			dayIndex = 0
			c.startDate = zeroTime(now)
		} else {
			copy(c.value, c.value[shift:])
			copy(c.value[len(c.value)-shift:], make([]int, shift))
			dayIndex -= shift
			c.startDate = c.startDate.Add(time.Duration(shift) * DAY_DURATION)
		}
	}
	c.value[dayIndex] += value
	c.total = c.calcTotal()
}

func (c *CounterByDays) GetTotal() int {
	return c.total
}

func (c *CounterByDays) calcTotal() int {
	total := 0
	for _, v := range c.value {
		total += v
	}
	return total
}

func (c *CounterByDays) Reset(values []int, startDate time.Time) {
	c.startDate = zeroTime(startDate)
	c.value = values
	c.total = c.calcTotal()
}
