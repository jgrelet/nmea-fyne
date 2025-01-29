package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tarm/serial"
)

type GPSData struct {
	Latitude  string
	Longitude string
	Time      string
	Date      string
}

func main() {
	gpsChannel := make(chan GPSData)

    go readSerialPort("COM1", gpsChannel)
    //go readEthernet("localhost:12345", gpsChannel)

	a := app.New()
	w := a.NewWindow("GPS Data")

	label := widget.NewLabel("Waiting for GPS data...")
	w.SetContent(container.NewVBox(label))
	w.Resize(fyne.NewSize(400, 200))

	go func() {
		for data := range gpsChannel {
			label.SetText(fmt.Sprintf("Lat: %s, Lon: %s, Time: %s, Date: %s",
				data.Latitude, data.Longitude, data.Time, data.Date))
		}
	}()

	w.ShowAndRun()
}

func readSerialPort(portName string, gpsChannel chan GPSData) {
    c := &serial.Config{Name: portName, Baud: 9600}
    s, err := serial.OpenPort(c)
    if err != nil {
        log.Fatal(err)
    }
    defer s.Close()

	reader := bufio.NewReader(s)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading from serial port:", err)
			continue
		}
		processNMEA(line, gpsChannel)
	}
}

func readEthernet(address string, gpsChannel chan GPSData) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading from ethernet connection:", err)
			continue
		}
		processNMEA(line, gpsChannel)
	}
}

func processNMEA(line string, gpsChannel chan GPSData) {
	if strings.HasPrefix(line, "$GPRMC") {
		fields := strings.Split(line, ",")
		if len(fields) < 12 {
			return
		}
		latitude := fields[3] + " " + fields[4]
		longitude := fields[5] + " " + fields[6]
		time := fields[1]
		date := fields[9]

		gpsData := GPSData{
			Latitude:  latitude,
			Longitude: longitude,
			Time:      time,
			Date:      date,
		}
		gpsChannel <- gpsData
	}
}
