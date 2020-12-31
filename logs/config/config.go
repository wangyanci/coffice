package config

type LogConfig struct {
	FileConfig
	ConsoleConfig
}

type FileConfig struct {
	Perm       string   `json:"perm,omitempty"`
	Level      int      `json:"level,omitempty"`
	Daily      bool     `json:"daily,omitempty"`
	Hourly     bool     `json:"hourly,omitempty"`
	Rotate     bool     `json:"rotate,omitempty"`
	MaxSize    int      `json:"maxsize,omitempty"`
	MaxDays    int64    `json:"maxdays,omitempty"`
	MaxHours   int64    `json:"maxhours,omitempty"`
	Filename   string   `json:"filename,omitempty"`
	MaxLines   int      `json:"maxlines,omitempty"`
	MaxFiles   int      `json:"maxfiles,omitempty"`
	RotatePerm string   `json:"rotateperm,omitempty"`
	Separate   []string `json:"separate,omitempty"`
}

type ConsoleConfig struct {
	Level    int  `json:"level,omitempty"`
	Colorful bool `json:"color,omitempty"` //this filed is useful only when system's terminal supports color
}
