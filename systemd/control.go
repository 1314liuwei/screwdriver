package systemd

import (
	"context"
	"errors"
	"strings"

	"github.com/godbus/dbus/v5"
	"github.com/gogf/gf/v2/util/gconv"
)

type Systemd struct {
	conn *dbus.Conn
	obj  dbus.BusObject
}

func (c *Systemd) Connect() error {
	bus, err := dbus.SystemBus()
	if err != nil {
		return err
	}
	c.conn = bus
	c.obj = bus.Object(Systemd1Dest, Systemd1ObjectPath)
	return nil
}

func (c *Systemd) Close() error {
	return c.conn.Close()
}

func (c *Systemd) ListService(ctx context.Context) ([]Service, error) {
	reply := c.obj.CallWithContext(ctx, ListServicesMethod, 0, []string{}, []string{"*.service"})
	if reply.Err != nil {
		return nil, reply.Err
	}

	res, ok := reply.Body[0].([][]interface{})
	if ok {
		var out []Service
		for i := 0; i < len(res); i++ {
			out = append(out, Service{
				Name:        gconv.String(res[i][0]),
				Description: gconv.String(res[i][1]),
				LoadState:   gconv.String(res[i][2]),
				ActiveState: gconv.String(res[i][3]),
				SubState:    gconv.String(res[i][4]),
			})
		}
		return out, nil
	} else {
		return nil, errors.New("assert fail")
	}
}

func (c *Systemd) GetService(ctx context.Context, name string) (Service, error) {
	if strings.Contains(name, ".service") {
		name += ".service"
	}
	reply := c.obj.CallWithContext(ctx, ListServicesMethod, 0, []string{name}, []string{"*.service"})
	if reply.Err != nil {
		return Service{}, nil
	}

	res, ok := reply.Body[0].([][]interface{})
	if ok {
		var out Service
		out = Service{
			Name:        gconv.String(res[0][0]),
			Description: gconv.String(res[0][1]),
			LoadState:   gconv.String(res[0][2]),
			ActiveState: gconv.String(res[0][3]),
			SubState:    gconv.String(res[0][4]),
		}
		return out, nil
	} else {
		return Service{}, errors.New("assert fail")
	}
}

func (c *Systemd) StartService(ctx context.Context, s Service, mode ...Mode) error {
	var m Mode
	if len(mode) == 0 {
		m = Replace
	} else {
		m = mode[0]
	}

	res := c.obj.CallWithContext(ctx, StartServicesMethod, 0, s.Name, m.String())
	if res.Err != nil {
		return res.Err
	}
	return nil
}

func (c *Systemd) StopService(ctx context.Context, s Service, mode ...Mode) error {
	var m Mode
	if len(mode) == 0 {
		m = Replace
	} else {
		m = mode[0]
	}

	res := c.obj.CallWithContext(ctx, StopServicesMethod, 0, s.Name, m.String())
	if res.Err != nil {
		return res.Err
	}
	return nil
}

func (c *Systemd) RestartService(ctx context.Context, s Service, mode ...Mode) error {
	var m Mode
	if len(mode) == 0 {
		m = Replace
	} else {
		m = mode[0]
	}

	res := c.obj.CallWithContext(ctx, RestartServicesMethod, 0, s.Name, m.String())
	if res.Err != nil {
		return res.Err
	}
	return nil
}

func (c *Systemd) ReloadService(ctx context.Context, s Service, mode ...Mode) error {
	var m Mode
	if len(mode) == 0 {
		m = Replace
	} else {
		m = mode[0]
	}

	res := c.obj.CallWithContext(ctx, ReloadServicesMethod, 0, s.Name, m.String())
	if res.Err != nil {
		return res.Err
	}
	return nil
}
