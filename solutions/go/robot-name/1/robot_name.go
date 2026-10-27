package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// totalNames 是所有可能名字的数量：26*26*10*10*10 = 676000
// （测试文件里已有一个 maxNames 常量，这里换个名字避免冲突）
const totalNames = 26 * 26 * 10 * 10 * 10

// used 是「已分配名字」的全局集合。
//
// 必须是包级全局的，不能放在 Robot 结构体里：
// 不同机器人之间的名字也必须互不相同，而每个机器人各持一份池子的话，
// 两个机器人会各自随机到同一个名字，导致重名。
var used = make(map[string]bool)

// Robot 表示一台机器人。
type Robot struct {
	name string
}

// Name 返回机器人的名字，必要时生成一个新的。
// 同一个机器人多次调用返回的是同一个名字。
func (r *Robot) Name() (string, error) {
	// 已经有名字了，直接返回（名字在两次调用之间保持不变）
	if r.name != "" {
		return r.name, nil
	}
	// 名字空间已耗尽，无法再分配
	if len(used) >= totalNames {
		return "", errors.New("no more robot names available")
	}
	// 随机生成直到撞上一个没被用过的名字
	for {
		name := makeName()
		if used[name] {
			continue
		}
		used[name] = true
		r.name = name
		return name, nil
	}
}

// Reset 把机器人重置为出厂设置：清除它自己的名字，
// 下次调用 Name() 时会拿到一个新的随机名字。
//
// 注意旧名字不会重新放回可用池中——它已经被标记为「用掉了」，
// 这样保证任何名字都绝不会被分配给第二台机器人。
func (r *Robot) Reset() {
	r.name = ""
}

// makeName 生成一个形如 "AB123" 的随机名字：两位大写字母 + 三位数字。
func makeName() string {
	// 'A' + 0~25 得到 A~Z
	letter1 := rune('A' + rand.Intn(26))
	letter2 := rune('A' + rand.Intn(26))
	// %03d 补足三位，不足则前面补 0（保证是 000~999）
	return fmt.Sprintf("%c%c%03d", letter1, letter2, rand.Intn(1000))
}
