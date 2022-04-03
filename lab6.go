//พัฒนาโดย นาย โชคชัย แจ่มน้อย
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	pages  []string
	frames []string
	f1     []string
	f2     []string
	f3     []string
	f4     []string
	fault  int
)

//ค่าเริ่มต้น
func initx() {
	pages = make([]string, 100)
	frames = make([]string, 3)
	f1 = make([]string, 100)
	f2 = make([]string, 100)
	f3 = make([]string, 100)
	f4 = make([]string, 100)

	fault = 0
}

//แสดงตาราง
func showTable() {
	fmt.Println()
	for i := range pages {
		if pages[i] != "" {
			fmt.Printf("----")
		}
	}

	fmt.Println()
	for i := range pages {
		if pages[i] != "" {
			fmt.Printf("| %s ", pages[i])
		}
	}

	fmt.Println()
	for i := range pages {
		if pages[i] != "" {
			fmt.Printf("----")
		}
	}

	fmt.Println()
	for i := range pages {
		if pages[i] != "" {
			fmt.Printf("| %s ", f1[i])
		}
	}

	fmt.Println()
	for i := range pages {
		if pages[i] != "" {
			fmt.Printf("| %s ", f2[i])
		}
	}

	fmt.Println()
	for i := range pages {
		if pages[i] != "" {
			fmt.Printf("| %s ", f3[i])
		}
	}

	fmt.Println()
	for i := range pages {
		if pages[i] != "" {
			fmt.Printf("| %s ", f4[i])
		}
	}

	fmt.Println()
	fmt.Printf("Page fualt = %d\n", fault)

}

//รับค่าตัวเลข
func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

//ฟังก์ชัน LRU
func LRU() {
	fault = 0          //ทุกรอบให้ reset ค่าเป็น 0
	not_fault := false //ทุกรอบให้ reset ค่าเป็น false ไว้เช็คสถานะว่าเกิด page fault หรือไม่

	//ให้ทุกค่าในเฟรมเป็น -
	for j := range frames {
		frames[j] = "-"
	}

	for i := range pages {
		f4[i] = "-"       //แต่ละรอบให้ส่วนแสดง page fault เริ่มต้นเป็น -
		not_fault = false //ทุกรอบให้ reset ค่าเป็น false ไว้เช็คสถานะว่าเกิด page fault หรือไม่

		//รอบ 0-2 ยังไงก็เกิด page fault
		if i < 3 {
			frames[i] = pages[i]
			fault++
			f1[i] = frames[0]
			f2[i] = frames[1]
			f3[i] = frames[2]
			f4[i] = "f" //ส่วนแสดง page fault เป็น "f" นั้นคือเกิด page fault

			//ตั้งแต่รอบที่ 3 - ตัวเลขสุดท้ายใน pages
		} else if pages[i] != "" {

			for j := range frames {
				if frames[j] == pages[i] { //เช็คว่าตัวเลขที่เรากำลังสนใจอยู่ มีใน frames ไหม ถ้ามีจะไม่เกิด page fault
					f1[i] = frames[0]
					f2[i] = frames[1]
					f3[i] = frames[2]
					not_fault = true //ไม่เกิด page fault = จริง
					break
				}
			}

			//จะเข้าเงื่อนไขนี้เมื่อการันตีได้ว่าจะเกิด page fault เพราะไม่เข้าเงื่อนไขที่ Line 124
			if not_fault == false {
				//ตัวแปร สำหรับไว้จำตำแหน่งของตัวเลขใน frames เพื่อไว้เปรียบเทียบว่าตัวไหน ไกลที่สุด (LRU)
				no_f1 := 0
				no_f2 := 0
				no_f3 := 0

				fault++ //การันตีได้ว่าจะเกิด page fault ให้ เพิ่มค่า fault+1

				//ให้วนถอยหลังจากตำแหน่งที่ทำอยู่ และจำตำแหน่งเอาไว้
				for j := (i - 1); j >= 0; j-- {
					if pages[j] == frames[0] {
						no_f1 = j
					} else if pages[j] == frames[1] {
						no_f2 = j
					} else if pages[j] == frames[2] {
						no_f3 = j
					}
					//เมื่อได้ค่าตำแหน่งครบทั้ง 3 ตัวแล้ว ให้หลุดออกจาก loop
					if no_f1 != 0 && no_f2 != 0 && no_f3 != 0 {
						break
					}
				}

				//เปรียบเทียบว่าตำแหน่งไหนไกลสุด ให้แทนที่ตำแหน่งนั้นใน frames
				if no_f1 < no_f2 && no_f1 < no_f3 {
					frames[0] = pages[i]
				} else if no_f2 < no_f1 && no_f2 < no_f3 {
					frames[1] = pages[i]
				} else {
					frames[2] = pages[i]
				}

				//จำค่าเอาไว้เพื่อไว้นำไปแสดงผล
				f1[i] = frames[0]
				f2[i] = frames[1]
				f3[i] = frames[2]
				f4[i] = "f"
			}

		}
	}

}

func main() {
	initx()
	for {
		fmt.Printf("\nreq>")
		command := getCommand()
		commandx := strings.Split(command, " ")
		for i := range commandx {
			pages[i] = commandx[i]
		}
		LRU()
		showTable()
		initx()
	}
}
