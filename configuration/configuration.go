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
	Methods       []string `yaml:"methods"`
	Credentials   bool     `yaml:"credentials"`
	Origins       []string `yaml:"origins"`
	AllowHeaders  []string `yaml:"allowHeaders"`
	JwtCookieName string   `yaml:"jwtName"`
}

// EndpointTarget : Target proxy for the endpoint
type EndpointTarget struct {
	HostTarget string `yaml:"hostTarget"`
	URLTarget  string `yaml:"urlTarget"`
	Method     string `yaml:"method"`

	// Additional or Helper
	// Used to increase request performance
	IsStar bool `yaml:"isStar"`
}

// EndpointJWT
type EndpointJWT struct {
	Alghorithm string `yaml:"alghorithm"`
	Key        string `yaml:"key"`
	Signature  string `yaml:"signature"`
}

// Endpoint : Endpoint struct, for proxy server
type Endpoint struct {
	Endpoint   string           `yaml:"endpoint"`
	Method     string           `yaml:"method"`
	Sequential bool             `yaml:"sequential"`
	DeepMerge  bool             `yaml:"deepMerge"`
	Merge      bool             `yaml:"merge"`
	Targets    []EndpointTarget `yaml:"targets"`

	// Rate Limiter
	RateLimiter  int `yam:"rateLimiter"`
	RateDuration int `yaml:"rateDuration"`

	// Additional or Helper
	Segments    []string    `yaml:"segments"`
	ParamsIndex []int       `yaml:"paramsIndex"`
	Jwt         EndpointJWT `yaml:"jwt"`
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
