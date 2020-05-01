package main

import (
	"encoding/json"
	"regexp"
	"stefan/go-astilecton-test/strip"

	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

type PayloadDeviceList struct {
	DeviceList map[int]strip.DeviceData `json:"DeviceList"`
}

type PayloadChangeColor struct {
	Device strip.DeviceData     `json:"device"`
	Color  strip.StripColorData `json:"color"`
}

type PayloadData struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type PayloadAppReady struct {
	Debug bool `json:"debug"`
}

var deviceList map[int]strip.DeviceData

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {

	re, _ := regexp.Compile("^(.*?)(#.*?)?$")
	match := re.FindStringSubmatch(m.Name)

	switch string(match[1]) {

	case "app-ready":
		payloadAppReady := PayloadAppReady{Debug: debug}
		return payloadAppReady, nil
		break

	case "refresh.devices":
		deviceList, err = strip.FindConnectedDevices(true)
		if err != nil {
			return nil, err
		}

		payloadDevices := PayloadDeviceList{DeviceList: deviceList}
		return payloadDevices, nil
		break

	case "data.set":

		var payloadData PayloadData
		if err = json.Unmarshal(m.Payload, &payloadData); err == nil {
			appData[payloadData.Key] = payloadData.Value
			return "saved", nil
		}
		break

	case "data.get":
		var key string
		if err = json.Unmarshal(m.Payload, &key); err == nil {
			payloadDataGet := PayloadData{Key: key, Value: appData[key]}
			return payloadDataGet, nil
		}
		break

	case "change.color":
		var payloadChangeColor PayloadChangeColor
		if err = json.Unmarshal(m.Payload, &payloadChangeColor); err == nil {
			device, ok := deviceList[payloadChangeColor.Device.Id]

			if ok {
				colorData := strip.StripColorData{
					Color:      payloadChangeColor.Color.Color,
					Brightness: payloadChangeColor.Color.Brightness,
					Power:      payloadChangeColor.Color.Power,
					Start:      0,
					End:        payloadChangeColor.Device.NumLeds}
				err = device.Change(colorData)
				if err != nil {
					return nil, err
				}

				return "changed", nil
			}
		}
		return "not changed", nil
		break
	}
	return nil, err
}
