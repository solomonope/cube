package node

type Node struct {
	Name             string
	Ip               string
	Memory           int
	MemoryAllocatted int
	Disk             int
	DiskAllocated    int
	TaskCount        int
	Cores            int
	Role             string
}
