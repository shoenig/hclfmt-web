package format

import "fmt"

func Markdown(kind, content string) string {
	switch kind {
	case "github":
		return mdGithub(content)
	case "jira":
		return mdJira(content)
	default:
		return content
	}
}

func mdGithub(content string) string {
	return fmt.Sprintf("```hcl\n%s\n```", content)
}

func mdJira(content string) string {
	return fmt.Sprintf("{code}\n%s\n{code}", content)
}
