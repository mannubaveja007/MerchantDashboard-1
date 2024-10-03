# Set the start time to the epoch (1970-01-01T00:00:00Z)
$startTime = "1970-01-01T00:00:00Z"

# Set the end time to the end of yesterday
$endTime = (Get-Date).AddDays(-1).Date.AddHours(23).AddMinutes(59).AddSeconds(59).ToString("yyyy-MM-ddTHH:mm:ssZ")

# Call the AWS CLI command
aws cloudwatch get-metric-statistics `
    --namespace AWS/DynamoDB `
    --metric-name ConsumedReadCapacityUnits `
    --dimensions Name=TableName,Value=products `
    --start-time $startTime `
    --end-time $endTime `
    --period 86400 `
    --statistics Sum
