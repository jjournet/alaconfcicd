package alacritty

import (
	"errors"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	yml *yaml.Node
	loc string
}

func GetColorTheme(config *Config) string {
	elements := config.yml.Content[0].Content
	for _, elt := range elements {
		if elt.Value == "colors" {
			return elt.HeadComment
		}
	}
	return ""
}

func ChangeColorTheme(config *Config, theme string) {
	// check if file exists with extension yml or yaml
	path1 := filepath.Join(config.loc, "colors", theme+".yml")
	path2 := filepath.Join(config.loc, "colors", theme+".yaml")
	file := ""
	if _, err := os.Stat(path1); err == nil {
		file = path1
	} else if errors.Is(err, os.ErrNotExist) {
		if _, err := os.Stat(path2); err == nil {
			file = path2
		}
	}
	if file == "" {
		log.Fatal("theme not found")
	}
	// read content of new color theme from file
	var colornode yaml.Node
	content2, err2 := os.ReadFile(file)
	if err2 != nil {
		log.Fatal(err2)
	}
	err3 := yaml.Unmarshal(content2, &colornode)
	if err3 != nil {
		log.Fatal(err3)
	}
	elements := config.yml.Content[0].Content
	idx := 0
	for i, elt := range elements {
		if elt.Value == "colors" {
			idx = i
			break
		}
	}
	// spew.Dump(config.yml)
	elements[idx].HeadComment = theme
	elements[idx+1] = colornode.Content[0].Content[1]
}

func SaveConfig(config *Config) {
	// write new config file
	f, err5 := os.Create(filepath.Join(config.loc, "alacritty.yml"))
	if err5 != nil {
		log.Fatal(err5)
	}
	defer f.Close()
	results, _ := yaml.Marshal(config.yml.Content[0])
	// write to file
	_, err6 := f.Write(results)
	if err6 != nil {
		log.Fatal(err6)
	}
}

func NewConfig(location string) Config {
	// if location empty, use default location
	if location == "" {
		location = "/.config/alacritty/alacritty.yml"
		usr, _ := user.Current()
		dir := usr.HomeDir
		location = filepath.Join(dir, location)
	}
	newConf := Config{loc: filepath.Dir(location)}
	// read yaml config file
	content, err := os.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	var node yaml.Node
	newConf.yml = &node
	err2 := yaml.Unmarshal(content, &node)
	if err2 != nil {
		log.Fatal(err2)
	}

	return newConf
}

