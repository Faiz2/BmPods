package BmPodsDefine

type Conf struct {
	Resource string `yaml: "resource"`
	Collection string `yaml: "collection"`
	Relationships struct{
		One2one []string
		One2many []string
	}
}
