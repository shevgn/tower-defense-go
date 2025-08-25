package terminal

type BorderType string

const (
	BorderTopLeft     BorderType = "┌"
	BorderTop         BorderType = "─"
	BorderTopRight    BorderType = "┐"
	BorderRight       BorderType = "│"
	BorderBottomRight BorderType = "┘"
	BorderBottom      BorderType = "─"
	BorderBottomLeft  BorderType = "└"
	BorderLeft        BorderType = "│"
	BorderVertical    BorderType = "│"
	BorderHorizontal  BorderType = "─"

	BorderRoundedTopLeft     BorderType = "╭"
	BorderRoundedBottomLeft  BorderType = "╰"
	BorderRoundedTopRight    BorderType = "╮"
	BorderRoundedBottomRight BorderType = "╯"

	BoxDrawingHeavyVertical   BorderType = "┃"
	BoxDrawingHeavyHorizontal BorderType = "━"
	BoxDrawingHeavyCross      BorderType = "┿"

	Cross BorderType = "┼"

	BoxDrawingLightVertical   BorderType = "│"
	BoxDrawingLightHorizontal BorderType = "─"
	BoxDrawingLightCross      BorderType = "┼"
)
