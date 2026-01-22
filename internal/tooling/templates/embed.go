package templates

import "embed"

//go:embed templates/**
var embeddedTemplates embed.FS
