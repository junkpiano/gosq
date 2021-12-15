package gosq

type Option func(map[string]string)

func SetBranch(branch string) Option {
	return func(args map[string]string) {
		args["branch"] = branch
	}
}

func SetProject(project string) Option {
	return func(args map[string]string) {
		args["project"] = project
	}
}
