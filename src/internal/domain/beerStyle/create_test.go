package BeerStyleEntity

import (
	"errors"
	"testing"
)

// Deve criar com sucesso uma nova entidade BeerStyle
func Test_ValidPayload(t *testing.T) {
	name := "IPA"
	minTemp := -7.0
	maxtemp := 10.0

	beerStyle, err := New(nil, name, nil, minTemp, maxtemp, nil)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if beerStyle == nil {
		t.Errorf("expected beerStyle, got nil")
	}

	if beerStyle.Name != name {
		t.Errorf("expected name %s, got %s", name, beerStyle.Name)
	}

	if beerStyle.TempType != Celsius {
		t.Errorf("expected temperature type %s, got %s", Celsius, beerStyle.TempType)
	}
}

// Deve calcular corretamente a temperatura média de um BeerStyle
func Test_CalculateAverage(t *testing.T) {
	beerStyle, _ := New(nil, "Test", nil, -4, 2, nil)

	avg := beerStyle.AverageTemperature()

	expected := -1.0
	if avg != expected {
		t.Errorf("expected average %v, got %v", expected, avg)
	}
}

// Deve retornar nil quando o nome for vazio
func Test_EmptyName(t *testing.T) {
	name := ""
	minTemp := 2.0
	maxtemp := 3.0

	_, err := New(nil, name, nil, minTemp, maxtemp, nil)

	if err == nil {
		t.Fatal("expected error, got nil.")
	}

	if !errors.Is(err, ErrNameRequired) {
		t.Errorf("expected ErrNameRequired, got %v", err)
	}
}

// Deve retornar nil quando temperatura mínima for maior que máxima
func Test_MinTempGreaterThanMax(t *testing.T) {
	name := "Test"
	minTemp := 3.0
	maxtemp := 2.0

	_, err := New(nil, name, nil, minTemp, maxtemp, nil)

	if err == nil {
		t.Fatal("expected error, got nil.")
	}

	if !errors.Is(err, ErrInvalidTemperature) {
		t.Errorf("expected ErrInvalidtemperature, got %v", err)
	}
}

// Deve retornar nil quando o tipo de temperatura for diferente de Celsius, Fahrenheit e Kelvin
func Test_TemperatureTypeErr(t *testing.T) {
	tempType := "Non exist"

	_, err := New(nil, "TempTest", nil, -5, 5, (*TemperatureType)(&tempType))

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !errors.Is(err, ErrInvalidTemperatureType) {
		t.Errorf("expected ErrInvalidTemperatureType, got %v", err)
	}
}
