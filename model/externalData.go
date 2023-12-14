package model

import "os"

const POOLLIMIT = 40
const MAXROUTINE = POOLLIMIT

type TokenNames struct {
	Sources  []string
	Channels []string
}

func GatherTokens() TokenNames {
	return TokenNames{
		Sources:  []string{"AB", "SS", "SM"},
		Channels: []string{"Amazon", "Walmart"},
	}
}

// Assumption token name index -> token
func RequestTokens() ([]string, []string) {
	return []string{os.Getenv("FLX_AB_TOKEN"), os.Getenv("FLX_SS_TOKEN"), os.Getenv("FLX_SM_TOKEN")}, []string{os.Getenv("FLX_AZ_TOKEN"), os.Getenv("FLX_WALMART_TOKEN")}
}

func RequestAccToken() string {
	return os.Getenv("FLX_API_TOKEN")
}
