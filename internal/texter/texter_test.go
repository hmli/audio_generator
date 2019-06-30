package texter

import (
	"testing"
)

var testStr = `
	伴随着二次元文化而生的A站，从诞生之初就充斥着亚文化的气质：“ACFun”取意自“Anime Comic Fun”，以及“天下漫友是一家”，
一小群屡遭误解、排挤的二次元文化爱好者聚拢在一起，用小圈层的自娱自乐反叛而不羁地解构主流文化	
`

func TestTextAutoMerge(t *testing.T) {

}

func TestTextEveryLine(t *testing.T) {
	result := TextEveryLine([]byte(testStr))
	for _, l := range result {
		t.Log(l)
	}
}
