package flex

import (
	"testing"
)

func TestAlignContentFlexStart(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 130)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 10)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeStyleSetHeight(rootChild4, 10)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 130, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 130, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 80, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 80, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_flex_start_without_height_on_children(t *testing.T) {

	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 10)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_flex_start_with_flex(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 120)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0, 0)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetFlexBasisPercent(rootChild1, 0)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild3, 1)
	NodeStyleSetFlexShrink(rootChild3, 1)
	NodeStyleSetFlexBasisPercent(rootChild3, 0)
	NodeStyleSetWidth(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 120, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 120, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 120, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 120, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignContent(root, AlignFlexEnd)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 10)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeStyleSetHeight(rootChild4, 10)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 40, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 20, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 30, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 40, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 100, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 100, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_spacebetween(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignSpaceBetween)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 130)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 10)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeStyleSetHeight(rootChild4, 10)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 130, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 45, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 45, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 90, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 130, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 80, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 30, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 80, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 45, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 30, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 45, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 80, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 90, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild4))

}

func TestAlign_content_spacearound(t *testing.T) {

	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignSpaceAround)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 140)
	NodeStyleSetHeight(root, 120)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeStyleSetHeight(rootChild0, 10)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 10)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeStyleSetHeight(rootChild2, 10)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeStyleSetHeight(rootChild3, 10)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeStyleSetHeight(rootChild4, 10)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 140, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 120, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 15, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 15, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 55, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 55, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 95, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 140, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 120, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 90, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 15, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 15, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 90, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 55, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 40, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 55, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 90, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 95, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_children(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0Child0, 1)
	NodeStyleSetFlexShrink(rootChild0Child0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0Child0, 0)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_flex(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetFlexShrink(rootChild1, 1)
	NodeStyleSetFlexBasisPercent(rootChild1, 0)
	NodeStyleSetWidth(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild3, 1)
	NodeStyleSetFlexShrink(rootChild3, 1)
	NodeStyleSetFlexBasisPercent(rootChild3, 0)
	NodeStyleSetWidth(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 100, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 100, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_flex_no_shrink(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetFlexShrink(rootChild1, 1)
	NodeStyleSetFlexBasisPercent(rootChild1, 0)
	NodeStyleSetWidth(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild3, 1)
	NodeStyleSetFlexBasisPercent(rootChild3, 0)
	NodeStyleSetWidth(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 100, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 100, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_margin(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild1, EdgeLeft, 10)
	NodeStyleSetMargin(rootChild1, EdgeTop, 10)
	NodeStyleSetMargin(rootChild1, EdgeRight, 10)
	NodeStyleSetMargin(rootChild1, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetMargin(rootChild3, EdgeLeft, 10)
	NodeStyleSetMargin(rootChild3, EdgeTop, 10)
	NodeStyleSetMargin(rootChild3, EdgeRight, 10)
	NodeStyleSetMargin(rootChild3, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 60, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 40, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 60, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 40, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 10, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 40, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 40, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 40, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 100, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetPadding(rootChild1, EdgeLeft, 10)
	NodeStyleSetPadding(rootChild1, EdgeTop, 10)
	NodeStyleSetPadding(rootChild1, EdgeRight, 10)
	NodeStyleSetPadding(rootChild1, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetPadding(rootChild3, EdgeLeft, 10)
	NodeStyleSetPadding(rootChild3, EdgeTop, 10)
	NodeStyleSetPadding(rootChild3, EdgeRight, 10)
	NodeStyleSetPadding(rootChild3, EdgeBottom, 10)
	NodeStyleSetWidth(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_single_row(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild1))
}

func TestAlign_content_stretch_row_with_fixed_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetHeight(rootChild1, 60)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 60, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 80, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 80, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_max_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetMaxHeight(rootChild1, 20)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 20, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_row_with_min_height(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 150)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 50)
	NodeStyleSetMinHeight(rootChild1, 80)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 90, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 90, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 100, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 90, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 90, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 90, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 150, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 100, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 90, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 90, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 90, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 100, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 90, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 90, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignContent(root, AlignStretch)
	NodeStyleSetFlexWrap(root, WrapWrap)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 150)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild0, 50)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild0Child0, 1)
	NodeStyleSetFlexShrink(rootChild0Child0, 1)
	NodeStyleSetFlexBasisPercent(rootChild0Child0, 0)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetFlexGrow(rootChild1, 1)
	NodeStyleSetFlexShrink(rootChild1, 1)
	NodeStyleSetFlexBasisPercent(rootChild1, 0)
	NodeStyleSetHeight(rootChild1, 50)
	NodeInsertChild(root, rootChild1, 1)

	rootChild2 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild2, 50)
	NodeInsertChild(root, rootChild2, 2)

	rootChild3 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild3, 50)
	NodeInsertChild(root, rootChild3, 3)

	rootChild4 := NewNodeWithConfig(config)
	NodeStyleSetHeight(rootChild4, 50)
	NodeInsertChild(root, rootChild4, 4)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 150, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 100, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 50, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 150, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0Child0))

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild1))

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 50, NodeLayoutGetTop(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild2))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild2))

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 100, NodeLayoutGetTop(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild3))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild3))

	assertFloatEqual(t, 0, rootChild4.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild4))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild4))
}

func TestAlign_content_stretch_is_not_overriding_align_items(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetAlignContent(root, AlignStretch)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionRow)
	NodeStyleSetAlignContent(rootChild0, AlignStretch)
	NodeStyleSetAlignItems(rootChild0, AlignCenter)
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 100)
	NodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetAlignContent(rootChild0Child0, AlignStretch)
	NodeStyleSetWidth(rootChild0Child0, 10)
	NodeStyleSetHeight(rootChild0Child0, 10)
	NodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 45, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 90, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 45, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0Child0))
}
