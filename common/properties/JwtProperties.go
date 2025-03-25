package properties

type JwtProperties struct {
	AdminSecretKey string
	AdminTtl       int64
	AdminTokenName string
	UserSecretKey  string
	UserTtl        int64
	UserTokenName  string
}
