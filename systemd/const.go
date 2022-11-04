package systemd

const (
	GetPropertyMethod    = "org.freedesktop.DBus.Properties.Get"
	GetAllPropertyMethod = "org.freedesktop.DBus.Properties.GetAll"
	IntrospectMethod     = "org.freedesktop.DBus.Introspectable.Introspect"

	Systemd1DBusName              = "systemd1"
	Systemd1Dest                  = "org.freedesktop.systemd1"
	Systemd1ObjectPath            = "/org/freedesktop/systemd1"
	Systemd1ObjectMangerInterface = "org.freedesktop.systemd1.Manager"
	ListServicesMethod            = "org.freedesktop.systemd1.Manager.ListUnitsByPatterns"
	StartServicesMethod           = "org.freedesktop.systemd1.Manager.StartUnit"
	StopServicesMethod            = "org.freedesktop.systemd1.Manager.StopUnit"
	RestartServicesMethod         = "org.freedesktop.systemd1.Manager.RestartUnit"
	ReloadServicesMethod          = "org.freedesktop.systemd1.Manager.ReloadUnit"
)

type Mode int

const (
	Replace Mode = iota
	Fail
	Isolate
	IgnoreDependencies
	IgnoreRequirements
)

func (m Mode) String() string {
	return []string{"replace", "fail", "isolate", "ignore-dependencies", "ignore-requirements"}[m]
}
