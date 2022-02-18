package kcron

import (
	"sync"
	"time"

	"github.com/guojiangli/picasso/pkg/klog/baselogger"
	"github.com/guojiangli/picasso/pkg/worker/kcron/consistenthash"
	"github.com/guojiangli/picasso/pkg/worker/kcron/driver"
)

const (
	defaultReplicas         = 10
	defaultIntervalDuration = 30 * time.Second
)

// NodePool is a node pool
type NodePool struct {
	workerName string
	NodeID     string
	interval   time.Duration
	mu         sync.Mutex
	nodes      *consistenthash.Map
	logger     baselogger.KcronLogger
	Driver     driver.Driver
	opts       poolOption
	kcron      *Kcron
}

// poolOption is a pool option
type poolOption struct {
	Replicas int
	HashFn   consistenthash.Hash
}

func (k *Kcron) newNodePool(workerName string, driver driver.Driver, kcron *Kcron) (*NodePool, error) {
	nodePool := new(NodePool)
	nodePool.interval = defaultIntervalDuration
	nodePool.Driver = driver
	err := nodePool.Driver.Ping()
	if err != nil {
		k.logger.Error(err, "picasso.kcron")
		return nil, err
	}
	nodePool.workerName = workerName
	nodePool.kcron = kcron
	nodePool.logger = k.logger
	option := poolOption{
		Replicas: defaultReplicas,
	}
	nodePool.opts = option
	return nodePool, nil
}

func (np *NodePool) InitPoolGrabService() {
	np.NodeID = np.Driver.RegisterServiceNode(np.workerName, np.interval)
	go np.Driver.DoHeartBeat(np.NodeID, np.interval)
}

func (np *NodePool) updatePool() {
	np.mu.Lock()
	defer np.mu.Unlock()
	nodes, err := np.Driver.GetServiceNodeList(np.workerName)
	if err != nil {
		np.logger.Error(err, "picasso.kcron", "update redis Pool lock fail %s")
		return
	} else if nodes == nil || len(nodes) == 0 {
		go np.InitPoolGrabService()
		return
	}
	np.logger.Info("服务节点列表", "picasso.kcron", nodes)
	np.nodes = consistenthash.New(np.opts.Replicas, np.opts.HashFn)
	for _, node := range nodes {
		np.nodes.Add(node)
	}
}

func (np *NodePool) tickerUpdatePool() {
	tickers := time.NewTicker(np.interval)
	for range tickers.C {
		if np.kcron.isRun {
			np.updatePool()
		}
	}
}

// PickNodeByJobName : 使用一致性hash算法根据任务名获取一个执行节点
func (np *NodePool) PickNodeByJobName(jobName string) string {
	np.mu.Lock()
	defer np.mu.Unlock()
	if np.nodes == nil || np.nodes.IsEmpty() {
		return ""
	}
	return np.nodes.Get(jobName)
}
