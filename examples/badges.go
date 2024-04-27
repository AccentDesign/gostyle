package examples

import (
	"github.com/AccentDesign/gcss/props/align"
	"github.com/AccentDesign/gcss/props/border"
	"github.com/AccentDesign/gcss/props/colors"
	"github.com/AccentDesign/gcss/props/display"
	"github.com/AccentDesign/gcss/props/font"
	"github.com/AccentDesign/gcss/props/unit"
	"github.com/AccentDesign/gcss/style"
)

var Badges = []style.Style{
	{
		Selector: ".badge",
		Props: style.Props{
			AlignItems:    align.ItemsCenter,
			Display:       display.InlineFlex,
			BorderRadius:  radiusFull,
			FontSize:      fontXs,
			FontWeight:    font.WeightSemiBold,
			LineHeight:    leadingNone,
			PaddingTop:    size1,
			PaddingRight:  size3,
			PaddingBottom: size1,
			PaddingLeft:   size3,
		},
	},
	{
		Selector: ".badge-primary",
		Props: style.Props{
			BackgroundColor: primary,
			Color:           primaryForeground,
		},
	},
	{
		Selector: ".badge-secondary",
		Props: style.Props{
			BackgroundColor: secondary,
			Color:           secondaryForeground,
		},
	},
	{
		Selector: ".badge-destructive",
		Props: style.Props{
			BackgroundColor: destructive,
			Color:           destructiveForeground,
		},
	},
	{
		Selector: ".badge-outline",
		Props: style.Props{
			BackgroundColor: colors.Transparent(),
			Border: border.Border{
				Color: borderColor,
				Style: border.StyleSolid,
				Width: unit.Px(1),
			},
		},
	},
}
