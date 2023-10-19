package sql_queries

import _ "embed"

//go:embed count.sql
var CountSql string

//go:embed average.sql
var AverageSql string

//go:embed weekly_deployment.sql
var WeeklyDeploymentSql string

//go:embed monthly_deployment.sql
var MonthlyDeploymentSql string

//go:embed quarterly_deployment.sql
var QuarterlyDeploymentSql string
