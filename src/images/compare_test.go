package images

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareImages_Equals(t *testing.T) {
	imgPath1 := "../../images/normal.jpg"
	imgPath2 := "../../images/equal-normal.jpg"

	assert.True(t, CompareImages(imgPath1, imgPath2))
}

func TestCompareImages_Diff_Cursor(t *testing.T) {
	diffFolderPath = "../../diff/"
	imgNormal := "../../images/normal.jpg"
	imgCursor := "../../images/error-cursor.jpg"
	imgCursorDiff := "../../diff/diff-error-cursor.jpg"
	imgTestDiff := "../../test-diff-images/diff-error-cursor.jpg"

	assert.True(t, !CompareImages(imgCursor, imgNormal))
	assert.True(t, CompareImages(imgCursorDiff, imgTestDiff))
}

func TestCompareImages_Diff_Button(t *testing.T) {
	diffFolderPath = "../../diff/"
	imgNormal := "../../images/normal.jpg"
	imgCursor := "../../images/error-button.jpg"
	imgCursorDiff := "../../diff/diff-error-button.jpg"
	imgTestDiff := "../../test-diff-images/diff-error-button.jpg"

	assert.True(t, !CompareImages(imgCursor, imgNormal))
	assert.True(t, CompareImages(imgCursorDiff, imgTestDiff))
}
