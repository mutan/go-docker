package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	//response, _ := http.Get("https://iss.moex.com/iss/engines/stock/markets/bonds/boards/TQCB/securities.xml")
	//bytes, _ := ioutil.ReadAll(response.Body)

	//fmt.Println(string(bytes))

	data := []byte(`<document>
    <data id="securities">
        <rows>
            <row SECID="RU000A0JNYN1" BOARDID="TQCB" SHORTNAME="МГор48-об" />
        </rows>
    </data>
</document>`)

	var res Document
	err := xml.Unmarshal(data, &res)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", res)
	fmt.Printf("%#v\n", res)
}

type Document struct {
	XMLName xml.Name `xml:"document"`
	Data    []Data   `xml:"data"`
}

type Data struct {
	Id   string `xml:"id,attr"`
	Rows []Rows `xml:"rows"`
}

type Rows struct {
	Row []Row `xml:"row"`
}

type Row struct {
	Secid     string `xml:"SECID,attr"`
	BOARDID   string `xml:"BOARDID,attr"`
	SHORTNAME string `xml:"SHORTNAME,attr"`
}
