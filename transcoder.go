package transcoder

import "github.com/easonlin404/transcoder/aws"

//type Command interface {
//	Execute() string
//}

type Transcoder struct {
	profile Profile
}

func (trans *Transcoder) Execute(p Profile) error {
	trans.profile = p

	trans.readInput()
	return nil
}

func (t *Transcoder) readInput() error {

	switch t.profile.Transcode.Input.Type {
	case "S3":
		//TODO: read S3 object
		aws.ListBuckets(t.profile.Transcode.Input.Bucket, t.profile.Transcode.Input.Source)
	case "File":
		//TODO: File
	default:
		panic("not supported")

	}

	return nil

}


