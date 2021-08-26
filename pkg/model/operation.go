package model

import "gorm.io/gorm"

func CreateCluster(name string, version string, podnetcidr string, registry string) *Cluster {
	newCluster := &Cluster{
		Name:       name,
		Version:    version,
		PodNetCIDR: podnetcidr,
		Registry:   registry,
	}
	result := db.Omit("Node").Create(newCluster)
	db.Save(newCluster)
	if result.Error != nil {
		panic("db operate failed, unable create item")
	}
	return newCluster
}

func RegisterNode(hostname string, role string, cid string) *Node {
	newNode := &Node{
		Hostname:  hostname,
		Role:      role,
		ClusterID: cid,
	}

	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(newNode).Error; err != nil {
			return err
		}
		return nil
	})

	return newNode
}

func QueryClusterByID(id string) *Cluster {
	cluster := &Cluster{}
	db.First(cluster, "id = ?", id)
	db.Preload("Nodes").Find(cluster)
	return cluster
}

func ListAllCluster() []Cluster {
	var clusters []Cluster
	db.Find(&clusters)
	db.Preload("Nodes").Find(&clusters)
	return clusters
}
