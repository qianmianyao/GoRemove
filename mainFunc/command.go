package mainFunc

// Command 判断命令行参数
func Command(par *[]string) {
	defer func() {
		if err := recover(); err != nil {
			Help() // 如果后面什么参数都没有
		}
	}()
	pars := *par
	switch pars[1] {
	// -r 选项直接删除
	case "-r":
		file := pars[2:]
		Remover(&file)
	// -m 选项直接回收
	case "-m":
		file := pars[2:]
		Recycle(&file)
	// -e 清空回收站
	case "-e":
		Empty()
	// -b 恢复文件
	case "-b":
		Restore()
	// -h help
	case "-h":
		Help()
	// 初始化
	case "init":
		Init()
	default:
		file := pars[1:]
		MainFunc(&file)
	}
}
