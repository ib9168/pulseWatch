package config

//Central place to read env vars
import (
	"log"
	"os"
)

var HTTPPort string
var DBPath string
var JWTSecret string
var GeminiAPI string

// Reads the required and optional environmental variables
func Load() {
	HTTPPort = os.Getenv("PULSEWATCH_HTTP_PORT")
	if HTTPPort == "" {
		HTTPPort = ":8080" //fallback default so can still run locally

	}
	//DB Path(optional)
	DBPath = os.Getenv("PULSEWATCH_DB_PATH")
	if DBPath == "" {
		DBPath = "./pulsewatch.db"
	}

	JWTSecret = os.Getenv("PULSEWATCH_JWT_SECRET")
	if JWTSecret == "" {
		log.Println("⚠️ WARNING: PULSEWATCH_JWT_SECRET not set — using insecure default (dev mode only)")
		//weak fallback for dev environment
		JWTSecret = "dev-secret-change-me"

	}
	GeminiAPI = os.Getenv("PULSEWATCH_GEMINI_API_KEY")
	if GeminiAPI == "" {
		log.Println("⚠️ WARNING: Gemini API key not set — AI summaries will NOT work until configured")
	}

}
