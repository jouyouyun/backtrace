package main

import (
	"flag"
	"log"
	"time"

	"github.com/fatih/color"
)

var (
	confFile    = flag.String("f", "./config.yaml", "The backtrace configuration file")
	targetType  = flag.String("t", "mainland", "The backtrace target type")
	useFirstHop = flag.String("n", "yes", "Using the first hop node asn as result, available value: yes or no")

	homepage = "项目地址：github.com/zhanghanyun/backtrace"
	title    = "正在测试三网回程路由..."
	timeout  = time.Duration(10)
)

func init() {
	flag.Parse()

	cfg, _ := loadConfig(*confFile)
	if cfg == nil {
		return
	}

	if cfg.Timeout > 0 {
		timeout = time.Duration(cfg.Timeout)
	}
	if len(cfg.Homepage) != 0 {
		homepage = cfg.Homepage
	}
	if len(cfg.Title) != 0 {
		title = cfg.Title
	}
	if len(cfg.IPList) != 0 {
		infos := cfg.makeIPInfoList(*targetType)
		rIp = infos.ListIP()
		rName = infos.ListName()
	}
	if len(cfg.ASNList) != 0 {
		m = cfg.makeASNMap()
	}
}

func main() {
	var (
		s = make([]string, len(rIp))
		c = make(chan Result)
		t = time.After(time.Second * timeout)
	)

	head := color.New(color.FgHiBlue).Add(color.Bold).SprintFunc()
	note := color.New(color.FgGreen).SprintFunc()
	log.Println(head(homepage))
	log.Println(note(title))

	for i := range rIp {
		go trace(c, i)
	}

loop:
	for range s {
		select {
		case o := <-c:
			s[o.i] = o.s
		case <-t:
			break loop
		}
	}

	for _, r := range s {
		log.Println(r)
	}
}
