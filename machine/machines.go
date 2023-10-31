package machine

type Machine int32

var machineMap map[Machine]func(x, y, z int) int

func register(m Machine, fn func(x, y, z int) int) {
	machineMap[m] = fn
}

func GetAns(m Machine, x, y, z int) int {
	if _, ok := machineMap[m]; !ok {
		return -1
	}
	return machineMap[m](x, y, z)
}

func CheckMachines(ms []Machine) {
	println("check 1 start, machines = ", ms)
	f := true
	for _, m := range ms {
		if _, ok := machineMap[m]; !ok {
			println("machine ", m, " is unset.")
			f = false
		}
	}
	if f {
		println("pass check")
	}
}

func CheckAllMachines() {
	println("check all start")
	f := true
	for i := 1; i <= 24; i++ {
		m := Machine(i)
		if _, ok := machineMap[m]; !ok {
			println("machine ", m, " is unset.")
			f = false
		}
	}
	if f {
		println("pass check")
	}
}

func init() {
	machineMap = make(map[Machine]func(x, y, z int) int, 0)

	register(1, func(x, y, z int) int {
		if x == 1 {
			return 1
		}
		return 2
	})
	/*
		本验证者验证...the  number compared to 3
		蓝色小于3 蓝色等于3 蓝色大于3
	*/
	register(2, func(x, y, z int) int {
		if x < 3 {
			return 1
		}
		if x == 3 {
			return 2
		}
		return 3
	})

	register(3, func(x, y, z int) int {
		if y < 3 {
			return 1
		}
		if y == 3 {
			return 2
		}
		return 3
	})

	register(4, func(x, y, z int) int {
		if y < 4 {
			return 1
		}
		if y == 4 {
			return 2
		}
		return 3
	})

	register(5, func(x, y, z int) int {
		if x%2 == 1 {
			return 2
		}
		return 1
	})

	register(6, func(x, y, z int) int {
		if y%2 == 1 {
			return 2
		}
		return 1
	})

	/*
			本验证者验证...if z is even or odd
		is even
		is odd7
	*/

	register(7, func(x, y, z int) int {
		if z%2 == 1 {
			return 2
		}
		return 1
	})

	register(8, func(x, y, z int) int {
		c := 1
		if x == 1 {
			c++
		}
		if y == 1 {
			c++
		}
		if z == 1 {
			c++
		}
		return c
	})

	/*
			密码中3的个数
		没有3 一个3 2个3 3个3 9
	*/

	register(9, func(x, y, z int) int {
		c := 1
		if x == 3 {
			c++
		}
		if y == 3 {
			c++
		}
		if z == 3 {
			c++
		}
		return c
	})

	register(10, func(x, y, z int) int {
		c := 1
		if x == 4 {
			c++
		}
		if y == 4 {
			c++
		}
		if z == 4 {
			c++
		}
		return c
	})

	register(11, func(x, y, z int) int {
		if x == y {
			return 2
		}
		if x < y {
			return 1
		}
		return 3
	})

	register(12, func(x, y, z int) int {
		if x == z {
			return 2
		}
		if x < z {
			return 1
		}
		return 3
	})

	register(13, func(x, y, z int) int {
		if y == z {
			return 2
		}
		if y < z {
			return 1
		}
		return 3
	})

	register(14, func(x, y, z int) int {
		if x < y && x < z {
			return 1
		}
		if y < x && y < z {
			return 2
		}
		if z < x && z < y {
			return 3
		}
		return 0
	})

	register(15, func(x, y, z int) int {
		if x > y && x > z {
			return 1
		}
		if y > x && y > z {
			return 2
		}
		if z > x && z > y {
			return 3
		}
		return 0
	})
	/*
			偶数奇数本验证者验证...偶数和奇数的数量相比
		偶数更多(例如 245) | 奇数数字更多(例如 332)
	*/
	register(16, func(x, y, z int) int {
		a, b := 0, 0
		if x%2 == 1 {
			a++
		} else {
			b++
		}
		if y%2 == 1 {
			a++
		} else {
			b++
		}
		if z%2 == 1 {
			a++
		} else {
			b++
		}
		if a > b {
			return 2
		}
		return 1
	})

	/*
			偶数本验证者验证...密码中有多少偶数
		没有偶数1个偶数2个偶数3个偶数17
	*/

	register(17, func(x, y, z int) int {
		a, b := 0, 0
		if x%2 == 1 {
			a++
		} else {
			b++
		}
		if y%2 == 1 {
			a++
		} else {
			b++
		}
		if z%2 == 1 {
			a++
		} else {
			b++
		}
		return b + 1
	})

	/*
			偶数奇数本验证者验证...全部数字的总和是偶数还是奇数
		偶数奇数数字的总和是偶数数字的总和是奇数18
	*/
	register(18, func(x, y, z int) int {
		if (x+y+z)%2 == 0 {
			return 1
		}
		return 2
	})

	/*
			本验证者验证...the sum of  and  compared to 6
		蓝色和黄色数字的和小于6 | 蓝色和黄色的和等于6 | 蓝色和黄色数字的和大于6
	*/
	register(19, func(x, y, z int) int {
		if x+y < 6 {
			return 1
		}
		if x+y == 6 {
			return 2
		}
		return 3
	})

	/*
		本验证者验证...如果密码中有数字重复
		三个重复|两个重复|没有重复

	*/
	register(20, func(x, y, z int) int {
		if x == y && x == z {
			return 1
		}
		if x == y || y == z || x == z {
			return 2
		}
		return 3
	})

	/*
		本验证者验证...是否有数字正好出现两次
		否 是 21

	*/
	register(21, func(x, y, z int) int {
		if x == y && x == z {
			return 1
		}
		if x == y || y == z || x == z {
			return 2
		}
		return 1
	})

	/*
			本验证者验证...密码中的3个数字是升序、降序还是无序
		升序降序无序22
	*/
	register(22, func(x, y, z int) int {
		if x < y && y < z {
			return 1
		}
		if x > y && y > z {
			return 2
		}
		return 3
	})

	// // No. 1
	// register(1, func(x,y,z int) int{

	// 	return 0
	// })
}
