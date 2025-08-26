package terminal

type borderType rune

const (
	BorderTopLeft     borderType = '┌'
	BorderTop         borderType = '─'
	BorderTopRight    borderType = '┐'
	BorderRight       borderType = '│'
	BorderBottomRight borderType = '┘'
	BorderBottom      borderType = '─'
	BorderBottomLeft  borderType = '└'
	BorderLeft        borderType = '│'
	BorderVertical    borderType = '│'
	BorderHorizontal  borderType = '─'

	BorderRoundedTopLeft     borderType = '╭'
	BorderRoundedBottomLeft  borderType = '╰'
	BorderRoundedTopRight    borderType = '╮'
	BorderRoundedBottomRight borderType = '╯'

	BoxDrawingHeavyVertical   borderType = '┃'
	BoxDrawingHeavyHorizontal borderType = '━'
	BoxDrawingHeavyCross      borderType = '┿'

	Cross borderType = '┼'

	BoxDrawingLightVertical   borderType = '│'
	BoxDrawingLightHorizontal borderType = '─'
	BoxDrawingLightCross      borderType = '┼'
)
