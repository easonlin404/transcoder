package transcoder

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
		//TODO: S3
	case "File":
		//TODO: File
	default:
		panic("not supported")

	}

	return nil

}
