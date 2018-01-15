package transcoder

import (
	"testing"
)

func TestTranscoder(t *testing.T) {
	trans := Transcoder{}

	profile := Profile{
		Transcode: Transcode{VideoCodec: "123",AudioCodec:"234"},
	}

	trans.Execute(profile)
}
