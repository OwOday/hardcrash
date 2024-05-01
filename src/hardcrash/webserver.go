package hardcrash

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"

	"golang.org/x/net/html"
)

//TODO
//change tab like structure to a interconnected nodegraph, then use that same nodegraph code to make a browser, and use the native browser version when hardcrash launches that browser.

type Component struct {
	name         string
	doc          []byte
	guiCallbacks map[string]func(w http.ResponseWriter, r *http.Request)
}

type Webserver struct {
	doc         []byte
	components  []Component
	rendereddoc string
}

func loadGui() []byte {
	_, base, _, _ := runtime.Caller(0) // Relative to the runtime Dir
	//fmt.Println((base))
	dir := path.Join(path.Dir(base))
	// its better to use os specific path separator
	htmlDir := filepath.Join(dir, "gui.html")
	htmlBytes, err := os.ReadFile(htmlDir)
	if err != nil {
		fmt.Println(err)
		return []byte("There was an error running goglitch: " + err.Error()) //display the error as html for people new to the program //I dont think the program can even get here?
	}
	return htmlBytes
}

func insertComponents(n *html.Node, components [][]byte) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, v := range n.Attr {
			if v.Key == "class" {
				if v.Val == "tabs" {
					//n is the tabs node
					for _, v := range components {
						componenthtml, err := html.Parse(bytes.NewReader(v))
						if err != nil {
							fmt.Println("Error parsing HTML, is your user supplied component presenting invalid html?:", err)
							return
						}
						n.AppendChild(componenthtml)
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		insertComponents(c, components)
	}
}

func (webserver *Webserver) mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, webserver.rendereddoc)
}

func (webserver *Webserver) InitWebServer() {
	//webserver.doc = loadGui()
	htmldoc, err := html.Parse(bytes.NewReader(webserver.doc))
	if err != nil {
		fmt.Println("Error parsing HTML, is your user supplied component presenting invalid html?:", err)
		return
	}
	var guidocs [][]byte
	//loop over components, binding functions and building a slice of the docs
	for _, c := range webserver.components {
		guidocs = append(guidocs, c.doc)
		for e, g := range c.guiCallbacks {
			http.HandleFunc(e, g)
			//fmt.Println("whatever something that is not hi", g)
		}
		//fmt.Println("hi", c)
		//TODO SMILEY MAKE CALLBACK FUNCTIONS GO BRR HERE
	}
	insertComponents(htmldoc, guidocs)
	//	htmldoc.AppendChild()
	buf := new(bytes.Buffer)
	html.Render(buf, htmldoc)
	webserver.rendereddoc = buf.String()
	http.HandleFunc("/", webserver.mainHandler)
	go openBrowser(fmt.Sprintf("http://localhost:%d", 8080)) // Open browser in a goroutine
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		fmt.Println("Unsupported platform.")
	}

	if err != nil {
		fmt.Println("Error opening the web browser:", err)
	}
}
