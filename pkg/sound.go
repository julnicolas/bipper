package sound

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Player interface {
	Read(file string)
	Play()
}

type BeepPlayer struct {
	streamer beep.Streamer
	format beep.Format
}

func NewPlayer() Player {
	return &BeepPlayer{}
}

func (o *BeepPlayer) Read(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	var err error
	o.streamer, o.format, err = mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer o.streamer.Close()

	speaker.Init(o.format.SampleRate, o.format.SampleRate.N(time.Second/10))
}

func (o *BeepPlayer) Play() {
	speaker.Play(o.streamer)
}