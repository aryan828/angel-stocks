package main

type BaseResponse struct {
	Status    bool
	Message   string
	ErrorCode string
}

type ConnectionResponse struct {
	BaseResponse
	Data struct {
		JwtToken     string
		RefreshToken string
		FeedToken    string
	}
}

type ProfileResponse struct {
	BaseResponse
	Data struct {
		ClientCode    string
		Name          string
		Email         string
		Mobileno      string
		Exchanges     []string
		Products      []string
		LastLoginTime string
		BrokerID      string
	}
}
