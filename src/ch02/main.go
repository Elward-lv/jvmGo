package main

import "fmt"

func main() {
	cmd := praseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}

// test
// PS C:\workspace\jvmGo> bin\ch01.exe -cp foo/bar MyApp arg1 arg2
// classpath:foo/bar class:MyApp args:[arg1 arg2]
