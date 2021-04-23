package main

import (
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/hostname", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, getHostname())
	})

	http.HandleFunc("/ps", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, getAllProcesses())
	})

	http.HandleFunc("/free", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, showFreeMemory())
	})

	http.HandleFunc("/ps-httpd", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, showHttpd())
	})

	http.HandleFunc("/memhog", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, showMemhog())
	})

	http.HandleFunc("/ls-usr-bin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, showLsUsrBin())
	})

	http.HandleFunc("/ipaddress", func(w http.ResponseWriter, r *http.Request) {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			os.Stderr.WriteString("Oops: " + err.Error() + "\n")
			os.Exit(1)
		}

		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					os.Stdout.WriteString(ipnet.IP.String() + "\n")
					fmt.Fprintf(w, ipnet.IP.String())
				}
			}
		}
	})

	http.HandleFunc("/ps-mongo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, showMongoProcess())
	})

	http.HandleFunc("/ps-maria", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, showMariaProcess())
	})

	http.HandleFunc("/ps-concerto", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, showConcertoProcess())
	})

	http.HandleFunc("/ps-haproxy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, showHaProxyProcess())
	})

	http.HandleFunc("/df-h", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, showDiskFree())
	})

	log.Fatal(http.ListenAndServe(":8082", nil))

}

func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return name
}

func showFreeMemory() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "free"
	cmdArgs := []string{"-hg"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error calling "+cmdName, err)
		os.Exit(1)
	}
	myResult := string(cmdOut)
	firstHundred := myResult[:100]
	fmt.Println("Successfully calling "+cmdName, firstHundred)

	return myResult
}

func getAllProcesses() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "ps"
	cmdArgs := []string{"-elf"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error calling "+cmdName, err)
		os.Exit(1)
	}
	myResult := string(cmdOut)
	firstHundred := myResult[:100]
	fmt.Println("Successfully calling "+cmdName, firstHundred)

	return myResult
}

func showHttpd() string {

	// using the function
	mydir, errdir := os.Getwd()
	if errdir != nil {
		fmt.Println(errdir)
	}
	fmt.Println(mydir)

	var (
		cmdOut []byte
		err    error
	)
	cmdName := "/home/ea856a/monitoring-agent/go-servers/ps-efH-grep-httpd.sh" // ps -efH | grep httpd
	cmdArgs := []string{}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error calling "+cmdName, err)

	}
	myResult := string(cmdOut)
	firstHundred := myResult[:100]
	fmt.Println("Successfully calling "+cmdName, firstHundred)

	return myResult
}

func showMemhog() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "ps" // ps aux --sort -rss
	cmdArgs := []string{"aux", "--sort", "-rss"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error calling "+cmdName, err)

	}
	myResult := string(cmdOut)
	firstHundred := myResult[:100]
	fmt.Println("Successfully calling "+cmdName, firstHundred)

	return myResult
}

func showLsUsrBin() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "/home/ea856a/monitoring-agent/go-servers/ls-usr-bin.sh" // ps aux --sort -rss
	cmdArgs := []string{}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error calling "+cmdName, err)

	}
	myResult := string(cmdOut)
	firstHundred := myResult[:100]
	fmt.Println("Successfully calling "+cmdName, firstHundred)

	return myResult
}

func showMongoProcess() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "/home/ea856a/monitoring-agent/go-servers/ps-efH-grep-mongo.sh" // ps -efH | grep mongo
	cmdArgs := []string{}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error calling "+cmdName, err)

	}
	myResult := string(cmdOut)
	firstHundred := myResult[:100]
	fmt.Println("Successfully calling "+cmdName, firstHundred)

	return myResult
}

func showMariaProcess() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "/home/ea856a/monitoring-agent/go-servers/ps-efH-grep-maria.sh" // ps -efH | grep maria
	cmdArgs := []string{}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error calling "+cmdName, err)

	}
	myResult := string(cmdOut)
	firstHundred := myResult[:100]
	fmt.Println("Successfully calling "+cmdName, firstHundred)

	return myResult
}

func showConcertoProcess() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "/home/ea856a/monitoring-agent/go-servers/ps-efH-grep-concerto.sh" // ps -efH | grep maria
	cmdArgs := []string{}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error calling "+cmdName, err)

	}
	myResult := string(cmdOut)
	firstHundred := myResult[:100]
	fmt.Println("Successfully calling "+cmdName, firstHundred)

	return myResult
}

func showHaProxyProcess() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "/home/ea856a/monitoring-agent/go-servers/ps-efH-grep-haproxy.sh" // ps -efH | grep maria
	cmdArgs := []string{}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error calling "+cmdName, err)

	}
	myResult := string(cmdOut)
	firstHundred := myResult[:100]
	fmt.Println("Successfully calling "+cmdName, firstHundred)

	return myResult
}

func showDiskFree() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "df"
	cmdArgs := []string{"-h"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error calling "+cmdName, err)
		os.Exit(1)
	}
	myResult := string(cmdOut)
	firstHundred := myResult[:100]
	fmt.Println("Successfully calling "+cmdName, firstHundred)

	return myResult
}
