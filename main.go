package main

import (
	"encoding/json"
	"fmt"
	"turing_machine/machine"
)

func Cal(machines []machine.Machine) []int {
	ansList := []int{}
	// ansMap := make(map[int][]int)
	ansCount := make([]int, 1000000)
	nearCount := make([]int, 1000000)

	// check rules

	nums := []int{1, 2, 3, 4, 5}
	for _, i := range nums {
		for _, j := range nums {
			for _, k := range nums {
				valid := true
				ans := 0
				for _, m := range machines {
					a := machine.GetAns(machine.Machine(m), i, j, k)
					ans = ans*10 + a
					if a == 0 {
						valid = false
					}
					if a < 0 {
						println(i, j, k, m, a, ans)
						panic("unset machine")
					}
				}
				// ansMap[ans] = append(ansMap[ans], i*100+j*10+k)
				ind := 1
				for range machines {
					near := ans/(ind*10)*(ind*10) + ans%ind
					nearCount[near]++
					ind *= 10
				}

				if !valid {
					continue
				}

				ansCount[ans]++
			}
		}
	}

	// println("ans below:")

	for _, i := range nums {
		for _, j := range nums {
			for _, k := range nums {
				fl := true
				valid := true
				ans := 0
				for _, m := range machines {
					a := machine.GetAns(machine.Machine(m), i, j, k)
					ans = ans*10 + a
					// println(i, j, k, m)
					if a == 0 {
						valid = false
					}
				}
				if !valid {
					continue
				}
				if ansCount[ans] > 1 {
					fl = false
					// println(i, j, k, "failed, because ans = ", ans, "has multiple ans:", ansCount[ans])
					// println("multiple ans are: ", JsonMarshal(ansMap[ans]))
				}
				ind := 1
				for range machines {
					mm := ind * 10
					near := ans/mm*mm + ans%ind
					for near < 0 {

						println(ans, ind, mm, near)
						panic("invalid near")
					}
					if nearCount[near] == 1 {
						fl = false
						// println(i, j, k, "failed, because ans = ", ans, "unused machine, near = ", near)
						break
					}
					ind *= 10
				}
				if fl {
					ansList = append(ansList, i*100+j*10+k)
				}
			}
		}
	}
	return ansList
}

/*
machine  23  is unset.
machine  24  is unset.
*/
func main() {

	// machines := []machine.Machine{2, 7, 13, 14}
	// ansList := Cal(machines)
	// println(JsonMarshal(ansList))

	// CalFourMachines()

	// CalFiveMachines()

	CalSixMachines()
}

func CalFourMachines() {
	ansCount := make([]int, 10000)
	total := 0

	for i := 1; i <= 22; i++ {
		for j := i + 1; j <= 22; j++ {
			for k := j + 1; k <= 22; k++ {
				for l := k + 1; l <= 22; l++ {
					ms := []machine.Machine{
						machine.Machine(i),
						machine.Machine(j),
						machine.Machine(k),
						machine.Machine(l),
					}
					ansList := Cal(ms)
					if len(ansList) > 20 {
						println("len > 20, len = ", len(ansList), ", machines = ", MarshalMachines(ms))
					}
					// if len(ansList) == 1 {
					// 	// if i == 1 && j == 13 && k == 17 && l == 19 {
					// 	// 	println(JsonMarshal(ansList))
					// 	// }
					// 	// println(PrintMachinesAsTuple(ms))

					// 	// println("len ==  1, machines = ", MarshalMachines(ms))
					// }
					ansCount[len(ansList)]++
					total++
				}
			}
		}
	}

	for i := 0; i <= 30; i++ {
		println("ans with ", i, ": ", ansCount[i])
	}
	println("total = ", total)
}

func CalFiveMachines() {
	ansCount := make([]int, 10000)
	total := 0

	for i := 1; i <= 22; i++ {
		for j := i + 1; j <= 22; j++ {
			for k := j + 1; k <= 22; k++ {
				for l := k + 1; l <= 22; l++ {
					for m := l + 1; m <= 22; m++ {
						ms := []machine.Machine{
							machine.Machine(i),
							machine.Machine(j),
							machine.Machine(k),
							machine.Machine(l),
							machine.Machine(m),
						}
						ansList := Cal(ms)
						if len(ansList) > 20 {
							println("len > 20, len = ", len(ansList), ", machines = ", MarshalMachines(ms))
						}
						// if len(ansList) == 1 {
						// 	// if i == 1 && j == 13 && k == 17 && l == 19 {
						// 	// 	println(JsonMarshal(ansList))
						// 	// }
						// 	println(PrintMachinesAsTuple(ms))

						// 	// println("len ==  1, machines = ", MarshalMachines(ms))
						// }
						ansCount[len(ansList)]++
						total++
					}
				}
			}
		}
	}

	for i := 0; i <= 30; i++ {
		println("ans with ", i, ": ", ansCount[i])
	}
	println("total = ", total)
}

func CalSixMachines() {
	ansCount := make([]int, 10000)
	total := 0

	for i := 1; i <= 22; i++ {
		for j := i + 1; j <= 22; j++ {
			for k := j + 1; k <= 22; k++ {
				for l := k + 1; l <= 22; l++ {
					for m := l + 1; m <= 22; m++ {
						for n := m + 1; n <= 22; n++ {
							ms := []machine.Machine{
								machine.Machine(i),
								machine.Machine(j),
								machine.Machine(k),
								machine.Machine(l),
								machine.Machine(m),
								machine.Machine(n),
							}
							ansList := Cal(ms)
							if len(ansList) > 20 {
								println("len > 20, len = ", len(ansList), ", machines = ", MarshalMachines(ms))
							}
							// if len(ansList) == 1 {
							// 	// if i == 1 && j == 13 && k == 17 && l == 19 {
							// 	// 	println(JsonMarshal(ansList))
							// 	// }
							// 	println(PrintMachinesAsTuple(ms))

							// 	// println("len ==  1, machines = ", MarshalMachines(ms))
							// }
							ansCount[len(ansList)]++
							total++
						}
					}
				}
			}
		}
	}

	for i := 0; i <= 30; i++ {
		println("ans with ", i, ": ", ansCount[i])
	}
	println("total = ", total)
}

func JsonMarshal(i interface{}) string {
	bytes, _ := json.Marshal(i)
	return string(bytes)
}

func MarshalMachines(ms []machine.Machine) string {
	res := ""
	for _, m := range ms {
		res += fmt.Sprintf("%02d", m)
	}
	return res
}

func PrintMachinesAsTuple(ms []machine.Machine) string {
	res := ""
	for _, m := range ms {
		if len(res) > 0 {
			res += ", "
		}
		res += fmt.Sprintf("%d", m)
	}
	return "(" + res + "),"
}
