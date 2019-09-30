package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"

	flag "github.com/spf13/pflag"
)

type selpgArgs struct {
	startPage   int    //Start page
	endPage     int    //End page
	filename    string //Input filename
	linePerPage int    //Line per page
	pageType    bool   //Whether paging by paging break
	dest        string //Print destionation
}

func main() {
	fin, err = os.Cr(args.filename)
	args := new(selpgArgs)
	get(args)
	processArgs(args)
	processInput(args)
}

func get(args *selpgArgs) { //Parsing args in commands
	//Bind flag to selpgArgs struct
	flag.IntVarP(&args.startPage, "startpage", "s", -1, "Start page")
	flag.IntVarP(&args.endPage, "endpage", "e", -1, "End page")
	flag.IntVarP(&args.linePerPage, "lines", "l", 72, "Lines per page")
	flag.BoolVarP(&args.pageType, "flag", "f", false, "Whether paging by paging break")
	flag.StringVarP(&args.dest, "destionation", "d", "", "Print destionation")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"\nUSAGE: %s -s startPage -e endPage [-l linePerPage | -f] [-d dest] input_file [>output_file] [2>error_file]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	//Get the file name
	if len(flag.Args()) == 1 {
		args.filename = flag.Args()[0]
	} else {
		args.filename = ""
	}
}

func processArgs(args *selpgArgs) { //Processing input errors
	if args == nil {
		fmt.Fprintln(os.Stderr, "\n[Error]The args is nil!")
		os.Exit(1)
	} else if args.startPage == -1 || args.endPage == -1 {
		fmt.Fprintln(os.Stderr, "\n[Error]The startPage and endPage are not allowed empty!")
		os.Exit(2)
	} else if args.startPage < 0 || args.startPage > (math.MaxInt32-1) {
		fmt.Fprintln(os.Stderr, "\n[Error]The startPage is invalid!")
		os.Exit(3)
	} else if args.endPage < 0 || args.endPage > (math.MaxInt32-1) {
		fmt.Fprintln(os.Stderr, "\n[Error]The endPage is invalid!")
		os.Exit(4)
	} else if args.pageType != false && args.linePerPage != 72 {
		fmt.Fprintln(os.Stderr, "\n[Error]-f and -l are not allowed to exist at the same time!")
		os.Exit(5)
	}
}

func processInput(args *selpgArgs) {
	var fin io.Reader

	//Open file
	if args.filename == "" {
		fin = os.Stdin
	} else {
		var err error
		fin, err = os.Open(args.filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "[Error]Failed to open input file!")
			os.Exit(0)
		}
	}

	//Page files
	var buf []string
	scanner := bufio.NewScanner(fin)
	if !args.pageType {
		//-l类型
		count := 0
		for scanner.Scan() {
			line := scanner.Text()
			if count/args.linePerPage == args.endPage {
				break
			}
			if count/args.linePerPage+1 >= args.startPage {
				buf = append(buf, line)
			}
			count++
		}
	} else {
		//-f类型
		count := 0
		onSp := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			for i := 0; i < len(data); i++ {
				if data[i] == '\f' {
					return i + 1, data[:i], nil
				}
			}
			if atEOF {
				return 0, data, bufio.ErrFinalToken
			}
			return 0, nil, nil
		}
		scanner.Split(onSp)
		for scanner.Scan() {
			line := scanner.Text()
			if count == args.endPage {
				break
			}
			if count+1 >= args.startPage {
				buf = append(buf, line)
			}
			count++
		}
	}

	//Output contents of the file
	if args.dest == "" {
		//Output to stdout
		for _, v := range buf {
			fmt.Fprintln(os.Stdout, v)
		}
	} else {
		//Output to a file
		cmd := exec.Command("./" + args.dest)
		inpipe, err := cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer inpipe.Close()
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()

		for _, v := range buf {
			inpipe.Write([]byte(v))
		}
	}
}
