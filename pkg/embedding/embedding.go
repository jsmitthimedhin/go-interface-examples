package embedding

type (
	walkable interface {
		walk()
	}
	bahtable interface {
		requiresBath() bool
	}
	talkable interface {
		walkable
		bahtable
	}
)
