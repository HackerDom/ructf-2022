package game

type Status byte

const (
	STATUS_UNFINISHED Status = 0
	STATUS_WIN        Status = iota
	STATUS_LOSE       Status = iota
)
