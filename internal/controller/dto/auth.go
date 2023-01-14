package dto

type CreateCertificate struct {
	PublicKey []byte `json:"publicKey"`
	Time      string `json:"time"`
}

func ConvertTime(time string) int {
	timeTable := map[string]int{
		"5m":  5,
		"10m": 10,
		"30m": 30,
		"60m": 60,
		"3h":  180,
		"6h":  360,
		"12h": 720,
	}

	return timeTable[time]
}
