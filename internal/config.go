package internal

type Config struct {
	Version     string `url:"v"`
	Method      string `url:"method"`
	AccessToken string `url:"access_token"`
	AppKey      string `url:"app_key"`
	SignMethod  string `url:"sign_method"`
	Format      string `url:"format"`
	Timestamp   string `url:"timestamp"`
	Sign        string `url:"sign"`
	ParamJson   string `url:"param_json"`
}
