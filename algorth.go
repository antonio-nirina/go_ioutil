package main

import(
	"fmt"
	"log"
	"os/exec"
)

func tesst() {
	ts := [...]float32{7, -10, 13, 8, 4, -7.2, -12, -3.7, 3.5, -9.6, 6.5, -1.7, -6.2, 7, 0.5, -0.3}
	var pos []float32
	var neg []float32

	for _, val := range ts {
		if val > 0 {
			pos = append(pos, val)
		} else {
			neg = append(neg, val)
		}

	}

	minPos := pos[0]
	maxNeg := neg[0]

	for _, valp := range pos {
		if minPos > valp {
			minPos = valp
		}
	}

	for _, valn := range neg {
		if maxNeg < valn {
			maxNeg = valn
		}
	}

	if minPos > (-1)*maxNeg {
		fmt.Println(maxNeg)
	} else {
		fmt.Println(minPos)
	}
}

func printer() {
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 50; j++ {
			if j > i {
				for k := 1; k <= i; k++ {
					fmt.Print("*")
				}
				break
			} else {
				fmt.Print("*")
			}
		}

		fmt.Println("")
	}

	for m := 1; m <= 5; m++ {
		for n := 1; n <= 25; n++ {
			if n >= 20 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println("")
	}
}

func monitor() {
	cmd := exec.Command("pwd")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(string(out))
}