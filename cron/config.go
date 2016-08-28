package cron

var config = struct {
	QList []string
}{}

func init() {
	config.QList = []string{"投資"}
}
