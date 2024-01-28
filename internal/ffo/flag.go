package ffo

type Flag struct {
	Name    string
	Value   string
	Enabled bool
}

type FeatureFlags struct {
	Flags []Flag
}
