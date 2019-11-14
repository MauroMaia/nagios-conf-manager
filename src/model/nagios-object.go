package model

type NagiosObject interface {
	String() string
	FormatToNagiosCFG() (string, error)
}
