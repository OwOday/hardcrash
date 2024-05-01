package hardcrash

import "go.bug.st/serial"

//import "fmt"
//everything needed to provide the user flows in gui.go
var mode = &serial.Mode{
	BaudRate: 115200,
}

type Hardcrash struct {
	toolComponents []Tool
	//glitcherComponents []Glitcher
	webserver Webserver
}

func New() *Hardcrash {
	return &Hardcrash{
		toolComponents: []Tool{},
		//glitcherComponents: []Glitcher{},
		//gui: Gui{
		//doc:        loadWebserver(),
		//	components: []Component{},
		//},
		webserver: Webserver{
			doc:        loadGui(),
			components: []Component{},
		},
	}
}

func (gog Hardcrash) Start() {
	//TODO bind components sanely
	gog.webserver.InitWebServer()
}

//register a tool *COMPONENT* that provides the tool interface (i.e. ender.go's Ender struct)
func (gog *Hardcrash) Registertool(tool Tool, name string) {
	gog.toolComponents = append(gog.toolComponents, tool)
	component := Component{
		name:         name,
		doc:          tool.GetGui(),
		guiCallbacks: tool.GetGuiFunctions(), //TODO SMILEY
	}
	gog.webserver.components = append(gog.webserver.components, component)
}

/*func (gog *Hardcrash) Registerglitcher(glitcher Glitcher, name string) {
	//gog.glitcherComponents = append(gog.glitcherComponents, glitcher)
	component := Component{
		name:         name,
		doc:          glitcher.GetGui(),
		guiCallbacks: glitcher.GetGuiFunctions(), //TODO SMILEY
	}
	gog.webserver.components = append(gog.webserver.components, component)
}*/
