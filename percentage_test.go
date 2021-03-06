package flex

import "testing"

func TestPercentage_width_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidthPercent(30)
	rootChild0.StyleSetHeightPercent(30)
	root.InsertChild(rootChild0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 140, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild0.LayoutGetHeight())
}

func TestPercentage_position_left_top(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(400)
	root.StyleSetHeight(400)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionPercent(EdgeLeft, 10)
	rootChild0.StyleSetPositionPercent(EdgeTop, 20)
	rootChild0.StyleSetWidthPercent(45)
	rootChild0.StyleSetHeightPercent(55)
	root.InsertChild(rootChild0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 400, root.LayoutGetWidth())
	assertFloatEqual(t, 400, root.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 180, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 220, rootChild0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 400, root.LayoutGetWidth())
	assertFloatEqual(t, 400, root.LayoutGetHeight())

	assertFloatEqual(t, 260, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 80, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 180, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 220, rootChild0.LayoutGetHeight())
}

func TestPercentage_position_bottom_right(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(500)
	root.StyleSetHeight(500)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionPercent(EdgeRight, 20)
	rootChild0.StyleSetPositionPercent(EdgeBottom, 10)
	rootChild0.StyleSetWidthPercent(55)
	rootChild0.StyleSetHeightPercent(15)
	root.InsertChild(rootChild0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 500, root.LayoutGetWidth())
	assertFloatEqual(t, 500, root.LayoutGetHeight())

	assertFloatEqual(t, -100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, -50, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 275, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 75, rootChild0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 500, root.LayoutGetWidth())
	assertFloatEqual(t, 500, root.LayoutGetHeight())

	assertFloatEqual(t, 125, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, -50, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 275, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 75, rootChild0.LayoutGetHeight())
}

func TestPercentage_flex_basis(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetFlexBasisPercent(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexGrow(1)
	rootChild1.StyleSetFlexBasisPercent(25)
	root.InsertChild(rootChild1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 125, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 125, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 75, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 75, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 125, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 75, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild1.LayoutGetHeight())
}

func TestPercentage_flex_basis_cross(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetFlexBasisPercent(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexGrow(1)
	rootChild1.StyleSetFlexBasisPercent(25)
	root.InsertChild(rootChild1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 125, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 125, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 75, rootChild1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 125, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 125, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 75, rootChild1.LayoutGetHeight())
}

func TestPercentage_flex_basis_cross_min_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetMinHeightPercent(60)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexGrow(2)
	rootChild1.StyleSetMinHeightPercent(10)
	root.InsertChild(rootChild1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 140, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 140, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 140, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 140, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild1.LayoutGetHeight())
}

func TestPercentage_flex_basis_main_max_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetFlexBasisPercent(10)
	rootChild0.StyleSetMaxHeightPercent(60)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexGrow(4)
	rootChild1.StyleSetFlexBasisPercent(10)
	rootChild1.StyleSetMaxHeightPercent(20)
	root.InsertChild(rootChild1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 52, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 120, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 52, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 148, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 148, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 52, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 120, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 148, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild1.LayoutGetHeight())
}

func TestPercentage_flex_basis_cross_max_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetFlexBasisPercent(10)
	rootChild0.StyleSetMaxHeightPercent(60)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexGrow(4)
	rootChild1.StyleSetFlexBasisPercent(10)
	rootChild1.StyleSetMaxHeightPercent(20)
	root.InsertChild(rootChild1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 120, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 120, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 120, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 120, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 40, rootChild1.LayoutGetHeight())
}

func TestPercentage_flex_basis_main_max_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetFlexBasisPercent(15)
	rootChild0.StyleSetMaxWidthPercent(60)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexGrow(4)
	rootChild1.StyleSetFlexBasisPercent(10)
	rootChild1.StyleSetMaxWidthPercent(20)
	root.InsertChild(rootChild1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 120, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 120, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 120, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild1.LayoutGetHeight())
}

