package transcoder


type Transcoder struct {
	profile Profile
	inputReader InputReader
}


type InputReader interface {
 Read(bucket, source string) ([]byte,error)
}

func (trans *Transcoder) Execute(p Profile) error {
	trans.profile = p

	 _,err:=trans.readInput()
	return err
}

func (t *Transcoder) readInput()([]byte, error) {
	bucket:=t.profile.Transcode.Input.Bucket
	source:=t.profile.Transcode.Input.Source
	return t.inputReader.Read(bucket,source)

}


