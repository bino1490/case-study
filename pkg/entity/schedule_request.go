package entity

type ScheduleRequest struct {
	ChannelId     string `json:"channel_id,omitempty"`
	ScheduleId    string `json:"schedule_id,omitempty"`
	StartEpoch    string `json:"start_epoch,omitempty"`
	Source        string `json:"source,omitempty"`
	SourceChannel string `json:"source_channel,omitempty"`
	NextPrg_Epoch string `json:"next_program_start_epoch,omitempty"`
	Encoding      string `json:"encoding,omitempty"`
	Pid           PID    `json:"pid,omitempty"`
}

type PID struct {
	VPid string     `json:"videoPid,omitempty"`
	APid []AudioPid `json:"audioPid,omitempty"`
}

type AudioPid struct {
	Lang string `json:"lang,omitempty"`
	Pid  string `json:"pid,omitempty"`
}
