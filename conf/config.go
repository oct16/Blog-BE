package conf

// Config ..
var Config = make(map[string]interface{})

func init() {
	var isProd = Env == "production"

	Config["Port"] = "3015"
	if isProd {
		Config["Port"] = "3016"
	}
}
