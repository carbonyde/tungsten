package tungsten

import (
	"github.com/evanw/esbuild/pkg/api"
)

func Build(entries []string, variables map[string]string) string {
	treeshake := api.TreeShakingFalse

	if Env.Watch {
		treeshake = api.TreeShakingFalse
	} else {
		treeshake = api.TreeShakingTrue
	}

	result := api.Build(api.BuildOptions{
		TreeShaking:       treeshake,
		Bundle:            !Env.Watch,
		MinifyWhitespace:  !Env.Watch,
		MinifyIdentifiers: !Env.Watch,
		MinifySyntax:      !Env.Watch,
		Format:            api.FormatESModule,
		Target:            api.ES2020,
		Platform:          api.PlatformBrowser,
		Write:             false,
		EntryPoints:       entries,
		Define:            variables,
		Engines: []api.Engine{
			{Name: api.EngineChrome, Version: "103"},
			{Name: api.EngineFirefox, Version: "115"},
			{Name: api.EngineSafari, Version: "11"},
			{Name: api.EngineEdge, Version: "117"},
		},
	})

	content := result.OutputFiles[0]

	return string(content.Contents[:])
}
