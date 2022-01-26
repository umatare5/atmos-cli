package domain

// Global: Name Flags
const (
	ServerFlagName string = "server"
)

// Global: Usage Flags
const (
	ServerFlagUsage string = "Set FQDN for ATMOS Platform"
)

// Global: Alias Flags
var (
	ServerFlagAlias []string = []string{"s"}
)

// Global: Value Flags
const (
	ServerFlagValue string = "http://localhost:8080"
)

// Name Flags
const (
	UsernameFlagName     string = "username"
	PasswordFlagName     string = "password"
	SaveProfileFlagName  string = "save-profile"
	AccessTokenFlagName  string = "access-token"
	LimitFlagName        string = "limit"
	CursorFlagName       string = "cursor"
	UserIDFlagName       string = "user-id"
	StatisticsFlagName   string = "statistics"
	DivelogIDFlagName    string = "divelog-id"
	PrettyFormatFlagName string = "pretty-format"
)

// Usage Flags
const (
	UsernameFlagUsage     string = "Set email as username"
	PasswordFlagUsage     string = "Set password"
	SaveProfileFlagUsage  string = "Save token to ~/.atmosrc"
	AccessTokenFlagUsage  string = "Access token to use API"
	LimitFlagUsage        string = "Limit data when querying"
	CursorFlagUsage       string = "Set a point to start querying"
	UserIDFlagUsage       string = "Set user ID"
	StatisticsFlagUsage   string = "Show user statistics"
	DivelogIDFlagUsage    string = "Set divelog ID"
	PrettyFormatFlagUsage string = "Show response as pretty format"
)

// Value Flags
const (
	SaveProfileFlagDefaultValue  bool   = false
	LimitFlagDefaultValue        int    = 10
	CursorFlagDefaultValue       string = ""
	UserIDFlagDefaultValue       string = ""
	DivelogIDFlagDefaultValue    string = ""
	StatisticsFlagDefaultValue   bool   = false
	PrettyFormatFlagDefaultValue bool   = false
)

// Required Flags
const (
	UsernameFlagRequired    bool = true
	PasswordFlagRequired    bool = true
	AccessTokenFlagRequired bool = true
)

// Flag aliases
var (
	UsernameFlagAliases     = []string{"u"}
	PasswordFlagAliases     = []string{"p"}
	SaveProfileFlagAliases  = []string{"s"}
	AccessTokenFlagAliases  = []string{"t"}
	LimitFlagAliases        = []string{"l"}
	CursorFlagAliases       = []string{"c"}
	UserIDFlagAliases       = []string{"i", "uid"}
	StatisticsFlagAliases   = []string{"S", "stats"}
	DivelogIDFlagAliases    = []string{"L", "logid"}
	PrettyFormatFlagAliases = []string{"P", "pretty"}
)

// Flag envvars
var (
	ServerFlagEnvVars      = []string{"ATMOS_API"}
	AccessTokenFlagEnvVars = []string{"ATMOS_TOKEN"}
	UserIDFlagEnvVars      = []string{"ATMOS_USERID"}
)
