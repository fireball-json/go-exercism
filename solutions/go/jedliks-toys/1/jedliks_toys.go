package jedlik

import "fmt"

// Drive 让车前进一次：行驶距离增加 speed（米），电池减少 batteryDrain（%）。
// 必须用指针接收者，否则对接收者的修改不会传回调用方。
// 如果剩余电量低于耗电量（驾驶后电池会变成负数），就放弃本次驾驶，车身状态保持不变。
func (c *Car) Drive() {
	if c.battery < c.batteryDrain {
		return
	}
	c.distance += c.speed
	c.battery -= c.batteryDrain
}

// DisplayDistance 返回 LED 屏上显示的行驶总距离。
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery 返回 LED 屏上显示的剩余电量百分比。
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish 判断车能否在电量耗尽前跑完整条赛道。
// 每驾驶一次前进 speed 米、消耗 batteryDrain 电量，
// 所以剩余电量最多能支撑 (battery / batteryDrain) 次驾驶，总共能跑该次数乘以 speed 米。
// 电量刚好够跑完赛道时（剩余距离 == 0）也算能完成。
func (c Car) CanFinish(trackDistance int) bool {
	maxDistance := (c.battery / c.batteryDrain) * c.speed
	return maxDistance >= trackDistance
}
