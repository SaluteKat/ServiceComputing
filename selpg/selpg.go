package main

import (
	"bufio"
	//"flag"
	"github.com/spf13/pflag"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
)

type selpg_args struct {
	startPage        int
	endPage          int
	inFile           string
	pageLen          int
	pageType         bool //ture for -f, false for -l
	printDestination string
}

var programName []byte

func main() {
	args := new(selpg_args)
	receive_args(args)
	check_args(args)
	process_input(args)
}

func receive_args(args *selpg_args) {
	//提示信息，如果给出的参数不正确或者需要查看帮助 -help，那么会给出这里指定的字符串
	pflag.Usage = usage
	pflag.IntVarP(&(args.startPage), "start", "s", 0, "start page")
	pflag.IntVarP(&(args.endPage), "end", "e", 0, "end page")
	pflag.IntVarP(&(args.pageLen), "line", "l", 72, "page len")
	pflag.StringVarP(&(args.printDestination), "destionation", "d", "", "print destionation")
	pflag.BoolVarP(&(args.pageType), "type", "f", false, "type of print")
	pflag.Parse()
	//从第一个不能解析的参数开始，后面的所有参数都是无法解析的。即使后面的参数中含有预定义的参数
	//其他参数
	othersArg := pflag.Args()
	if len(othersArg) > 0 {
		args.inFile = othersArg[0]
	} else {
		args.inFile = ""
	}
}

func check_args(args *selpg_args) {
	if args.startPage == -1 || args.endPage == -1 {
		os.Stderr.Write([]byte("you should input -s -e at least\n"))
		os.Exit(0)
	}
	if args.startPage < 1 {
		os.Stderr.Write([]byte("invalid start page\n"))
		os.Exit(1)
	}
	if args.endPage < 1 || args.endPage > (math.MaxInt32-1) || args.endPage < args.startPage {
		os.Stderr.Write([]byte("invalid end page\n"))
		os.Exit(2)
	}
	if args.pageLen < 1 || args.pageLen > (math.MaxInt32-1) {
		os.Stderr.Write([]byte("invalid page length\n"))
		os.Exit(3)
	}
}

func process_input(args *selpg_args) {
	fin := os.Stdin
	fout := os.Stdout
	var (
		page_ctr int
		line_ctr int
		err error
		err1 error
		cmd *exec.Cmd
		stdin io.WriteCloser
	)
	/* set the input source */
	if args.inFile != "" {
		fin, err1 = os.Open(args.inFile)
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "Can not open input file \"%s\"\n",  args.inFile)
			os.Exit(11)
		}
	}

	if args.printDestination != "" {
		cmd = exec.Command("cat", "-n")
		stdin, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		stdin = nil
	}

	rd := bufio.NewReader(fin)
	writer := bufio.NewWriter(os.Stdout)
	if args.pageType == true {
		process_input_f(rd, writer, args, &page_ctr)
	} else {
		process_input_l(rd, writer, args, &page_ctr, &line_ctr)
	}

	if page_ctr < args.startPage {
		fmt.Fprintf(os.Stderr, "The start page %d is greater than total pages %d\n", args.startPage, page_ctr)
		os.Exit(12)
	} else if page_ctr < args.endPage {
			fmt.Fprintf(os.Stderr, "The end page %d is greater than total pages %d\n", args.endPage, page_ctr)
			os.Exit(13)
	}

	if args.printDestination != "" {
		stdin.Close()
		cmd.Stdout = fout
		cmd.Run()
	}
	fmt.Fprintf(os.Stderr,"\n----------------------\n  Process Finished!\n")
	fin.Close()
	fout.Close()
}

func process_input_f(reader *bufio.Reader, writer *bufio.Writer, args *selpg_args, pageCtr *int) {
	*pageCtr = 1
	for {
		char, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			os.Stderr.Write([]byte("read byte from Reader fail\n"))
			os.Exit(7)
		}
		if *pageCtr >= args.startPage && *pageCtr <= args.endPage {
			errW := writer.WriteByte(char)
			if errW != nil {
				os.Stderr.Write([]byte("Write byte to out fail\n"))
				os.Exit(8)
			}
			writer.Flush()
		}
		if char == '\f' {
			(*pageCtr)++
		}
	}
}

func process_input_l(reader *bufio.Reader, writer *bufio.Writer, args *selpg_args, pageCtr *int, lineCtr *int) {
	*lineCtr = 0
	*pageCtr = 1
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			//遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			os.Stderr.Write([]byte("read bytes from Reader error\n"))
			os.Exit(5)
		}
		*lineCtr++
		if *lineCtr > args.pageLen{
			*lineCtr = 0
			*pageCtr++
		}
		if *pageCtr >= args.startPage && *pageCtr <= args.endPage {
			_, errW := writer.Write(line)
			//_, errW := fmt.Fprintf(os.Stdout, "%s", line)
			if errW != nil {
				os.Stderr.Write([]byte("Write to file fail\n"))
				os.Exit(6)
			}
			writer.Flush()
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr,"Usage error!\n")
	fmt.Fprintf(os.Stderr,"Usage:")
	fmt.Fprintf(os.Stderr,"\tselpg -s Number -e Number [options] [filename]\n\n")
	fmt.Fprintf(os.Stderr,"\t-s=Number\t开始页数\n")
	fmt.Fprintf(os.Stderr,"\t-e=Number\t结束页数\n")
	fmt.Fprintf(os.Stderr,"\t-l=Number\t每页行数(可选)，默认为72\n")
	fmt.Fprintf(os.Stderr,"\t-f\t\t是否用换页符来换页(可选)\n")
	fmt.Fprintf(os.Stderr,"\t[filename]\t从文件读，省略为标准输入\n\n")
}