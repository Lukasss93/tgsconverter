package main

import "fmt"
import "io/ioutil"
import "path/filepath"
import "os"
import "strings"
import "libtgsconverter"

func main() {
    // check if the user provided 3 arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: tgsconverter path/to/tgs format (apng, gif, png, webp)")
		os.Exit(1)
	}
	
	// get the path to the output file
	outPath := os.Args[2]
	
	// get the format of the output file
	extension := strings.TrimPrefix(filepath.Ext(outPath), ".")
	
	// check if the format is valid
	if !libtgsconverter.SupportsExtension(extension) {
		fmt.Println("Unsupported extension: " + extension)
		os.Exit(2)
	}
	
	// get the path to the tgs file
	inPath := os.Args[1]
	
	// convert the tgs file to the output format
	opt := libtgsconverter.NewConverterOptions()
	opt.SetExtension(extension)
	ret, err := libtgsconverter.ImportFromFile(inPath, opt)
	if err != nil {
		fmt.Println("Error in tgsconverter.ImportFromFile:" + err.Error())
		os.Exit(3)
	}
	
	// ensure the output directory exists
	dirName := filepath.Dir(outPath)
      if _, serr := os.Stat(dirName); serr != nil {
        merr := os.MkdirAll(dirName, os.ModePerm)
        if merr != nil {
            fmt.Println("Error creating folder:" + merr.Error())
            os.Exit(4)
        }
      }
	
	// write the output file
	err = ioutil.WriteFile(outPath, ret, 0666)
	if err != nil {
        fmt.Println("Error saving file:" + err.Error())
        os.Exit(5)
    }
}