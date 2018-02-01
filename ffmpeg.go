package transcoder

import (
	"html/template"
)


type FFmpeg struct {
	
}

// Excute FFmpeg
func (f *FFmpeg) Execute(p Profile) error {


}


var ffmpegTemplate = template.Must(template.New("").Parse("ffmpeg -y -i")