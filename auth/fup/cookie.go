package fup

import (
	"github.com/gin-gonic/gin"
	"github.com/wernerdweight/api-auth-go/auth/config"
	"github.com/wernerdweight/api-auth-go/auth/constants"
	"github.com/wernerdweight/api-auth-go/auth/contract"
	"log"
	"net/http"
)

// CookieFUPChecker is an implementation of the FUPCheckerInterface for the URL path-based access model
type CookieFUPChecker struct {
	CookieName string
}

func (ch CookieFUPChecker) Check(scope *contract.FUPScope, c *gin.Context, key string) contract.FUPScopeLimits {
	name := "api-auth-go-fup"
	if "" != ch.CookieName {
		name = ch.CookieName
	}
	cookie, err := c.Cookie(name)
	log.Printf("cookie: %s", cookie)
	if nil != err && http.ErrNoCookie != err {
		return contract.FUPScopeLimits{
			Accessible: constants.ScopeAccessibilityForbidden,
			Error:      contract.NewInternalError(contract.InvalidFUPCookie, nil),
		}
	}
	if nil == scope || "" == cookie {
		// no limitations by default
		return contract.FUPScopeLimits{Accessible: constants.ScopeAccessibilityUnlimited}
	}
	if !scope.HasLimit(constants.FUPCookieKey) {
		return contract.FUPScopeLimits{Accessible: constants.ScopeAccessibilityUnlimited}
	}
	if !config.ProviderInstance.IsCacheEnabled() {
		return contract.FUPScopeLimits{
			Error: contract.NewInternalError(contract.FUPCacheDisabled, nil),
		}
	}
	cacheDriver := config.ProviderInstance.GetCacheDriver()
	cookieLimits, scopeLimits := checkLimits(scope, key, cookie, constants.FUPCookieKey, cacheDriver)
	if nil != scopeLimits {
		return *scopeLimits
	}
	return contract.FUPScopeLimits{
		Accessible: constants.ScopeAccessibilityAccessible,
		Limits:     cookieLimits,
		Error:      nil,
	}
}
