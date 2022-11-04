# alaconf
Go program to switch Alacritty colorshemes and other configuration options

## Prerequisites

This program was developed first for MacOS, and should work seamlessly for Linux. The assumption is that the program is run as the user, Alacritty config is in `~/.config/alacritty/alacritty.yml` and color schemes are in `~/.config/alacritty/colors/`.

## Usage

### Set colorscheme

```shell
alaconf --color=color1
```
This command will get `~/.config/alacritty/colors/color1.yml` (or yaml) and replace the colors section in Alacritty config with the content of the color file.

### Get colorscheme

```shell
alaconf --getcolor
```
This command will get the colorscheme currently set (by the previous command) in Alacritty config and display the name (using the header comment of the section)

## Todo

This command is under development, and here is the list of feature that should be developed to make it more robust and useful:
- Ability to specify or configure the config folders (Alacritty config and colors), or get it from default values
- Set the colorscheme when none is present
- Compile and test for Windows
- Change additional parameters, for instance:
    - font
    - Window decoration
    - mouse/key bindings

## Notes

### reference
The reference folder contains the full example of configuration from Alacritty deliverable as a reference. Can be used to map fields, or use (Yaml-to-go)[https://mengzhuo.github.io/yaml-to-go/] to automatically create a config mapper.