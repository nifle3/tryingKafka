package cfg

type Cfg struct {
	SMTPCfg
	KafkaCfg
}

type SMTPCfg struct {
}

type KafkaCfg struct {
}

func New() Cfg {
	return Cfg{}
}
