// Command simple is a chromedp example demonstrating how to do a simple google
// search.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
	"github.com/chromedp/chromedp/runner"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Version string

func main() {
	confPathPtr := flag.String("config", "config.yml", "set config path")
	versionReqPtr := flag.Bool("version", false, "see version")
	flag.Parse()
	if *versionReqPtr {
		fmt.Println("web-rotate " + Version)
		os.Exit(0)
	}
	var err error
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-sigs
		os.Exit(0)
	}()
	conf, err := retrieveConfig(*confPathPtr)
	if err != nil {
		log.Fatal(err)
	}
	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()
	opts := []chromedp.Option{
		chromedp.WithRunnerOptions(
			runner.CmdOpt(defaultUserData),
			runner.Flag("start-fullscreen", conf.Fullscreen),
		),
	}
	ln, err := net.Listen("tcp", "127.0.0.1:9222")
	if err != nil {
		opts = append(opts, chromedp.WithTargets(client.New().WatchPageTargets(ctxt)))
	}
	if ln != nil {
		ln.Close()
	}

	// create chrome instance
	c, err := chromedp.New(ctxt, opts...)
	if err != nil {
		log.Fatal(err)
	}

	actions := pagesToTasks(conf.Pages)
	go func() {
		for {
			err = c.Run(ctxt, actions)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	c.Wait()
}

func pagesToTasks(pages []*Page) chromedp.Tasks {
	b, _ := yaml.Marshal(pages[0])
	fmt.Println(string(b))
	network.Enable()
	tasks := chromedp.Tasks{
		network.Enable(),
	}

	for _, p := range pages {
		tasks = append(tasks, pageToTask(p)...)
	}
	return tasks
}

func pageToTask(page *Page) chromedp.Tasks {
	tasks := make(chromedp.Tasks, 0)
	if len(page.Headers) > 0 {
		tasks = append(tasks, network.SetExtraHTTPHeaders(page.Headers))
	}
	tasks = append(tasks, chromedp.Navigate(page.Url))
	if page.Cartridge != "" {
		tasks = append(tasks, chromedp.ActionFunc(cartridge(page)))
	}
	tasks = append(tasks, chromedp.Sleep(time.Duration(page.Duration)))
	return tasks
}

func retrieveConfig(configPath string) (*Config, error) {
	config := &Config{}
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(b, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
