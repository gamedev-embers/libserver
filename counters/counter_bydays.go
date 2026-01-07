package counters

import (
	"slices"
	"time"
)

const (
	DAY_DURATION = 24 * time.Hour
)

type CounterByDays struct {
	StartDate time.Time
	Value     []int32
	total     int
}

func NewCounterByDays(days int, starDate time.Time) *CounterByDays {
	return &CounterByDays{
		StartDate: zeroTime(starDate),
		Value:     make([]int32, days),
	}
}

func NewCounterByDaysWith(days int, starDate time.Time, values []int32) *CounterByDays {
	if len(values) <= 0 {
		values = make([]int32, days)
	} else if len(values) < days {
		prefix := make([]int32, days-len(values))
		values = append(prefix, values...)
	} else if len(values) > days {
		values = values[len(values)-days:]
	}
	return &CounterByDays{
		StartDate: zeroTime(starDate),
		Value:     values,
	}
}

func (c *CounterByDays) Add(value int32, now time.Time) {
	dayIndex := int(now.Sub(c.StartDate) / DAY_DURATION)
	if dayIndex < 0 {
		return
	}
	if int(dayIndex) >= len(c.Value) {
		// slice 值往前移动, 使用 copy 直接 shift
		shift := int(dayIndex) - len(c.Value) + 1
		if shift >= len(c.Value) {
			for i := 0; i < len(c.Value); i++ {
				c.Value[i] = 0
			}
			dayIndex = 0
			c.StartDate = zeroTime(now)
		} else {
			copy(c.Value, c.Value[shift:])
			copy(c.Value[len(c.Value)-shift:], make([]int32, shift))
			dayIndex -= shift
			c.StartDate = c.StartDate.Add(time.Duration(shift) * DAY_DURATION)
		}
	}
	c.Value[dayIndex] += value
	c.total = c.calcTotal()
}

func (c *CounterByDays) GetTotal() int {
	return c.total
}

func (c *CounterByDays) calcTotal() int {
	total := 0
	for _, v := range c.Value {
		total += int(v)
	}
	return total
}

func (c *CounterByDays) SetValue(values []int32, startDate time.Time) {
	c.StartDate = zeroTime(startDate)
	c.Value = values
	c.total = c.calcTotal()
}

func (c *CounterByDays) GetValue() (values []int32, startDate time.Time) {
	values = slices.Clone(c.Value)
	startDate = c.StartDate
	return
}
