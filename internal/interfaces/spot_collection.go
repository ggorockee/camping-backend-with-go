package interfaces

type Spot interface {
	GetId() uint
}

// Spot을 인터페이스로 정의하여 구체적인 구현에 의존하지 않게 함
// 이러한 변경으로 interfaces 패키지는 entities 패키지에 의존하지 않게 되어
// 순환 참조 문제를 해결할 수 있.
// 구체적인 Spot 구현은 다른 패키지에서 이 인터페이스를 구현
type SpotCollection interface {
	Filter(id int) SpotCollection
	Exists() bool
	Remove(spot Spot)
	Add(spot Spot)
	ToSlice() []Spot
}

type SpotCollectionFactory func([]Spot) SpotCollection
