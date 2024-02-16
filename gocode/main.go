package main

import (
	"fmt"

	"os"
	"strings"

	"io/ioutil"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

//This md2html executable must run in the main repo, not inside gocode.
//Ex.: ./gocode/md2html

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	mdFileNames := []string{"about", "preparation", "company",
		"migration", "distribution", "competition",
		"finance", "european", "transport",
		"insurance", "environmental", "energy",
		"realestate", "labour", "it",
		"dispute", "simonas"}

	slotContentFileNames := []string{"banner01", "banner02", "banner03"}

	langs := []string{"de", "en", "es", "fr", "it", "lt", "no", "pl", "ru"}

	for _, lang := range langs {
		// Read in the template HTML file
		template, err := os.ReadFile("lang/" + lang + "/index_template.html")
		check(err)
		output := string(template)

		for _, name := range mdFileNames {

			inputFile := "lang/" + lang + "/md/" + name + ".md"

			fmt.Printf("Processing file %s.\n", inputFile)

			// Read in the Markdown file
			md, err := os.ReadFile(inputFile)
			check(err)

			// Convert the Markdown to HTML
			//md = []byte(md)
			// Create a parser
			extensions := parser.CommonExtensions | parser.Footnotes
			parser := parser.NewWithExtensions(extensions)

			html := markdown.ToHTML(md, parser, nil)

			// Replace the {src} placeholder with the HTML
			output = strings.Replace(output, "{src-"+name+"}", string(html), -1)

		}

		for _, name := range slotContentFileNames {

			inputFile := "lang/" + lang + "/txt/" + name + ".txt"

			fmt.Printf("Processing file %s.\n", inputFile)

			// Read in the text file
			slotContent, err := ioutil.ReadFile(inputFile)
			check(err)

			// Replace the {src} placeholder with the HTML
			output = strings.Replace(output, "{src-"+name+"}", string(slotContent), -1)

		}

		// Write the output to the specified file
		err = os.WriteFile("lang/"+lang+"/index.html", []byte(output), 0644)
		check(err)
	}
	fmt.Printf("Done.\n")

}
