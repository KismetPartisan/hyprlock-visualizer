# hyprlock-visualizer




*Inspired by https://github.com/Fubi118/hyprlock-equalizer but written in go cause I was having performance issues*


Set of scripts to read from cava's raw ascii output and display it on your lock screen




https://github.com/user-attachments/assets/36e7f34c-2e13-45e4-9ba0-c5bf22372c3f

*Example lockscreen adapted from https://github.com/SatanAlphabet/dotfiles/blob/main/dot_config/hypr/hyprlock/theme-1-alt.conf*



## Installation
1. Clone repo: `git clone https://github.com/kismetpartisan/hyprlock-visualizer`
2. Move to repo and build with go: `mv hyprlock-visualizer && go build .`
3. Create a label in your hyprlock.conf to run the visualizer, something like this:
   ```
   label {
     text = cmd[update:16] ~/hyprlock-visualizer/visualizer
     font_size = 12
     position = 0, 0
     halign = center
     valign = bottom
     zindex = 2
   }
4. Run launch-hyprlock.sh wherever you would normally just run hyprlock


## Configuration
To change the number of bars displayed, change `bars = 40` in cava_config  
To change the height of the bars, set `const LAYERS = 13` in visualizer.go to the number of text lines you would like the visualizer to take up, then set `ascii_max_range = 104` in cava_config to 8 times whatever you set LAYERS to, finally rebuild with `go build .`
