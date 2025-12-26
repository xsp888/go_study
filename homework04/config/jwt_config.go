package config

type JWTConfig struct {
	JWTSecret string
	JWTExpire int // 过期时间（小时）
}

var AppConfig = &JWTConfig{
	JWTSecret: "xsp",
	JWTExpire: 24, // 24小时
}
