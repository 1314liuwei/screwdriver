package systemd

type Service struct {
	Name        string
	Description string
	LoadState   string
	ActiveState string
	SubState    string
}
