package mitmproxy

import (
	"5ub5i5t/bughuntingtoolbox/mitmproxy/custom"
	"flag"
	"fmt"
	rawLog "log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kardianos/mitmproxy/addon"
	"github.com/kardianos/mitmproxy/cert"
	"github.com/kardianos/mitmproxy/proxy"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	debug    int
	version  bool
	certPath string

	addr         string
	webAddr      string
	ssl_insecure bool

	dump      string // dump filename
	dumpLevel int    // dump level

	mapperDir string
}

func loadConfig() *Config {
	config := new(Config)

	flag.IntVar(&config.debug, "debug", 0, "debug mode: 1 - print debug log, 2 - show debug from")
	flag.BoolVar(&config.version, "version", false, "show version")
	flag.StringVar(&config.addr, "addr", ":9080", "proxy listen addr")
	flag.StringVar(&config.webAddr, "web_addr", ":9081", "web interface listen addr")
	flag.BoolVar(&config.ssl_insecure, "ssl_insecure", true, "not verify upstream server SSL/TLS certificates.")
	flag.StringVar(&config.dump, "dump", "", "dump filename")
	flag.IntVar(&config.dumpLevel, "dump_level", 0, "dump level: 0 - header, 1 - header + body")
	flag.StringVar(&config.mapperDir, "mapper_dir", "", "mapper files dirpath")
	flag.StringVar(&config.certPath, "cert_path", "", "path of generate cert files")
	flag.Parse()

	return config
}

func StartProxy(context *gin.Context) {
	config := loadConfig()

	if config.debug > 0 {
		rawLog.SetFlags(rawLog.LstdFlags | rawLog.Lshortfile)
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	if config.debug == 2 {
		log.SetReportCaller(true)
	}
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	l, err := cert.NewPathLoader(config.certPath)
	if err != nil {
		log.Fatal(err)
	}
	ca, err := cert.New(l)
	if err != nil {
		log.Fatal(err)
	}

	opts := &proxy.Options{
		Debug:                 config.debug,
		Addr:                  config.addr,
		StreamLargeBodies:     1024 * 1024 * 5,
		InsecureSkipVerifyTLS: config.ssl_insecure,
		CA:                    ca,
	}

	p, err := proxy.NewProxy(opts)
	if err != nil {
		log.Fatal(err)
	}

	if config.version {
		fmt.Println("go-mitmproxy: " + p.Version)
		os.Exit(0)
	}

	log.Infof("go-mitmproxy version %v\n", p.Version)

	//p.AddAddon(&addon.LogAddon{})
	//p.AddAddon(&custom.ChangeHtml{})
	//p.AddAddon(web.NewWebAddon(config.webAddr))
	p.AddAddon(&custom.CustomLogAddon{})
	//p.AddAddon(&custom.CustomHttpxAddon{})

	if config.dump != "" {
		dumper := addon.NewDumperWithFilename(config.dump, config.dumpLevel)
		p.AddAddon(dumper)
	}

	if config.mapperDir != "" {
		mapper := addon.NewMapper(config.mapperDir)
		p.AddAddon(mapper)
	}

	log.Fatal(p.Start())
}

func StartProxyBasic() {
	config := loadConfig()

	if config.debug > 0 {
		rawLog.SetFlags(rawLog.LstdFlags | rawLog.Lshortfile)
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	if config.debug == 2 {
		log.SetReportCaller(true)
	}
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	l, err := cert.NewPathLoader(config.certPath)
	if err != nil {
		log.Fatal(err)
	}
	ca, err := cert.New(l)
	if err != nil {
		log.Fatal(err)
	}

	opts := &proxy.Options{
		Debug:                 config.debug,
		Addr:                  config.addr,
		StreamLargeBodies:     1024 * 1024 * 5,
		InsecureSkipVerifyTLS: config.ssl_insecure,
		CA:                    ca,
	}

	p, err := proxy.NewProxy(opts)
	if err != nil {
		log.Fatal(err)
	}

	if config.version {
		fmt.Println("go-mitmproxy: " + p.Version)
		os.Exit(0)
	}

	log.Infof("go-mitmproxy version %v\n", p.Version)

	//p.AddAddon(&addon.LogAddon{})
	//p.AddAddon(&custom.ChangeHtml{})
	//p.AddAddon(web.NewWebAddon(config.webAddr))
	p.AddAddon(&custom.CustomLogAddon{})
	p.AddAddon(&custom.SaveFlowAddon{})
	//p.AddAddon(&custom.CustomHttpxAddon{})

	if config.dump != "" {
		dumper := addon.NewDumperWithFilename(config.dump, config.dumpLevel)
		p.AddAddon(dumper)
	}

	if config.mapperDir != "" {
		mapper := addon.NewMapper(config.mapperDir)
		p.AddAddon(mapper)
	}

	log.Fatal(p.Start())
}
