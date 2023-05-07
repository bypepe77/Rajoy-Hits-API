package structs

import "fmt"

// atributos
type Streamer struct {
	Name    string
	Channel string
	Viewers int
}

// constructora
func NewStreamer(name, channel string, viewers int) *Streamer {
	return &Streamer{
		Name:    name,
		Channel: channel,
		Viewers: viewers,
	}
}

// setter
func (s *Streamer) SetName(name string) {
	s.Name = name
}

// getter
func (s *Streamer) GetName() string {
	return s.Name
}

// setter
func (s *Streamer) SetChannel(channel string) {
	s.Channel = channel
}

// getter
func (s *Streamer) GetChannel() string {
	return s.Channel
}

// setter
func (s *Streamer) SetViewers(viewers int) {
	s.Viewers = viewers
}

// getter
func (s *Streamer) GetViewers() int {
	return s.Viewers
}

func (s *Streamer) ToString() string {
	//`Mi nombre es ${nombre}, tengo ${edad} años y mido ${altura} metros.`
	return fmt.Sprintf("Streamer:{Name:  %s, Channel: %s, Viewers: %d}", s.GetName(), s.GetChannel(), s.GetViewers())
}
