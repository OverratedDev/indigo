package process

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/struCoder/pidusage"
	"github.com/vereas/indigo/utils"
)

type ProcStatus struct {
	Status    string
	Restarts  int
	StartTime int64
	Uptime    string
	Sys       *pidusage.SysInfo
}

func (proc_status *ProcStatus) SetStatus(status string) {
	proc_status.Status = status
}

func (proc_status *ProcStatus) AddRestart() {
	proc_status.Restarts++
}

func (proc_status *ProcStatus) InitUptime() {
	proc_status.StartTime = time.Now().Unix()
}

func (proc_status *ProcStatus) SetUptime() {
	proc_status.Uptime = utils.FormatUptime(proc_status.StartTime, time.Now().Unix())
}

func (proc_status *ProcStatus) ResetUptime() {
	proc_status.Uptime = "0s"
}

func (proc_status *ProcStatus) SetSysInfo(pid int) {
	var err error
	proc_status.Sys, err = pidusage.GetStat(pid)
	if err != nil {
		log.Error(err)
	}
}