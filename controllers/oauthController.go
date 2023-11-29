package controllers

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8000/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),     // 使用環境變數名稱
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"), // 使用環境變數名稱
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// // func GoogleOAuth()gin.HandlerFunc {
// // 	return func(c *gin.Context){
// // 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// // 		defer cancel()
// // 		ctx.Done()

// // }

// // func OauthGoogleCallback() gin.HandlerFunc {
// // 	return func(c *gin.Context){
// // 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// // 		defer cancel()
// // 		ctx.Done()

// // }

// // func generateStateOauthCookie(w http.ResponseWriter) string {
// // 	var expiration = time.Now().Add(20 * time.Minute)

// // 	b := make([]byte, 16)
// // 	rand.Read(b)
// // 	state := base64.URLEncoding.EncodeToString(b)
// // 	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
// // 	http.SetCookie(w, &cookie)

// // 	return state
// // }

// // func getUserDataFromGoogle(code string) ([]byte, error) {
// // 	// Use code to get token and get user info from Google.

// // 	token, err := googleOauthConfig.Exchange(context.Background(), code)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
// // 	}
// // 	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
// 	}
// 	defer response.Body.Close()
// 	contents, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed read response: %s", err.Error())
// 	}
// 	return contents, nil
// }
