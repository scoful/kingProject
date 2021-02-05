package cron

type AddCron struct {
    CronName   string `v:"required#请输入定时任务名"`
    TargetTime string `v:"required#请输入目标时间"`
    Content    string `v:"required#请输入内容"`
}

type DeleteOrSearchCron struct {
    CronName string `v:"required#请输入定时任务名"`
}
