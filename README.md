## LumatoneHtnlGen

This tool is a quick harmonic table note layout generator for .ltn files, which can be loaded on to a Lumatone.io midi controller via the Lumatone Editor, provided by the controller manufacturer at https://www.lumatone.io/support/software. 

### Basic usage

    lumatoneHtnlGen.exe -f [input.json] > [output.ltn]
 
This will generate a basic harmonic table note layout .ltn file, supporting one channel, given a properly formatted json file for input.

### ltn_def JSON format

The format of the json is an array of layout definitions - one for each section:

    {
        "Board": 0,
        "Key0Pitch": 48,
        "Channel": 1,
        "ColorDefault": "FFFF00",
        "ColorEdge": "00FF7F",
        "ColorC": "0000FF",
        "ColorMiddleC": "FF0080"
    }

Where ...

- "Board" is the id of the board id for the section, and is numbered 0 through 4. Each board value must be unique.
- "Key0Pitch" is the midi pitch value for the upper-leftmost key for a section - Key_0. (The value for middle C is 60).
- "Channel" is the midi channel for the section.

The color settings are what I prefer for harmonic table note layout.

- "ColorDefault" is the "filler" color I use.
- "ColorEdge" outlines an area where each pitch is uniquely named. (Note: this will span a couple of octaves, and skip pitches along the way, in according to harmonic table note layout).
- "ColorC" is the color to use for any key assigned to C (midi value divisible by 12).
- "ColorMiddleC" is the color to apply to any key assigned to middle-C (midi value 60).

### Color Selection

I use tetrad color schemes. Although there are other free online tools for creating color schemes, I used https://www.coderstool.com/tetrad-color-scheme while working on this project.
