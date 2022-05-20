package utils
import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

func GetNumBetween(min, max int) int{
	return randGenerator.Intn(max - min) + min
}