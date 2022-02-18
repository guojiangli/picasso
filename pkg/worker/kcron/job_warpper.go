package kcron

import "github.com/robfig/cron/v3"


//JobWarpper is a job warpper
type JobWrapper struct {
	ID      cron.EntryID
	Kcron   *Kcron
	Name    string
	CronStr string
	Func    func()
	Job     cron.Job
}

//Run is run job
func (job JobWrapper) Run() {
	//如果该任务分配给了这个节点 则允许执行
	if job.Kcron.allowThisNodeRun(job.Name) {
		if job.Func != nil {
			job.Func()
		}
		if job.Job != nil {
			job.Job.Run()
		}
	}
}