// type Config struct {
// 	Import []string `yaml:"import"`
// 	Env    struct {
// 		TERM string `yaml:"TERM"`
// 	} `yaml:"env"`
// 	Window struct {
// 		Dimensions struct {
// 			Columns int `yaml:"columns"`
// 			Lines   int `yaml:"lines"`
// 		} `yaml:"dimensions"`
// 		Position struct {
// 			X int `yaml:"x"`
// 			Y int `yaml:"y"`
// 		} `yaml:"position"`
// 		Padding struct {
// 			X int `yaml:"x"`
// 			Y int `yaml:"y"`
// 		} `yaml:"padding"`
// 		DynamicPadding bool    `yaml:"dynamic_padding"`
// 		Decorations    string  `yaml:"decorations"`
// 		Opacity        float64 `yaml:"opacity"`
// 		StartupMode    string  `yaml:"startup_mode"`
// 		Title          string  `yaml:"title"`
// 		DynamicTitle   bool    `yaml:"dynamic_title"`
// 		Class          struct {
// 			Instance string `yaml:"instance"`
// 			General  string `yaml:"general"`
// 		} `yaml:"class"`
// 		DecorationsThemeVariant string `yaml:"decorations_theme_variant"`
// 	} `yaml:"window"`
// 	Scrolling struct {
// 		History    int `yaml:"history"`
// 		Multiplier int `yaml:"multiplier"`
// 	} `yaml:"scrolling"`
// 	Font struct {
// 		Normal struct {
// 			Family string `yaml:"family"`
// 			Style  string `yaml:"style"`
// 		} `yaml:"normal"`
// 		Bold struct {
// 			Family string `yaml:"family"`
// 			Style  string `yaml:"style"`
// 		} `yaml:"bold"`
// 		Italic struct {
// 			Family string `yaml:"family"`
// 			Style  string `yaml:"style"`
// 		} `yaml:"italic"`
// 		BoldItalic struct {
// 			Family string `yaml:"family"`
// 			Style  string `yaml:"style"`
// 		} `yaml:"bold_italic"`
// 		Size   float64 `yaml:"size"`
// 		Offset struct {
// 			X int `yaml:"x"`
// 			Y int `yaml:"y"`
// 		} `yaml:"offset"`
// 		GlyphOffset struct {
// 			X int `yaml:"x"`
// 			Y int `yaml:"y"`
// 		} `yaml:"glyph_offset"`
// 		BuiltinBoxDrawing bool `yaml:"builtin_box_drawing"`
// 	} `yaml:"font"`
// 	DrawBoldTextWithBrightColors bool `yaml:"draw_bold_text_with_bright_colors"`
// 	Colors                       struct {
// 		Primary struct {
// 			Background       string `yaml:"background"`
// 			Foreground       string `yaml:"foreground"`
// 			DimForeground    string `yaml:"dim_foreground"`
// 			BrightForeground string `yaml:"bright_foreground"`
// 		} `yaml:"primary"`
// 		Cursor struct {
// 			Text   string `yaml:"text"`
// 			Cursor string `yaml:"cursor"`
// 		} `yaml:"cursor"`
// 		ViModeCursor struct {
// 			Text   string `yaml:"text"`
// 			Cursor string `yaml:"cursor"`
// 		} `yaml:"vi_mode_cursor"`
// 		Search struct {
// 			Matches struct {
// 				Foreground string `yaml:"foreground"`
// 				Background string `yaml:"background"`
// 			} `yaml:"matches"`
// 			FocusedMatch struct {
// 				Foreground string `yaml:"foreground"`
// 				Background string `yaml:"background"`
// 			} `yaml:"focused_match"`
// 		} `yaml:"search"`
// 		Hints struct {
// 			Start struct {
// 				Foreground string `yaml:"foreground"`
// 				Background string `yaml:"background"`
// 			} `yaml:"start"`
// 			End struct {
// 				Foreground string `yaml:"foreground"`
// 				Background string `yaml:"background"`
// 			} `yaml:"end"`
// 		} `yaml:"hints"`
// 		LineIndicator struct {
// 			Foreground string `yaml:"foreground"`
// 			Background string `yaml:"background"`
// 		} `yaml:"line_indicator"`
// 		FooterBar struct {
// 			Background string `yaml:"background"`
// 			Foreground string `yaml:"foreground"`
// 		} `yaml:"footer_bar"`
// 		Selection struct {
// 			Text       string `yaml:"text"`
// 			Background string `yaml:"background"`
// 		} `yaml:"selection"`
// 		Normal struct {
// 			Black   string `yaml:"black"`
// 			Red     string `yaml:"red"`
// 			Green   string `yaml:"green"`
// 			Yellow  string `yaml:"yellow"`
// 			Blue    string `yaml:"blue"`
// 			Magenta string `yaml:"magenta"`
// 			Cyan    string `yaml:"cyan"`
// 			White   string `yaml:"white"`
// 		} `yaml:"normal"`
// 		Bright struct {
// 			Black   string `yaml:"black"`
// 			Red     string `yaml:"red"`
// 			Green   string `yaml:"green"`
// 			Yellow  string `yaml:"yellow"`
// 			Blue    string `yaml:"blue"`
// 			Magenta string `yaml:"magenta"`
// 			Cyan    string `yaml:"cyan"`
// 			White   string `yaml:"white"`
// 		} `yaml:"bright"`
// 		Dim struct {
// 			Black   string `yaml:"black"`
// 			Red     string `yaml:"red"`
// 			Green   string `yaml:"green"`
// 			Yellow  string `yaml:"yellow"`
// 			Blue    string `yaml:"blue"`
// 			Magenta string `yaml:"magenta"`
// 			Cyan    string `yaml:"cyan"`
// 			White   string `yaml:"white"`
// 		} `yaml:"dim"`
// 		IndexedColors []struct {
// 			Index int    `yaml:"index"`
// 			Color string `yaml:"color"`
// 		} `yaml:"indexed_colors"`
// 		TransparentBackgroundColors bool `yaml:"transparent_background_colors"`
// 	} `yaml:"colors"`
// 	Bell struct {
// 		Animation string `yaml:"animation"`
// 		Duration  int    `yaml:"duration"`
// 		Color     string `yaml:"color"`
// 		Command   string `yaml:"command"`
// 	} `yaml:"bell"`
// 	Selection struct {
// 		SemanticEscapeChars string `yaml:"semantic_escape_chars"`
// 		SaveToClipboard     bool   `yaml:"save_to_clipboard"`
// 	} `yaml:"selection"`
// 	Cursor struct {
// 		Style struct {
// 			Shape    string `yaml:"shape"`
// 			Blinking string `yaml:"blinking"`
// 		} `yaml:"style"`
// 		ViModeStyle     string  `yaml:"vi_mode_style"`
// 		BlinkInterval   int     `yaml:"blink_interval"`
// 		BlinkTimeout    int     `yaml:"blink_timeout"`
// 		UnfocusedHollow bool    `yaml:"unfocused_hollow"`
// 		Thickness       float64 `yaml:"thickness"`
// 	} `yaml:"cursor"`
// 	LiveConfigReload bool `yaml:"live_config_reload"`
// 	Shell            struct {
// 		Program string   `yaml:"program"`
// 		Args    []string `yaml:"args"`
// 	} `yaml:"shell"`
// 	WorkingDirectory string `yaml:"working_directory"`
// 	AltSendEsc       bool   `yaml:"alt_send_esc"`
// 	IpcSocket        bool   `yaml:"ipc_socket"`
// 	Mouse            struct {
// 		DoubleClick struct {
// 			Threshold int `yaml:"threshold"`
// 		} `yaml:"double_click"`
// 		TripleClick struct {
// 			Threshold int `yaml:"threshold"`
// 		} `yaml:"triple_click"`
// 		HideWhenTyping bool `yaml:"hide_when_typing"`
// 	} `yaml:"mouse"`
// 	Hints struct {
// 		Alphabet string `yaml:"alphabet"`
// 		Enabled  []struct {
// 			Regex          string `yaml:"regex"`
// 			Hyperlinks     bool   `yaml:"hyperlinks"`
// 			Command        string `yaml:"command"`
// 			PostProcessing bool   `yaml:"post_processing"`
// 			Mouse          struct {
// 				Enabled bool   `yaml:"enabled"`
// 				Mods    string `yaml:"mods"`
// 			} `yaml:"mouse"`
// 			Binding struct {
// 				Key  string `yaml:"key"`
// 				Mods string `yaml:"mods"`
// 			} `yaml:"binding"`
// 		} `yaml:"enabled"`
// 	} `yaml:"hints"`
// 	MouseBindings []struct {
// 		Mouse  string `yaml:"mouse"`
// 		Action string `yaml:"action"`
// 		Mods   string `yaml:"mods,omitempty"`
// 		Mode   string `yaml:"mode,omitempty"`
// 	} `yaml:"mouse_bindings"`
// 	KeyBindings []struct {
// 		Key    string `yaml:"key"`
// 		Action string `yaml:"action"`
// 		Mods   string `yaml:"mods,omitempty"`
// 	} `yaml:"key_bindings"`
// 	Debug struct {
// 		RenderTimer       bool   `yaml:"render_timer"`
// 		PersistentLogging bool   `yaml:"persistent_logging"`
// 		LogLevel          string `yaml:"log_level"`
// 		PrintEvents       bool   `yaml:"print_events"`
// 		HighlightDamage   bool   `yaml:"highlight_damage"`
// 	} `yaml:"debug"`
// }
