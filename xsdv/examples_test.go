package xsdv_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/terminalstatic/go-xsd-validate/libxml2"
	"github.com/terminalstatic/go-xsd-validate/xsdv"
)

// An example on how to use the package.
// In some situations, e.g. programatically looping over xml documents you might have to explicitly free the handler without defer. You prabably want to call xsdv.Init() and xsdv.Cleanup() only once after app start and before app end.
func Example() {
	l2 := libxml2.New()
	defer l2.Shutdown()
	xsdhandler, err := xsdv.NewXsdHandlerUrl("../examples/test1_split.xsd", libxml2.ParsErrDefault)
	if err != nil {
		panic(err)
	}
	defer xsdhandler.Free()

	xmlFile, err := os.Open("../examples/test1_fail2.xml")
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	inXml, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}

	xmlhandler, err := xsdv.NewXmlHandlerMem(inXml, libxml2.ParsErrDefault)
	if err != nil {
		panic(err)
	}
	defer xmlhandler.Free()

	err = xsdhandler.Validate(xmlhandler, libxml2.ValidErrDefault)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// Validation error:
	// Element 'shipto': This element is not expected. Expected is ( orderperson ).
}