package calendar

import (
        "bufio"
)

func init() {
        registry["Day5"] = Day5
}

func Day5(input *bufio.Scanner, part int) (result int) {
        if part == 1 {

        } else {

        }

        for input.Scan() {
                log.Debug(input.Text())
        }

        return
}
