package transcoder

type Profile struct {
	Packages  []Package `json:"Package"`
	Transcode Transcode `json:"Transcode"`
}

type Transcode struct {
	Input struct {
		Bucket string `json:"Bucket"`
		Source string `json:"Source"`
		Type   string `json:"Type"`
	} `json:"Input"`
	Output struct {
		Bucket string `json:"Bucket"`
		Dist   string `json:"Dist"`
		Type   string `json:"Type"`
	} `json:"Output"`
	Renditions []struct {
		Bitrate int `json:"Bitrate"`
		Height  int `json:"height"`
		Width   int `json:"width"`
	} `json:"Renditions"`
	AudioCodec string `json:"AudioCodec"`
	VideoCodec string `json:"VideoCodec"`
	FrameRate  string `json:"FrameRate"`
	TwoPass    bool   `json:"TwoPass"`
}

type Package struct {
	Drm    []string `json:"Drm"`
	Format string   `json:"Format"`
}
