package models

// templateDate creats data for the template package
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FLoatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
