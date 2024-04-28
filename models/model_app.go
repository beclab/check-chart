package models

type AppInfo struct {
	ID int

	Name  string
	Owner string
	//OwnerId string
}

// Chart represents the structure of the Chart.yaml file
type Chart struct {
	APIVersion string `yaml:"apiVersion"`
	Name       string `yaml:"name"`
	Version    string `yaml:"version"`
}

// AppConfiguration represents the structure of the TerminusManifest.yaml file
type AppConfiguration struct {
	ConfigVersion string      `yaml:"terminusManifest.version" json:"terminusManifest.version"`
	Metadata      AppMetaData `yaml:"metadata" json:"metadata"`
	Spec          AppSpec     `yaml:"spec" json:"spec"`
}

type AppMetaData struct {
	Name        string   `yaml:"name" json:"name"`
	Icon        string   `yaml:"icon" json:"icon"`
	Description string   `yaml:"description" json:"description"`
	AppID       string   `yaml:"appid" json:"appid"`
	Title       string   `yaml:"title" json:"title"`
	Version     string   `yaml:"version" json:"version"`
	Categories  []string `yaml:"categories" json:"categories"`
	Rating      float32  `yaml:"rating" json:"rating"`
	Target      string   `yaml:"target" json:"target"`
}

type AppSpec struct {
	VersionName        string        `yaml:"versionName" json:"versionName"`
	FullDescription    string        `yaml:"fullDescription" json:"fullDescription"`
	UpgradeDescription string        `yaml:"upgradeDescription" json:"upgradeDescription"`
	PromoteImage       []string      `yaml:"promoteImage" json:"promoteImage"`
	PromoteVideo       string        `yaml:"promoteVideo" json:"promoteVideo"`
	SubCategory        string        `yaml:"subCategory" json:"subCategory"`
	Language           []string      `yaml:"language" json:"language"`
	Developer          string        `yaml:"developer" json:"developer"`
	RequiredMemory     string        `yaml:"requiredMemory" json:"requiredMemory"`
	RequiredDisk       string        `yaml:"requiredDisk" json:"requiredDisk"`
	SupportClient      SupportClient `yaml:"supportClient" json:"supportClient"`
	RequiredGPU        string        `yaml:"requiredGpu" json:"requiredGpu,omitempty"`
	RequiredCPU        string        `yaml:"requiredCpu" json:"requiredCpu"`

	Submitter string       `yaml:"submitter" json:"submitter"`
	Doc       string       `yaml:"doc" json:"doc"`
	Website   string       `yaml:"website" json:"website"`
	License   []TextAndURL `yaml:"license" json:"license"`
	Legal     []TextAndURL `yaml:"legal" json:"legal"`
}

type TextAndURL struct {
	Text string `yaml:"text" json:"text" bson:"text"`
	URL  string `yaml:"url" json:"url" bson:"url"`
}

type SupportClient struct {
	Edge    string `yaml:"edge" json:"edge" bson:"edge"`
	Android string `yaml:"android" json:"android" bson:"android"`
	Ios     string `yaml:"ios" json:"ios" bson:"ios"`
	Windows string `yaml:"windows" json:"windows" bson:"windows"`
	Mac     string `yaml:"mac" json:"mac" bson:"mac"`
	Linux   string `yaml:"linux" json:"linux" bson:"linux"`
}
