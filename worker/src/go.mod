module traverse-worker

go 1.22

replace traverse/pkg/utils => ./pkg/utils

replace traverse/pkg/logging => ./pkg/logging

replace traverse/pkg/proxy => ./pkg/proxy

replace traverse/pkg/secure_com => ./pkg/secure_com

require (
	traverse/pkg/logging v0.0.0-00010101000000-000000000000
	traverse/pkg/proxy v0.0.0-00010101000000-000000000000
	traverse/pkg/secure_com v0.0.0-00010101000000-000000000000
	traverse/pkg/utils v0.0.0-00010101000000-000000000000
)

require github.com/google/uuid v1.6.0 // indirect
