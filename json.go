package geogoth

// // FileFeatures ...
// type FileFeatures struct {
// 	Type     string        `json:"type"`
// 	Features []FileFeature `json:"features"`
// }

// // FileFeature ...
// type FileFeature struct {
// 	Type       string                 `json:"type"`
// 	Properties map[string]interface{} `json:"properties"`
// 	ID         string                 `json:"id"`

// 	Geometry struct {
// 		Type        string      `json:"type"`
// 		Coordinates interface{} `json:"coordinates"`
// 	} `json:"geometry"`
// }

// JSONFromFile ...
// func JSONFromFile(file string) (geo Features) {

// 	bytes, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	json.Unmarshal(bytes, &geo)

// 	return geo
// }