func TestPercentage_flex_basis_cross_max_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetFlexBasisPercent(10)
	rootChild0.StyleSetMaxWidthPercent(60)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexGrow(4)
	rootChild1.StyleSetFlexBasisPercent(15)
	rootChild1.StyleSetMaxWidthPercent(20)
	root.InsertChild(rootChild1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 120, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 150, rootChild1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 120, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 160, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 150, rootChild1.LayoutGetHeight())
}

func TestPercentage_flex_basis_main_min_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetFlexBasisPercent(15)
	rootChild0.StyleSetMinWidthPercent(60)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexGrow(4)
	rootChild1.StyleSetFlexBasisPercent(10)
	rootChild1.StyleSetMinWidthPercent(20)
	root.InsertChild(rootChild1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 120, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 120, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 120, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 80, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 200, rootChild1.LayoutGetHeight())
}

func TestPercentage_flex_basis_cross_min_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetFlexBasisPercent(10)
	rootChild0.StyleSetMinWidthPercent(60)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexGrow(4)
	rootChild1.StyleSetFlexBasisPercent(15)
	rootChild1.StyleSetMinWidthPercent(20)
	root.InsertChild(rootChild1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 150, rootChild1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 150, rootChild1.LayoutGetHeight())
}

func TestPercentage_multiple_nested_with_padding_margin_and_percentage_values(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetFlexBasisPercent(10)
	rootChild0.StyleSetMargin(EdgeLeft, 5)
	rootChild0.StyleSetMargin(EdgeTop, 5)
	rootChild0.StyleSetMargin(EdgeRight, 5)
	rootChild0.StyleSetMargin(EdgeBottom, 5)
	rootChild0.StyleSetPadding(EdgeLeft, 3)
	rootChild0.StyleSetPadding(EdgeTop, 3)
	rootChild0.StyleSetPadding(EdgeRight, 3)
	rootChild0.StyleSetPadding(EdgeBottom, 3)
	rootChild0.StyleSetMinWidthPercent(60)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetMargin(EdgeLeft, 5)
	rootChild0Child0.StyleSetMargin(EdgeTop, 5)
	rootChild0Child0.StyleSetMargin(EdgeRight, 5)
	rootChild0Child0.StyleSetMargin(EdgeBottom, 5)
	rootChild0Child0.StyleSetPaddingPercent(EdgeLeft, 3)
	rootChild0Child0.StyleSetPaddingPercent(EdgeTop, 3)
	rootChild0Child0.StyleSetPaddingPercent(EdgeRight, 3)
	rootChild0Child0.StyleSetPaddingPercent(EdgeBottom, 3)
	rootChild0Child0.StyleSetWidthPercent(50)
	rootChild0.InsertChild(rootChild0Child0, 0)

	rootChild0Child0Child0 := NewNodeWithConfig(config)
	rootChild0Child0Child0.StyleSetMarginPercent(EdgeLeft, 5)
	rootChild0Child0Child0.StyleSetMarginPercent(EdgeTop, 5)
	rootChild0Child0Child0.StyleSetMarginPercent(EdgeRight, 5)
	rootChild0Child0Child0.StyleSetMarginPercent(EdgeBottom, 5)
	rootChild0Child0Child0.StyleSetPadding(EdgeLeft, 3)
	rootChild0Child0Child0.StyleSetPadding(EdgeTop, 3)
	rootChild0Child0Child0.StyleSetPadding(EdgeRight, 3)
	rootChild0Child0Child0.StyleSetPadding(EdgeBottom, 3)
	rootChild0Child0Child0.StyleSetWidthPercent(45)
	rootChild0Child0.InsertChild(rootChild0Child0Child0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexGrow(4)
	rootChild1.StyleSetFlexBasisPercent(15)
	rootChild1.StyleSetMinWidthPercent(20)
	root.InsertChild(rootChild1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 5, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 190, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 48, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 8, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 8, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 92, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 36, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 6, rootChild0Child0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 58, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 142, rootChild1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 5, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 190, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 48, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 90, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 8, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 92, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 46, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 36, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 6, rootChild0Child0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 58, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 142, rootChild1.LayoutGetHeight())
}

func TestPercentage_margin_should_calculate_based_only_on_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(200)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetMarginPercent(EdgeLeft, 10)
	rootChild0.StyleSetMarginPercent(EdgeTop, 10)
	rootChild0.StyleSetMarginPercent(EdgeRight, 10)
	rootChild0.StyleSetMarginPercent(EdgeBottom, 10)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetWidth(10)
	rootChild0Child0.StyleSetHeight(10)
	rootChild0.InsertChild(rootChild0Child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 160, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 20, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 160, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 150, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetHeight())
}

