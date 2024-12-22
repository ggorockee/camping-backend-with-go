package common

type Serializer interface {
	Serialize() any
}

type Spot interface {
	GetId() string
}

type Review interface {
	GetId() string
}

type SpotDetailRes interface {
	GetId() string
}

type SpotTinyRes interface {
	GetId() string
}

type ReviewTinyRes interface {
	GetId() string
}

type UserTinyRes interface {
	GetId() string
}

type AmenityTinyRes interface {
	GetId() string
}

type CategoryTinyRes interface {
	GetId() string
}
