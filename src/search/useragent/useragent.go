package useragent

import (
	"math/rand"
	"time"
)

type userAgentWithHeaders struct {
	UserAgent       string
	SecCHUA         string
	SecCHUAMobile   string
	SecCHUAPlatform string
}

// UserAgents used when making requests and their corresponding Sec-Ch-Ua headers.
var userAgentArray = [...]userAgentWithHeaders{
	// Chrome 119.0.0, Windows
	{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
		`"Google Chrome";v="119", "Chromium";v="119", "Not=A?Brand";v="24"`,
		"?0",
		"\"Windows\"",
	},
	// Chrome 118.0.0, Windows
	{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
		`"Google Chrome";v="118", "Chromium";v="118", "Not=A?Brand";v="24"`,
		"?0",
		"\"Windows\"",
	},
	// Chrome 117.0.0, Windows
	{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
		`"Google Chrome";v="117", "Chromium";v="117", "Not=A?Brand";v="24"`,
		"?0",
		"\"Windows\"",
	},
}

func randomUserAgentStruct() userAgentWithHeaders {
	// WARNING: Will stop working after year 2262.
	randSrc := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSrc)
	return userAgentArray[randGen.Intn(len(userAgentArray))]
}

func RandomUserAgent() string {
	randomUA := randomUserAgentStruct()
	return randomUA.UserAgent
}

func RandomUserAgentWithHeaders() userAgentWithHeaders {
	return randomUserAgentStruct()
}
