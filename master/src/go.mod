module traverse-master

go 1.22

replace traverse/pkg/utils => ./pkg/utils

replace traverse/pkg/logging => ./pkg/logging

replace traverse/pkg/proxy => ./pkg/proxy

require (
	traverse/pkg/logging v0.0.0-00010101000000-000000000000 // indirect
	traverse/pkg/proxy v0.0.0-00010101000000-000000000000 // indirect
	traverse/pkg/utils v0.0.0-00010101000000-000000000000 // indirect
)
