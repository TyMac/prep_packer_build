package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"text/template"
)

func main() {
	// variables
	if runtime.GOOS == "windows" {
		fmt.Println("Windows Build Server Detected")
	} else if runtime.GOOS == "darwin" {
		fmt.Println("Macos Build Server Detected")
	} else {
		fmt.Println("Linux Build Server Detected")
	}
	// platform := flag.String("platform", "windows", "Directory to write autounattend.xml to")
	httpdir := flag.String("httpdir", "C:\\Users\\Public\\Documents\\", "Directory to write autounattend.xml to")
	templatedir := flag.String("templatedir", ".\\templates\\", "Directory to find autounattend.xml.tmpl template files in")
	// variables to be parsed in to the autounattend template
	vars := make(map[string]interface{})
	vars["PassWord"] = flag.String("password", "flapjacks123", "The Admin default password")
	vars["AesKey"] = flag.String("aeskey", "kekekekekekekekekekekekeke=", "An AES256 key to decrypt sensitive data")
	vars["StagingDirectory"] = flag.String("stagedir", "C:\\Windows\\Temp", "The directory on the VM for files used in the build")
	vars["ProxyAddress"] = flag.String("proxy", "1.1.1.1:8080", "The proxy IP")
	vars["InstallUpdates"] = flag.String("updates", "false", "If security updates should be applied")
	flag.Parse()
	// parse the template
	tmpl, _ := template.ParseFiles(*templatedir + "autounattend.xml.tmpl")
	// create a new file
	file, err := os.Create(*httpdir + "autounattend.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// apply the template to the vars map and write the result to file.
	tmpl.Execute(file, vars)
}
