package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/chengtiesheng/ContainerAnalyzer/analyse"
)

var (
	flagImage    = flag.String("image", "", "When analyse a local file, it selects a particular image to analyse. Format: IMAGE_NAME[:TAG]")
	flagDebug    = flag.Bool("debug", false, "Enables debug messages")
	flagInsecure = flag.Bool("insecure", false, "Uses unencrypted connections when fetching images")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "analyzer [--debug] IMAGE\n")
	fmt.Fprintf(os.Stderr, "  Where IMAGE is\n")
	fmt.Fprintf(os.Stderr, "    [--image=IMAGE_NAME[:TAG]] FILEPATH\n")
	fmt.Fprintf(os.Stderr, "  or\n")
	fmt.Fprintf(os.Stderr, "    docker://[REGISTRYURL/]IMAGE_NAME[:TAG]\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		usage()
		return
	}

	if err := analyseDockerImage(args[0], *flagImage, *flagDebug, *flagInsecure); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
