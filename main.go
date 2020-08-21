package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	showVersion = flag.Bool("version", false, "Print version information.")
	files       = flag.String("files", "", "Specify certificate files with comma separated values.")
	out         = flag.String("out", "local_ssl_exporter.prom", "Output file path.")
)

const version = "development"

type result struct {
	FilePath string
	UnixTime int64
	Days     int
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Fprintln(os.Stdout, "local ssl exporter:", version)
		os.Exit(0)
	}

	log.Println("Starting local-ssl-exporter")

	certFiles := strings.Split(*files, ",")
	certFiles = checkFiles(certFiles)
	results := checkCertificates(certFiles)

	file, err := os.OpenFile(*out, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, result := range results {
		fmt.Fprintf(file, "local_ssl_unixtime{file=\"%s\"} %d\n", result.FilePath, result.UnixTime)
		fmt.Fprintf(file, "local_ssl_sub_days{file=\"%s\"} %d\n", result.FilePath, result.Days)
	}

	log.Println("End local-ssl-exporter")
}

func checkFiles(certFiles []string) []string {
	checkFiles := make([]string, 0)
	for _, filename := range certFiles {
		if isExist(filename) {
			checkFiles = append(checkFiles, filename)
		}
	}
	return checkFiles
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func checkCertificates(certFiles []string) []*result {
	ch := make(chan *result)
	for _, file := range certFiles {
		go func(file string) {
			line := strings.Split("openssl x509 -noout -enddate -in "+file, " ")
			cmd := exec.Command(line[0], line[1:]...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			out = bytes.Trim(out, "notAfter=")
			out = bytes.TrimSpace(out)

			const layout = "Jan 2 15:04:05 2006 MST"
			notAfter, err := time.Parse(layout, string(out))
			if err != nil {
				log.Fatal(err)
			}
			sub := int(notAfter.Sub(time.Now()).Hours() / 24)

			ch <- &result{
				FilePath: file,
				UnixTime: notAfter.Unix(),
				Days:     sub,
			}
		}(file)
	}

	results := make([]*result, 0)
	for i := 0; i < len(certFiles); i++ {
		results = append(results, <-ch)
	}
	return results
}
