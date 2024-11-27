package humanize

import (
	"fmt"
	"strconv"
	"time"
)

func S(b uint64) string {
	return strconv.FormatUint(b, 10)
}

func TimeStamp(b uint64) string {
	return time.Unix(int64(b/1e9), 0).Format(time.RFC3339)
}

func Number(b uint64) string {
	switch {
	case b < 1000:
		return fmt.Sprintf("%d", b)
	case b < 1000_000:
		return fmt.Sprintf("%.1fK", float64(b)/1000)
	case b < 1000_000_000:
		return fmt.Sprintf("%.1fM", float64(b)/1000_000)
	case b < 1000_000_000_000:
		return fmt.Sprintf("%.1fG", float64(b)/1000_000_000_000)
	case b < 1000_000_000_000_000:
		return fmt.Sprintf("%.1fT", float64(b)/1000_000_000_000_000)
	case b < 1000_000_000_000_000_000:
		return fmt.Sprintf("%.1fP", float64(b)/1000_000_000_000_000_000)
	default:
		return fmt.Sprintf("%.1fE", float64(b)/1000_000_000_000_000_000_000)
	}
}

func Size(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
