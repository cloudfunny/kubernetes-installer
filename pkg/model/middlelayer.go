package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Before cluster object create
func (c *Cluster) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid generating
	c.ID = uuid.NewString()
	return
}

// Before cluster object create
func (c *Cluster) AfterCreate(tx *gorm.DB) (err error) {
	// uuid generating
	event := &Event{
		Opt: "create",
		Obj: c,
	}
	WatchClusterChan <- event
	return
}

// Before cluster object create
func (n *Node) BeforeCreate(tx *gorm.DB) (err error) {
	// uuid generating
	n.ID = uuid.NewString()
	return
}

// Before cluster object create
func (n *Node) AfterCreate(tx *gorm.DB) (err error) {
	// uuid generating
	event := &Event{
		Opt: "create",
		Obj: n,
	}
	WatchNodeChan <- event
	return
}
