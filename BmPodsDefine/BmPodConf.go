package BmPodsDefine

type Conf struct {
	Model string
	Resource string
	Storage string
	Relationships struct{
		One2one []map[string]string
		One2many []map[string]string
	}
}
