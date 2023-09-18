package calculatecontroller

import (
	"math"
	"testing"
)

func TestHitungAkar(t *testing.T) {
	testingData := [4]int{100, 144, 300, 400}
	i := 0
	for i < 4 {
		result := hitungakar(float64(testingData[i]))
		hasilSeharusnya := math.Sqrt(float64(testingData[i]))
		hasilSeharusnya = math.Floor(hasilSeharusnya * 100000) / 100000

		t.Logf("Input: %d, Hasil Akar: %f, Hasil Seharusnya: %f", testingData[i], result, hasilSeharusnya)
		if result != hasilSeharusnya {
			t.Errorf("SALAH! Untuk input %d, seharusnya %f", testingData[i], hasilSeharusnya)
		}
		i++
	}
}

func TestBoundary(t *testing.T) {
	testingData := [4]int{15000, 1, 2520000000000, 9000032800}
	i := 0
	for i < 4 {
		result := hitungakar(float64(testingData[i]))
		hasilSeharusnya := math.Sqrt(float64(testingData[i]))
		hasilSeharusnya = math.Floor(hasilSeharusnya*100000) / 100000

		t.Logf("Input: %d, Hasil Akar: %f, Hasil Seharusnya: %f", testingData[i], result, hasilSeharusnya)
		if result != hasilSeharusnya {
			t.Errorf("SALAH! Untuk input %d, seharusnya %f", testingData[i], hasilSeharusnya)
		}
		i++
	}
}
