package kcron

import (
	"sync"

	"picasso/pkg/klog/baselogger"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

// Kcron is main struct
type Kcron struct {
	jobs     map[string]*JobWrapper
	mu       sync.RWMutex
	cr       *cron.Cron
	nodePool *NodePool
	isRun    bool
	logger   baselogger.KcronLogger
}

// NewKcron create a Kcron
func New(option ...*Option) (k *Kcron, err error) {
	kcronOption := defaultOption().MergeOption(option...)
	if kcronOption.WorkerName == "" {
		err = errors.New("WorkerName不能为空")
		return nil, err
	}
	if kcronOption.Driver == nil {
		err = errors.New("driver不能为空")
		return nil, err
	}
	k = &Kcron{
		jobs:   make(map[string]*JobWrapper),
		mu:     sync.RWMutex{},
		cr:     cron.New(cron.WithLogger(kcronOption.Logger)),
		isRun:  false,
		logger: kcronOption.Logger,
	}
	k.nodePool, err = k.newNodePool(kcronOption.WorkerName, kcronOption.Driver, k)
	if err != nil {
		return nil, err
	}
	return k, nil
}

// AddJob  add a job
func (d *Kcron) AddJob(jobName, cronStr string, job cron.Job) (err error) {
	return d.addJob(jobName, cronStr, nil, job)
}

// AddFunc add a cron func
func (d *Kcron) AddFunc(jobName, cronStr string, cmd func()) (err error) {
	return d.addJob(jobName, cronStr, cmd, nil)
}

func (d *Kcron) addJob(jobName, cronStr string, cmd func(), job cron.Job) (err error) {
	if _, ok := d.jobs[jobName]; ok {
		return errors.New("jobName already exist")
	}
	innerJob := JobWrapper{
		Name:    jobName,
		CronStr: cronStr,
		Func:    cmd,
		Job:     job,
		Kcron:   d,
	}
	entryID, err := d.cr.AddJob(cronStr, innerJob)
	if err != nil {
		return err
	}
	innerJob.ID = entryID
	d.jobs[jobName] = &innerJob
	return err
}

// Remove Job
func (d *Kcron) Remove(jobName string) {
	if job, ok := d.jobs[jobName]; ok {
		delete(d.jobs, jobName)
		d.cr.Remove(job.ID)
	}
}

func (d *Kcron) allowThisNodeRun(jobName string) bool {
	return d.nodePool.NodeID == d.nodePool.PickNodeByJobName(jobName)
}

// Start start job
func (d *Kcron) Run() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.nodePool.NodeID == "" {
		d.nodePool.InitPoolGrabService()
	}
	d.isRun = true
	// 更新一次pool,再定时更新.
	d.nodePool.updatePool()
	go d.nodePool.tickerUpdatePool()
	d.cr.Run()
	return nil
}

// Stop stop job
func (d *Kcron) Stop() error {
	d.isRun = false
	d.cr.Stop()
	return nil
}
