package strip

import (
	"encoding/hex"
	"log"
	"time"

	serial "go.bug.st/serial.v1"
	"go.bug.st/serial.v1/enumerator"
)

type StripRGBData struct {
	Start      byte
	End        byte
	Red        byte
	Green      byte
	Blue       byte
	Brightness byte
}

func (s StripRGBData) AsByteArray() []byte {
	data := make([]byte, 7)
	data[0] = 'c'
	data[1] = s.Start
	data[2] = s.End
	data[3] = s.Red
	data[4] = s.Green
	data[5] = s.Blue
	data[6] = s.Brightness

	return data
}

type DeviceData struct {
	Id           int         `json:"id"`
	Vendor       string      `json:"vendor"`
	Product      string      `json:"product"`
	SerialNumber string      `json:"serialnumber"`
	Port         string      `json:"port"`
	Baud         int         `json:"baud"`
	NumLeds      int         `json:"numleds"`
	Handle       serial.Port `json:"-"`
	Connected    bool        `json:"Connected"`
}

type StripColorData struct {
	Color      string `json:"color"`
	Brightness int    `json:"brightness"`
	Power      bool   `json:"power"`
	Start      int    `json:"start"`
	End        int    `json:"end"`
}

func FindConnectedDevices(connect bool) (deviceList map[int]DeviceData, err error) {
	deviceList = map[int]DeviceData{}

	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return nil, err
	}
	if len(ports) == 0 {
		return nil, nil
	}
	for _, comport := range ports {
		var device DeviceData

		if comport.IsUSB {
			device.Port = comport.Name
			device.Vendor = comport.VID
			device.Product = comport.PID
			device.SerialNumber = comport.SerialNumber
			device.Baud = 9600

			if !connect {
				err = device.Open()
			} else {
				device.Connected, err = device.Connect()
			}

			if err == nil {
				time.Sleep(2 * time.Second)
				device.Handle.SetDTR(true)
				device.Handle.Write([]byte("i"))
				response, _ := device.Recieve(7)
				if validateDeviceInfo(response) {
					device.Id = int(response[2])
					device.NumLeds = int(response[6])
					deviceList[device.Id] = device
				}
				if !connect {
					device.Handle.Close()
				}
			}
		}
	}
	return deviceList, nil
}

func (dd *DeviceData) Open() (err error) {
	mode := &serial.Mode{
		BaudRate: dd.Baud,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	dd.Handle, err = serial.Open(dd.Port, mode)
	return err
}

func (dd *DeviceData) Connect() (state bool, err error) {
	err = dd.Open()

	if err == nil {
		time.Sleep(2 * time.Second)
		dd.Handle.SetDTR(true)
		dd.Handle.Write([]byte("i"))
		response, _ := dd.Recieve(7)
		dd.Connected = validateDeviceInfo(response)
		return dd.Connected, err
	}

	return false, err
}

func (dd *DeviceData) Recieve(count int) (response []byte, err error) {
	buff := make([]byte, count)
	var c int
	var n int
	c = 0

	for {
		n, err = dd.Handle.Read(buff)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		response = append(response[:], buff[:n]...)

		c = c + n
		if c >= count {
			break
		}
	}

	return response, nil
}

func validateDeviceInfo(response []byte) bool {

	if response[0] != 43 {
		return false
	}
	if response[1] != 51 {
		return false
	}
	if response[2] <= 100 {
		return false
	}
	if response[3] != 1 {
		return false
	}
	if response[4] != 0 {
		return false
	}
	if response[5] != 0 {
		return false
	}

	return true
}

func (dd *DeviceData) Change(data StripColorData) (err error) {
	var brightness byte
	if !data.Power {
		brightness = 0x00
	} else {
		brightness = byte(data.Brightness)
	}

	color := hexColorToByteArray(data.Color)
	stripdata := StripRGBData{byte(float64(data.Start)), byte(float64(data.End)), color[0], color[1], color[2], brightness}
	_, err = dd.Handle.Write(stripdata.AsByteArray())
	return err
}

func hexToByte(h string) byte {
	data, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}
	return data[0]
}

func hexColorToByteArray(hex string) []byte {
	data := make([]byte, 3)
	data[0] = hexToByte(hex[1:3])
	data[1] = hexToByte(hex[3:5])
	data[2] = hexToByte(hex[5:7])
	return data
}
