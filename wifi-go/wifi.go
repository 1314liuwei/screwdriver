package wifi

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
)

type WiFi interface {
	Close() error
	Scan() ([]Profile, error)
}

type wifi struct {
	ifname string
	conn   *net.UnixConn
}

func Open(ifname string) (WiFi, error) {
	// WPA  socket addr
	raddr, err := net.ResolveUnixAddr("unixgram", fmt.Sprintf("%s/%s", CtrlIfaceDir, ifname))
	if err != nil {
		return nil, err
	}

	// remove exists socket file
	laddrPath := fmt.Sprintf("%s/wifigo_%s", SockFileDir, ifname)
	err = os.RemoveAll(laddrPath)
	if err != nil {
		return nil, err
	}

	// local socket addr
	laddr, err := net.ResolveUnixAddr("unixgram", laddrPath)
	if err != nil {
		return nil, err
	}

	unix, err := net.DialUnix("unixgram", laddr, raddr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	_, err = unix.Write([]byte("PING"))
	if err != nil {
		return nil, err
	}

	buff := make([]byte, BuffSize)
	for i := 0; i < 5; i++ {
		n, err := unix.Read(buff)
		if err != nil {
			return nil, err
		}

		if strings.HasPrefix(string(buff[:n]), "PONG") {
			return &wifi{
				ifname: ifname,
				conn:   unix,
			}, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Connection to '%s' is broken!", CtrlIfaceDir))
}

func (w *wifi) Scan() ([]Profile, error) {
	_, err := w.execute("SCAN")
	if err != nil {
		return nil, err
	}

	reply, err := w.execute("SCAN_RESULTS")
	if err != nil {
		return nil, err
	}

	profiles := map[string]Profile{}
	result := strings.Split(string(reply[:len(reply)-1]), "\n")
	for i := 1; i < len(result); i++ {
		values := strings.Split(result[i], "\t")
		if len(values) < 5 {
			continue
		}

		var (
			ssid  = values[4]
			freqs []string
			akms  []string
		)

		freq := gconv.Int(values[1])
		if freq >= 2412 && freq <= 2484 {
			freqs = append(freqs, "2.4GHz")
		} else if freq >= 4915 && freq <= 5825 {
			freqs = append(freqs, "5GHz")
		}

		for _, v := range []string{WPA_PSK, WPA2_PSK, WPA_EAP, WPA2_EAP} {
			if strings.Contains(values[3], v) {
				akms = append(akms, v)
			}
		}

		if p, ok := profiles[ssid]; ok {
			freqs = append(freqs, p.Frequency...)
			freqs = gconv.Strings(set(gconv.Interfaces(freqs)))

			akms = append(akms, p.Akm...)
			akms = gconv.Strings(set(gconv.Interfaces(akms)))
			profiles[ssid] = Profile{
				SSID:      ssid,
				Frequency: freqs,
				Akm:       akms,
				RSSI:      max(p.RSSI, gconv.Int(values[2])),
			}
		} else {
			profiles[ssid] = Profile{
				SSID:      ssid,
				Frequency: freqs,
				Akm:       akms,
				RSSI:      gconv.Int(values[2]),
			}
		}
	}

	var data []Profile
	for _, profile := range profiles {
		data = append(data, profile)
	}
	return data, nil
}

func (w *wifi) Connect(p Profile) error {

	return nil
}

func (w *wifi) Networks() error {
	execute, err := w.execute("LIST_NETWORKS")
	if err != nil {
		return err
	}

	log.Println(execute)
	return nil
}

func (w *wifi) Close() error {
	return w.conn.Close()
}

func (w *wifi) execute(cmd string) ([]byte, error) {
	_, err := w.conn.Write([]byte(cmd))
	if err != nil {
		return nil, err
	}

	buff := make([]byte, BuffSize)
	n, err := w.conn.Read(buff)
	if err != nil {
		return nil, err
	}

	return buff[:n], nil
}
