package utils

var (
	visaDevURLSecurity      = "https://apitestenv.vnforapps.com/api.security/v1/security"
	visaDevURLSession       = "https://apitestenv.vnforapps.com/api.ecommerce/v2/ecommerce/token/session/%s"
	visaDevURLAuthorization = "https://apitestenv.vnforapps.com/api.authorization/v3/authorization/ecommerce/%s"
	visaPrdURLSecurity      = "https://apiprod.vnforapps.com/api.security/v1/security"
	visaPrdURLSession       = "https://apiprod.vnforapps.com/api.ecommerce/v2/ecommerce/token/session/%s"
	visaPrdURLAuthorization = "https://apiprod.vnforapps.com/api.authorization/v3/authorization/ecommerce/%s"
)

// GetSecurityAPIURL return string.
func GetSecurityAPIURL(isDevelop bool) string {
	if isDevelop {
		return visaDevURLSecurity
	}
	return visaPrdURLSecurity
}

// GetSessionAPIURL return string.
func GetSessionAPIURL(isDevelop bool) string {
	if isDevelop {
		return visaDevURLSession
	}
	return visaPrdURLSession
}

// GetAuthorizationAPIURL return string.
func GetAuthorizationAPIURL(isDevelop bool) string {
	if isDevelop {
		return visaDevURLAuthorization
	}
	return visaPrdURLAuthorization
}
