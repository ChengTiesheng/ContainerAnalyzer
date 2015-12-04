package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

var (
	flagImage    = flag.String("image", "", "When converting a local file, it selects a particular image to convert. Format: IMAGE_NAME[:TAG]")
	flagDebug    = flag.Bool("debug", false, "Enables debug messages")
	flagInsecure = flag.Bool("insecure", false, "Uses unencrypted connections when fetching images")
)

func analyseDockerImage(arg string, flagImage string, flagDebug bool, flagInsecure bool) error {
	if flagDebug {
		//TODO
	}

	// try to convert a local file
	u, err := url.Parse(arg)
	if err != nil {
		return fmt.Errorf("error parsing argument: %v", err)
	}
	if u.Scheme == "docker" {
		if flagImage != "" {
			return fmt.Errorf("flag --image works only with files.")
		}
		dockerURL := strings.TrimPrefix(arg, "docker://")

		indexServer := GetIndexName(dockerURL)

		var username, password string
		username, password, err = GetDockercfgAuth(indexServer)
		if err != nil {
			return fmt.Errorf("error reading .dockercfg file: %v", err)
		}

		err = Analyse(dockerURL, username, password, flagInsecure)
	} else {
		err = AnalyseFile(flagImage, arg)
	}
	if err != nil {
		return fmt.Errorf("conversion error: %v", err)
	}

	return nil
}

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
