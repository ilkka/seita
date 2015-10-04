package command

import (
	"log"
	"os"

	"github.com/codegangsta/cli"

	"github.com/ilkka/seita/config"
	"github.com/ilkka/seita/ui"
)

// Make is the command for creating a project based on an existing
// skeleton.
func Make(c *cli.Context) {
	projectName := c.Args().First()

	templates := getTemplates()
	choices := makeChoiceMap(templates)
	template := ui.Choose("Which template?", choices)

	os.Mkdir(projectName, 0755)
	ui.Printf("Would now populate %s with %s\n", projectName, template)
	// populate with skeleton, replacing name
}

func getTemplates() []string {
	d, err := os.Open(config.GetRepoPath())
	if err != nil {
		log.Fatalf("Could not open repo: %s", err)
	}
	names, err := d.Readdirnames(0)
	if err != nil {
		log.Fatalf("Could not read templates: %s", err)
	}
	return names
}

func makeChoiceMap(templates []string) map[string]string {
	choices := make(map[string]string)
	for index := 0; index < len(templates); index++ {
		choices[templates[index]] = "A fine templates"
	}
	return choices
}
