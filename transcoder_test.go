package transcoder

import (
	"testing"
)

func TestTranscoderFromS3(t *testing.T) {
	trans := Transcoder{}

	profile := Profile{
		Transcode: Transcode{Input: Input{
			Bucket:"xxx",
			Type: "S3",
		},
			VideoCodec: "123", AudioCodec: "234"},
	}

	trans.Execute(profile)
}
