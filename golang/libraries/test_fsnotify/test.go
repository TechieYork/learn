package main

import (
	//"github.com/golang/glog"
	"github.com/urfave/cli"
	"os"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"bufio"
	"time"
)

func begin_watch()  {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		fmt.Println("init fsnotify failed")
	}

	fmt.Println("init fsnotify successed")

	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			fmt.Println("begin watch")
			select {
			case event := <-watcher.Events:
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()

	fmt.Println("Add watch dir")
	err = watcher.Add("./")

	if err != nil {
		fmt.Println(err)
	}
	<-done

	fmt.Println("done")
}

func deal_cli()  {
	app := cli.NewApp()

	app.Name = "Test"
	app.Usage = "Test cli"
	app.Version = "1.0.0"

	var language string
	var test string

	app.Flags = []cli.Flag {
    cli.StringFlag{
      Name:        "lang",
      Value:       "english",
      Usage:       "language for the greeting",
      Destination: &language,
    },
	cli.StringFlag{
      Name:        "test",
      Value:       "test",
      Usage:       "fors test",
      Destination: &test,
    },
  }

	app.Action = func(c *cli.Context) error{
		name := "someone"

		fmt.Println(len(c.Args()))

		for _, arg := range c.Args(){
			fmt.Printf("Arg:%s\r\n", arg)
		}



		if c.NArg() > 0 {
      		name = c.Args()[0]
    	}
    	if language == "spanish" {
      		fmt.Println("Hola", name)
    	} else {
      		fmt.Println("Hello", name)
    	}

		return nil
	}

	app.Run(os.Args)
}

func readFile(filename string, log chan string)  {
	file, err := os.OpenFile(filename, os.O_APPEND | os.O_RDWR, 0755)

	if err != nil {
		fmt.Println("file open failed")
		return
	}

	reader := bufio.NewReader(file)

	for {
		line, prefix, err := reader.ReadLine()

		if err != nil {
			time.Sleep(1000 * time.Millisecond)
			continue
		}

		fmt.Println(filename + ":" + string(line), prefix, err)

		log <- string(line)
	}
}


func main()  {
	//deal_cli()

	c := make(chan string, 100)

	go func() {
		for  {
			select {
			case log := <- c:
				fmt.Println(log)
			default:
				fmt.Println("no log")
				time.Sleep(1 * time.Second)
			}
		}
	}()


	go readFile("./1.txt", c)
	//go readFile("./2.txt", c)
	//go readFile("./3.txt", c)
	//go readFile("./4.txt", c)
	//go readFile("./5.txt", c)
	//go readFile("./6.txt", c)
	//go readFile("./7.txt", c)
	//go readFile("./8.txt", c)
	//go readFile("./9.txt", c)
	//go readFile("./0.txt", c)

	for  {
		time.Sleep(10 * time.Second)
	}

	//defer glog.Flush()

	//glog.Infoln("test")
}
