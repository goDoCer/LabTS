package main

import "LabTS/labts/website"

//LabTSURL is the base url
const LabTSURL = "https://teaching.doc.ic.ac.uk/labts"

func main() {
	website.SetAuth("jas")
}
