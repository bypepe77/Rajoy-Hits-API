package structs

// definición de de una test suite
import "testing"

// definición de un test concreto sobre la estructura
// Prueba sobre la función constructora
func Test_Constructor(t *testing.T) {
	// GIVEN Situación inicial
	nameOffered := "altas"
	channelOffered := "altaskur"
	viewersOffered := 30
	// WHEN Acciones que realizas
	// ejecución dela función constructora
	streamer := NewStreamer(nameOffered, channelOffered, viewersOffered)
	// THEN Comprobaciones
	// comprobar que los datos están en su sitio
	// Assert (comprobación de estado de la variables)
	nameGetted := streamer.GetName()
	channelGetted := streamer.GetChannel()
	viewersGetted := streamer.GetViewers()

	// assert.Equals(t, viewersGetted, viewersOffered, "")

	if nameGetted != nameOffered {
		t.Errorf("Se esperaba que fuera: " + nameOffered)
	}
	if channelGetted != channelOffered {
		t.Errorf("Se esperaba que fuera: " + channelOffered)
	}
	if viewersGetted != viewersOffered {
		// `Mi nombre es ${nombre}, tengo ${edad} años y mido ${altura} metros.`
		t.Errorf("Se esperaba %d que fuera test", viewersGetted)
		t.Errorf("Se esperaba ${viewersGetted} que fuera test")
	}
}

func Test_Setters(t *testing.T) {
	// GIVEN Situación inicial
	nameOffered := "Pythonesa"
	channelOffered := "Patonesafan"
	viewersOffered := 1010
	// WHEN Acciones que realizas
	// ejecución dela función constructora
	streamer := NewStreamer("altas", "altaskur", 0)
	streamer.SetName(nameOffered)
	streamer.SetChannel(channelOffered)
	streamer.SetViewers(viewersOffered)
	// THEN Comprobaciones
	// comprobar que los datos están en su sitio
	// Assert (comprobación de estado de la variables)
	nameGetted := streamer.GetName()
	channelGetted := streamer.GetChannel()
	viewersGetted := streamer.GetViewers()

	// assert.Equals(t, viewersGetted, viewersOffered, "")

	if nameGetted != nameOffered {
		t.Errorf("Se esperaba que fuera: " + nameOffered)
	}
	if channelGetted != channelOffered {
		t.Errorf("Se esperaba que fuera: " + channelOffered)
	}
	if viewersGetted != viewersOffered {
		// `Mi nombre es ${nombre}, tengo ${edad} años y mido ${altura} metros.`
		t.Errorf("Se esperaba %d que fuera test", viewersGetted)
		t.Errorf("Se esperaba ${viewersGetted} que fuera test")
	}
}
