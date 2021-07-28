package configuration

// Configuration :
type Configuration struct {
	Name       string `yaml:"name"`
	version    string `yaml:"version"`
	Port       string `yaml:"port"`
	Address    string `yaml:"address"`
	Production bool   `yaml:"production"`

	Endpoints []Endpoint `yaml:"endpoints"`
	Header    Header     `yaml:"header"`
	Statics   []Static   `yaml:"statics"`
	Webs      []Web      `yaml:"webs"`
}

// Header : Configuration
type Header struct {
	Methods      []string `yaml:"methods"`
	Credentials  string   `yaml:"credentials"`
	Origins      []string `yaml:"origins"`
	AllowHeaders []string `yaml:"allowHeaders"`
}

// Endpoint : Endpoint struct, for proxy server
type Endpoint struct {
	Endpoint   string `yaml:"endpoint"`
	Method     string `yaml:"method"`
	HostTarget string `yaml:"hostTarget"`
	URLTarget  string `yaml:"urlTarget"`
	Jwt        string `yaml:"jwt"`
}

// Static : Passive Folder
type Static struct {
	Alias string `yaml:"alias"`
	Path  string `yaml:"path"`
}

// Web :
type Web struct {
	Path        string   `yaml:"path"`
	Windowspath string   `yaml:"windowsPath"`
	Alias       string   `yaml:"alias"`
	Extensions  []string `yaml:"extensions"` // What file extension supported by the gateway
}

// Tuple :
type Tuple struct {
	Key   string
	Value string
}
