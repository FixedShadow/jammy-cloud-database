package utils

import (
	"fmt"
	"github.com/mackerelio/go-osstat/uptime"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

func GetUptimeInSeconds() (int64, error) {
	osUptime, err := uptime.Get()
	if err != nil {
		return -1, errors.Wrap(err, "exec uptime.Get failed")
	}
	return int64(osUptime / time.Second), nil
}

func Float64From32Bits(f float64) float64 {
	if f < 0 {
		return 0
	}
	return f
}

func Keep2Decimal(number float64) float64 {
	limitedString := fmt.Sprintf("%.2f", number)
	limitedNumber, _ := strconv.ParseFloat(limitedString, 64)
	return limitedNumber
}
