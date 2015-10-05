package command

import (
	"log"
	"os"
	"path"
	"regexp"

	"github.com/codegangsta/cli"
	"github.com/hoisie/mustache"

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

	ctx := map[string]string{"projectname": projectName}
	populateDir(projectName, template, ctx)
}

func populateDir(dir string, template string, context map[string]string) {
	tpldir := path.Join(config.GetRepoPath(), template)
	files, err := getDirContents(tpldir)
	if err != nil {
		log.Fatalf("Could not get contents of template %s: %s", template, err)
	}
	for index := 0; index < len(files); index++ {
		fn := files[index]
		content := mustache.RenderFile(path.Join(tpldir, fn), context)
		f, err := os.Create(path.Join(dir, fn))
		if err != nil {
			log.Fatalf("Could not create %s: %s", fn, err)
		}
		f.WriteString(content)
		f.Close()
	}
}

func getDirContents(path string) (contents []string, err error) {
	d, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open %s: %s", path, err)
	}
	return d.Readdirnames(0)
}

func getTemplates() []string {
	names, err := getDirContents(config.GetRepoPath())
	if err != nil {
		log.Fatalf("Could not read templates: %s", err)
	}
	var templates []string
	ignore := regexp.MustCompile("^.git$")
	for idx := 0; idx < len(names); idx++ {
		if !ignore.MatchString(names[idx]) {
			templates = append(templates, names[idx])
		}
	}
	return templates
}

func makeChoiceMap(templates []string) map[string]string {
	choices := make(map[string]string)
	for index := 0; index < len(templates); index++ {
		choices[templates[index]] = "A fine templates"
	}
	return choices
}
