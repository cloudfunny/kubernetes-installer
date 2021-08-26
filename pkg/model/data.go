package model

import (
	"gorm.io/gorm"
)

type Cluster struct {
	ID         string `gorm:"primaryKey"`
	Name       string
	Version    string
	PodNetCIDR string
	Nodes      []Node
	Registry   string
	gorm.Model
}

type Node struct {
	ID        string `gorm:"primaryKey"`
	Hostname  string
	Role      string
	ClusterID string
	gorm.Model
}
