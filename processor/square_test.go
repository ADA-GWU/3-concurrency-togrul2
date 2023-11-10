package processor

import (
	"image/color"
	"testing"
)

// Unittest for testing average color func.
func TestCalcAverageColor(t *testing.T) {
	var colors = [][]color.Color{
		{color.RGBA{172, 10, 127, 255}, color.RGBA{140, 47, 170, 255}, color.RGBA{196, 151, 117, 255}, color.RGBA{166, 22, 183, 255}, color.RGBA{192, 204, 33, 255}, color.RGBA{216, 67, 179, 255}, color.RGBA{78, 154, 251, 255}, color.RGBA{82, 162, 219, 255}, color.RGBA{195, 118, 125, 255}, color.RGBA{139, 103, 125, 255}},
		{color.RGBA{229, 216, 9, 255}, color.RGBA{164, 116, 108, 255}, color.RGBA{211, 222, 161, 255}, color.RGBA{159, 21, 81, 255}, color.RGBA{89, 165, 242, 255}, color.RGBA{214, 102, 98, 255}, color.RGBA{36, 183, 5, 255}, color.RGBA{112, 87, 58, 255}, color.RGBA{43, 76, 70, 255}, color.RGBA{60, 75, 228, 255}},
		{color.RGBA{216, 189, 132, 255}, color.RGBA{14, 88, 154, 255}, color.RGBA{178, 246, 140, 255}, color.RGBA{205, 204, 69, 255}, color.RGBA{58, 57, 41, 255}, color.RGBA{98, 193, 66, 255}, color.RGBA{72, 122, 230, 255}, color.RGBA{125, 174, 202, 255}, color.RGBA{39, 74, 234, 255}, color.RGBA{207, 87, 168, 255}},
		{color.RGBA{101, 135, 174, 255}, color.RGBA{200, 223, 122, 255}, color.RGBA{88, 94, 107, 255}, color.RGBA{145, 81, 139, 255}, color.RGBA{141, 100, 165, 255}, color.RGBA{230, 243, 236, 255}, color.RGBA{25, 66, 9, 255}, color.RGBA{214, 77, 107, 255}, color.RGBA{47, 18, 72, 255}, color.RGBA{152, 95, 86, 255}},
		{color.RGBA{9, 27, 78, 255}, color.RGBA{22, 148, 151, 255}, color.RGBA{238, 165, 115, 255}, color.RGBA{8, 45, 5, 255}, color.RGBA{208, 19, 69, 255}, color.RGBA{94, 243, 146, 255}, color.RGBA{38, 213, 197, 255}, color.RGBA{30, 8, 245, 255}, color.RGBA{254, 71, 53, 255}, color.RGBA{199, 79, 7, 255}},
		{color.RGBA{238, 35, 175, 255}, color.RGBA{29, 185, 222, 255}, color.RGBA{192, 9, 190, 255}, color.RGBA{222, 82, 187, 255}, color.RGBA{134, 250, 99, 255}, color.RGBA{96, 62, 121, 255}, color.RGBA{216, 167, 149, 255}, color.RGBA{204, 177, 124, 255}, color.RGBA{8, 205, 243, 255}, color.RGBA{130, 35, 118, 255}},
		{color.RGBA{29, 3, 62, 255}, color.RGBA{133, 147, 194, 255}, color.RGBA{208, 199, 147, 255}, color.RGBA{12, 203, 173, 255}, color.RGBA{142, 59, 71, 255}, color.RGBA{30, 167, 97, 255}, color.RGBA{123, 184, 32, 255}, color.RGBA{221, 209, 163, 255}, color.RGBA{193, 63, 255, 255}, color.RGBA{148, 9, 205, 255}},
		{color.RGBA{178, 36, 185, 255}, color.RGBA{74, 145, 137, 255}, color.RGBA{127, 210, 213, 255}, color.RGBA{241, 32, 162, 255}, color.RGBA{52, 194, 31, 255}, color.RGBA{218, 151, 133, 255}, color.RGBA{202, 194, 28, 255}, color.RGBA{27, 244, 72, 255}, color.RGBA{39, 106, 151, 255}, color.RGBA{224, 61, 121, 255}},
		{color.RGBA{163, 234, 185, 255}, color.RGBA{67, 254, 121, 255}, color.RGBA{179, 47, 203, 255}, color.RGBA{45, 52, 198, 255}, color.RGBA{114, 171, 166, 255}, color.RGBA{188, 183, 68, 255}, color.RGBA{198, 116, 28, 255}, color.RGBA{216, 111, 55, 255}, color.RGBA{34, 227, 132, 255}, color.RGBA{145, 128, 168, 255}},
		{color.RGBA{157, 34, 128, 255}, color.RGBA{104, 207, 4, 255}, color.RGBA{164, 231, 250, 255}, color.RGBA{82, 53, 92, 255}, color.RGBA{29, 158, 133, 255}, color.RGBA{23, 81, 38, 255}, color.RGBA{38, 78, 178, 255}, color.RGBA{156, 232, 13, 255}, color.RGBA{234, 56, 244, 255}, color.RGBA{98, 239, 157, 255}},
	}
	expected := color.RGBA{141, 119, 115, 255}
	actual := calculateAverageColor(colors, 0, 0, 5, 5)

	if actual != expected {
		t.Fatal("Color is different from expected")
	}
}
