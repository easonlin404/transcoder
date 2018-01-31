package transcoder

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTranscoder(t *testing.T) {
	trans := Transcoder{
		inputReader:mockInputReader{},
	}

	profile := Profile{
		Transcode: Transcode{Input: Input{
			Bucket:"bucket",
			Source:"input.mpeg",
			Type: "S3",
		},
			VideoCodec: "123", AudioCodec: "234"},
	}

	err:=trans.Execute(profile)

	assert.NoError(t,err)
}



type mockInputReader struct {
}


func (mockInputReader) Read(bucket, source string) ([]byte,error){
	//TODO:
	return nil,nil
}