func TestPercentage_padding_should_calculate_based_only_on_width(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(200)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexGrow(1)
	rootChild0.StyleSetPaddingPercent(EdgeLeft, 10)
	rootChild0.StyleSetPaddingPercent(EdgeTop, 10)
	rootChild0.StyleSetPaddingPercent(EdgeRight, 10)
	rootChild0.StyleSetPaddingPercent(EdgeBottom, 10)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetWidth(10)
	rootChild0Child0.StyleSetHeight(10)
	rootChild0.InsertChild(rootChild0Child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 20, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 200, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 170, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetHeight())
}

func TestPercentage_absolute_position(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(200)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPositionPercent(EdgeLeft, 30)
	rootChild0.StyleSetPositionPercent(EdgeTop, 10)
	rootChild0.StyleSetWidth(10)
	rootChild0.StyleSetHeight(10)
	root.InsertChild(rootChild0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 60, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 60, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())
}

func TestPercentage_width_height_undefined_parent_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidthPercent(50)
	rootChild0.StyleSetHeightPercent(50)
	root.InsertChild(rootChild0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 0, root.LayoutGetWidth())
	assertFloatEqual(t, 0, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 0, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 0, root.LayoutGetWidth())
	assertFloatEqual(t, 0, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 0, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild0.LayoutGetHeight())
}

func TestPercent_within_flex_grow(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetWidth(350)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(100)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexGrow(1)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidthPercent(100)
	rootChild1.InsertChild(rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(100)
	root.InsertChild(rootChild2, 2)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 350, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 100, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 150, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 150, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 250, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild2.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 350, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 250, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 100, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 150, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 150, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 100, rootChild2.LayoutGetHeight())
}

func TestPercentage_container_in_wrapping_container(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetWidth(200)
	root.StyleSetHeight(200)

	rootChild0 := NewNodeWithConfig(config)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetFlexDirection(FlexDirectionRow)
	rootChild0Child0.StyleSetJustifyContent(JustifyCenter)
	rootChild0Child0.StyleSetWidthPercent(100)
	rootChild0.InsertChild(rootChild0Child0, 0)

	rootChild0Child0Child0 := NewNodeWithConfig(config)
	rootChild0Child0Child0.StyleSetWidth(50)
	rootChild0Child0Child0.StyleSetHeight(50)
	rootChild0Child0.InsertChild(rootChild0Child0Child0, 0)

	rootChild0Child0_child1 := NewNodeWithConfig(config)
	rootChild0Child0_child1.StyleSetWidth(50)
	rootChild0Child0_child1.StyleSetHeight(50)
	rootChild0Child0.InsertChild(rootChild0Child0_child1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 75, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0Child0Child0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0Child0_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0_child1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0Child0_child1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0Child0_child1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 200, root.LayoutGetWidth())
	assertFloatEqual(t, 200, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 75, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0Child0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0_child1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0Child0_child1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0Child0_child1.LayoutGetHeight())
}

func TestPercent_absolute_position(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(60)
	root.StyleSetHeight(50)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetFlexDirection(FlexDirectionRow)
	rootChild0.StyleSetPositionType(PositionTypeAbsolute)
	rootChild0.StyleSetPositionPercent(EdgeLeft, 50)
	rootChild0.StyleSetWidthPercent(100)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetWidthPercent(100)
	rootChild0.InsertChild(rootChild0Child0, 0)

	rootChild0child1 := NewNodeWithConfig(config)
	rootChild0child1.StyleSetWidthPercent(100)
	rootChild0.InsertChild(rootChild0child1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 60, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 60, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0child1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 60, root.LayoutGetWidth())
	assertFloatEqual(t, 50, root.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, -60, rootChild0child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0child1.LayoutGetTop())
	assertFloatEqual(t, 60, rootChild0child1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0child1.LayoutGetHeight())
}
