package secrets

import (
	// Register your backends here
	_ "github.com/ContainerSolutions/externalsecret-operator/backend/secrets/asm"
	_ "github.com/ContainerSolutions/externalsecret-operator/backend/secrets/dummy"
	_ "github.com/ContainerSolutions/externalsecret-operator/backend/secrets/onepassword"
)
