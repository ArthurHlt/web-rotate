// Command simple is a chromedp example demonstrating how to do a simple google
// search.
package main

import (
	"context"
	"flag"
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

func main() {
	confPathPtr := flag.String("config", "config.yml", "set config path")
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
	tasks := make(chromedp.Tasks, len(pages))
	for i, p := range pages {
		tasks[i] = pageToTask(p)
	}
	return tasks
}

func pageToTask(page *Page) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(page.Url),
		chromedp.ActionFunc(cartridge(page)),
		chromedp.Sleep(time.Duration(page.Duration)),
	}
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
