package measurers

import "time"

type Measurers[T Temperature | Energy] struct {
	Serial string    `json:"serial"`
	Date   time.Time `json:"date"`
	Values T         `json:"values"`
}

type Temperature struct {
	Temperature int `json:"temperature"`
}

type Energy struct {
	Energy int `json:"energy"`
}

var TemperatureMeasurer1 = &Measurers[Temperature]{Serial: "3043846b-14e3-4505-9230-75af208b6703"}
var TemperatureMeasurer2 = &Measurers[Temperature]{Serial: "a8bae29f-dd87-4ffb-a69a-5e68948fd1ff"}
var TemperatureMeasurer3 = &Measurers[Temperature]{Serial: "c4d2bac0-6f3f-4de9-9db4-2e6832dac6c1"}
var TemperatureMeasurer4 = &Measurers[Temperature]{Serial: "38d45aae-cecb-44d2-ae3c-bc91437d8be3"}
var TemperatureMeasurer5 = &Measurers[Temperature]{Serial: "79290c70-76f7-4cf7-8e3f-ad35629a07cf"}

var EnergyMeasuser1 = &Measurers[Energy]{Serial: "283bca27-17a6-41d5-9447-2bb7468d36d0"}
var EnergyMeasuser2 = &Measurers[Energy]{Serial: "fbddbad5-10a5-4ff3-b107-d0fc924aee0b"}
var EnergyMeasuser3 = &Measurers[Energy]{Serial: "6facea9e-fe99-4821-97b0-02249bce17df"}
var EnergyMeasuser4 = &Measurers[Energy]{Serial: "d7c74a44-e8a8-4be4-89c4-f5951bfdc000"}
var EnergyMeasuser5 = &Measurers[Energy]{Serial: "f5399d19-975d-4afd-87ee-3993fff662f9"}
