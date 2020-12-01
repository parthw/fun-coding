package oldgrpc

// OldInt interface
type OldInt interface {
	// If I put Something() here, then I will get error
	Myfunction()
}

// Register Function
func Register(o OldInt) {}
