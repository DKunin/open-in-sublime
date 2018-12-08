package settings

type Settings struct {
	port string
	editor string
}

func (s *Settings) SetEditor(editor string) {
	s.editor = editor
}

func (s *Settings) SetPort(port string) {
	s.port = port
}

func (s *Settings) GetPort() string {
	return s.port
}

func (s *Settings) GetEditor() string {
	return s.editor
}

func (s *Settings) GetSettings() string {
	return "{\"port\": "+ s.port + ", \"editor\": \"" + s.editor + "\"}"
}