package config

type EmailParam struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	EmailServer string `json:"email_server"`
	Port        string `json:"port"`
}
