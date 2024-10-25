package main

import (
    "fmt"
    "github.com/fatih/color"
)

type ground int

const (
    fg ground = iota
    bg
)

const (
    fgLabel = "Fg"
    bgLabel = "Bg"
)

var groundLabels = map[ground]string{
    fg: fgLabel,
    bg: bgLabel,
}

func (g ground) String() string {
    return groundLabels[g]
}

type dc struct {
    Ground ground
    Color  color.Attribute
    Label  string
}

func (d *dc) String() string {
    return fmt.Sprintf("%s::%s", d.Ground, d.Label)
}

var (
    colorDefault = color.Set(
        color.Reset,
        color.ResetBold,
        color.ResetBlinking,
        color.ResetConcealed,
        color.ResetCrossedOut,
        color.ResetItalic,
        color.ResetReversed,
        color.ResetUnderline,
    )

    foregroundColors = []*dc{
        &dc{
            Ground: fg,
            Color:  color.FgWhite,
            Label:  "White",
        },
        &dc{
            Ground: fg,
            Color:  color.FgBlack,
            Label:  "Black",
        },
        &dc{
            Ground: fg,
            Color:  color.FgMagenta,
            Label:  "Magenta",
        },
        &dc{
            Ground: fg,
            Color:  color.FgRed,
            Label:  "Red",
        },
        &dc{
            Ground: fg,
            Color:  color.FgGreen,
            Label:  "Green",
        },
        &dc{
            Ground: fg,
            Color:  color.FgYellow,
            Label:  "Yellow",
        },
        &dc{
            Ground: fg,
            Color:  color.FgCyan,
            Label:  "Cyan",
        },
        &dc{
            Ground: fg,
            Color:  color.FgBlue,
            Label:  "Blue",
        },
    }

    foregroundHiColors = []*dc{
        &dc{
            Ground: fg,
            Color:  color.FgHiWhite,
            Label:  "Hi White",
        },
        &dc{
            Ground: fg,
            Color:  color.FgHiBlack,
            Label:  "Hi Black",
        },
        &dc{
            Ground: fg,
            Color:  color.FgHiMagenta,
            Label:  "Hi Magenta",
        },
        &dc{
            Ground: fg,
            Color:  color.FgHiRed,
            Label:  "Hi Red",
        },
        &dc{
            Ground: fg,
            Color:  color.FgHiGreen,
            Label:  "Hi Green",
        },
        &dc{
            Ground: fg,
            Color:  color.FgHiYellow,
            Label:  "Hi Yellow",
        },
        &dc{
            Ground: fg,
            Color:  color.FgHiCyan,
            Label:  "Hi Cyan",
        },
        &dc{
            Ground: fg,
            Color:  color.FgHiBlue,
            Label:  "Hi Blue",
        },
    }

    backgroundColors = []*dc{
        &dc{
            Ground: bg,
            Color:  color.BgWhite,
            Label:  "White",
        },
        &dc{
            Ground: bg,
            Color:  color.BgBlack,
            Label:  "Black",
        },
        &dc{
            Ground: bg,
            Color:  color.BgMagenta,
            Label:  "Magenta",
        },
        &dc{
            Ground: bg,
            Color:  color.BgRed,
            Label:  "Red",
        },
        &dc{
            Ground: bg,
            Color:  color.BgGreen,
            Label:  "Green",
        },
        &dc{
            Ground: bg,
            Color:  color.BgYellow,
            Label:  "Yellow",
        },
        &dc{
            Ground: bg,
            Color:  color.BgCyan,
            Label:  "Cyan",
        },
        &dc{
            Ground: bg,
            Color:  color.BgBlue,
            Label:  "Blue",
        },
    }

    backgroundHiColors = []*dc{
        &dc{
            Ground: bg,
            Color:  color.BgHiWhite,
            Label:  "Hi White",
        },
        &dc{
            Ground: bg,
            Color:  color.BgHiBlack,
            Label:  "Hi Black",
        },
        &dc{
            Ground: bg,
            Color:  color.BgHiMagenta,
            Label:  "Hi Magenta",
        },
        &dc{
            Ground: bg,
            Color:  color.BgHiRed,
            Label:  "Hi Red",
        },
        &dc{
            Ground: bg,
            Color:  color.BgHiGreen,
            Label:  "Hi Green",
        },
        &dc{
            Ground: bg,
            Color:  color.BgHiYellow,
            Label:  "Hi Yellow",
        },
        &dc{
            Ground: bg,
            Color:  color.BgHiCyan,
            Label:  "Hi Cyan",
        },
        &dc{
            Ground: bg,
            Color:  color.BgHiBlue,
            Label:  "Hi Blue",
        },
    }
)

func main() {
    fmt.Println("Default text.")

    for _, bgColor := range backgroundColors {
        for _, fgColor := range foregroundColors {
            colorString := fmt.Sprintf("%s <> %s", bgColor, fgColor)
            fmt.Printf("%32s >> ", colorString)
            c := color.Set(fgColor.Color, bgColor.Color)
            c.Printf("%s <> %s", bgColor, fgColor)
            colorDefault.Println()
        }
    }

    for _, bgColor := range backgroundHiColors {
        for _, fgColor := range foregroundHiColors {
            colorString := fmt.Sprintf("%s <> %s", bgColor, fgColor)
            fmt.Printf("%32s >> ", colorString)
            c := color.Set(fgColor.Color, bgColor.Color)
            c.Printf("%s <> %s", bgColor, fgColor)
            colorDefault.Println()
        }
    }
}
