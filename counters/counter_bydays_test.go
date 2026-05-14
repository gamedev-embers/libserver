package counters

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCounterByDays(t *testing.T) {

	startDate := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	t.Run("basic", func(t *testing.T) {
		obj := NewCounterByDays(5, startDate)
		ts := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)

		// 添加数据到第1天
		obj.Add(10, ts)
		assert.Equal(t, 10, obj.GetTotal())
		assert.Equal(t, []int32{10, 0, 0, 0, 0}, obj.Value)

		// 添加数据到第2天
		ts = ts.Add(24 * time.Hour)
		obj.Add(20, ts)
		assert.Equal(t, 30, obj.GetTotal())
		assert.Equal(t, []int32{10, 20, 0, 0, 0}, obj.Value)

		// 添加数据到第3天
		ts = ts.Add(24 * time.Hour)
		obj.Add(30, ts)
		assert.Equal(t, 60, obj.GetTotal())
		assert.Equal(t, []int32{10, 20, 30, 0, 0}, obj.Value)

		// 添加数据到第4天
		ts = ts.Add(24 * time.Hour)
		obj.Add(40, ts)
		assert.Equal(t, 100, obj.GetTotal())
		assert.Equal(t, []int32{10, 20, 30, 40, 0}, obj.Value)

		// 添加数据到第5天
		ts = ts.Add(24 * time.Hour)
		obj.Add(50, ts)
		assert.Equal(t, 150, obj.GetTotal())
		assert.Equal(t, []int32{10, 20, 30, 40, 50}, obj.Value)

		// 添加数据到第6天，触发滑动窗口
		ts = ts.Add(24 * time.Hour)
		obj.Add(60, ts)
		assert.Equal(t, 200, obj.GetTotal())
		assert.Equal(t, []int32{20, 30, 40, 50, 60}, obj.Value)

		// 添加数据到第7天
		ts = ts.Add(24 * time.Hour)
		obj.Add(70, ts)
		assert.Equal(t, 250, obj.GetTotal()) // 第2天的数据被移除
		assert.Equal(t, []int32{30, 40, 50, 60, 70}, obj.Value)

		// 添加数据到第17天
		ts = ts.Add(10 * 24 * time.Hour)
		obj.Add(100, ts)
		assert.Equal(t, 100, obj.GetTotal()) // 第2天的数据被移除
		assert.Equal(t, []int32{100, 0, 0, 0, 0}, obj.Value)
	})
}

func TestCounterByDays_Init(t *testing.T) {
	startDate := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	t.Run("same-length", func(t *testing.T) {
		values := []int32{1, 2, 3, 4, 5}
		obj := NewCounterByDaysWith(5, startDate, values)
		assert.Equal(t, 15, obj.GetTotal())
		assert.Equal(t, values, obj.Value)

		// 添加数据到第1天
		obj.Add(10, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC))
		assert.Equal(t, []int32{11, 2, 3, 4, 5}, obj.Value)
	})

	t.Run("less-length", func(t *testing.T) {
		values := []int32{1, 2, 3, 4}
		obj := NewCounterByDaysWith(5, startDate, values)
		assert.Equal(t, 10, obj.GetTotal())
		assert.Equal(t, []int32{0, 1, 2, 3, 4}, obj.Value)

		// 添加数据到第1天
		obj.Add(10, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC))
		assert.Equal(t, []int32{10, 1, 2, 3, 4}, obj.Value)
	})

	t.Run("greater-length", func(t *testing.T) {
		values := []int32{1, 2, 3, 4, 5, 6}
		obj := NewCounterByDaysWith(5, startDate, values)
		assert.Equal(t, 20, obj.GetTotal())
		assert.Equal(t, []int32{2, 3, 4, 5, 6}, obj.Value)

		// 添加数据到第1天
		obj.Add(10, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC))
		assert.Equal(t, []int32{12, 3, 4, 5, 6}, obj.Value)
	})
}
