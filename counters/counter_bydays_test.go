package counters

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCounterByDays(t *testing.T) {
	startDate := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	obj := NewCounterByDays(5, startDate)

	// 添加数据到第1天
	obj.Add(10, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC))
	assert.Equal(t, 10, obj.GetTotal())
	assert.Equal(t, []int{10, 0, 0, 0, 0}, obj.Value)

	// 添加数据到第2天
	obj.Add(20, time.Date(2024, 1, 2, 9, 0, 0, 0, time.UTC))
	assert.Equal(t, 30, obj.GetTotal())
	assert.Equal(t, []int{10, 20, 0, 0, 0}, obj.Value)

	// 添加数据到第6天，触发滑动窗口
	obj.Add(30, time.Date(2024, 1, 6, 15, 0, 0, 0, time.UTC))
	assert.Equal(t, 50, obj.GetTotal()) // 第1天的数据被移除
	assert.Equal(t, []int{20, 0, 0, 0, 30}, obj.Value)

	// 添加数据到第7天
	obj.Add(40, time.Date(2024, 1, 7, 8, 0, 0, 0, time.UTC))
	assert.Equal(t, 70, obj.GetTotal()) // 第2天的数据被移除
	assert.Equal(t, []int{0, 0, 0, 30, 40}, obj.Value)

	// 添加数据到第17天
	obj.Add(100, time.Date(2024, 1, 17, 8, 0, 0, 0, time.UTC))
	assert.Equal(t, 100, obj.GetTotal()) // 第2天的数据被移除
	assert.Equal(t, []int{100, 0, 0, 0, 0}, obj.Value)
}
