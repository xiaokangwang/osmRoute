package mapctx

type Node interface {
	FindConnection(ConnectionSpec) []Connection
}

type Connection interface {
	From() Node
	To() Node
	GetCost() float64
}

type ConnectionSpec interface {
	CanWalk() bool
	CanDrive() bool
	CanPublicTransport() bool
}
