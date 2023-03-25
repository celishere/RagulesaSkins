package RagulesaSkins

import (
	"bytes"
	"fmt"
	"image"
)

type InvalidSkinException struct {
	msg string
}

func (e *InvalidSkinException) Error() string {
	return fmt.Sprintf("Invalid skin data: %s", e.msg)
}

var acceptedSkinBytes = []int{
	64 * 32 * 4,
	64 * 64 * 4,
	128 * 128 * 4,
}

var skinBytesToSize = map[int][2]int{
	64 * 32 * 4:   {64, 32},
	64 * 64 * 4:   {64, 64},
	128 * 128 * 4: {128, 128},
}

func checkSkinSize(len int) error {
	for _, size := range acceptedSkinBytes {
		if len == size {
			return nil
		}
	}
	return &InvalidSkinException{fmt.Sprintf("Invalid skin data size %d bytes (allowed sizes: %v)", len, acceptedSkinBytes)}
}

func ImageToSkinData(img *image.NRGBA) string {
	var outputBuf bytes.Buffer

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			_, err := outputBuf.Write([]byte{uint8(r), uint8(g), uint8(b), uint8(a)})

			if err != nil {
				return ""
			}
		}
	}

	return outputBuf.String()
}

func SkinDataToImage(skinData string) *image.NRGBA {
	size := len(skinData)

	err := checkSkinSize(size)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	imageSize := skinBytesToSize[size]

	xSize := imageSize[0]
	ySize := imageSize[1]

	b := []byte(skinData)

	img := image.NewNRGBA(image.Rect(0, 0, xSize, ySize))
	img.Pix = b

	return img
}
