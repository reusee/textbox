package textbox

import (
	"fmt"
	"log"
	"os"
)

var (
	sp     = fmt.Sprintf
	pt     = fmt.Printf
	logger = log.New(os.Stdout, "", 0)
	lg     = logger.Printf
)

func RuneWidth(r rune) int {
	switch {
	case r >= 0x4e00 && r <= 0x9fff, // CJK综合汉字
		r >= 0x2f00 && r <= 0x2fdf,   // 部首
		r >= 0x2e80 && r <= 0x2eff,   // 部首辅助
		r >= 0x2ff0 && r <= 0x2fff,   // 汉字构成记述文字
		r >= 0x3000 && r <= 0x303f,   // CJK标点
		r >= 0x3040 && r <= 0x309f,   // 平假名
		r >= 0x30A0 && r <= 0x30ff,   // 片假名
		r >= 0x31c0 && r <= 0x31ef,   // 笔画
		r >= 0x31f0 && r <= 0x31ff,   // 片假名扩展
		r >= 0x3400 && r <= 0x4dbf,   // 汉字扩展A
		r >= 0xf900 && r <= 0xfaff,   // CJK 互换汉字
		r >= 0xff00 && r <= 0xffef,   // 全角
		r >= 0x20000 && r <= 0x2ffff, // 汉字扩展
		r >= 0xe0100 && r <= 0xe01ef, // 异体汉字
		r >= 0x1b000 && r <= 0x1b0ff, // 假名辅助
		r >= 0x30000 && r <= 0x3ffff:
		return 2
	default:
		return 1
	}
}

func DisplayWidth(s string) (ret int) {
	for _, r := range s {
		ret += RuneWidth(r)
	}
	return
}

func MakeMove(fn func() Point, x, y int) func() Point {
	return func() Point {
		return fn().Move(x, y)
	}
}
