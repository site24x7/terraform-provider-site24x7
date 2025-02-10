package api

const (
	// Status constants denotes the status of the monitors.
	Down               Status = 0
	Up                 Status = 1
	Trouble            Status = 2
	Critical           Status = 3
	Suspended          Status = 5
	Maintenance        Status = 7
	Discovery          Status = 9
	ConfigurationError Status = 10

	// ResourceType constants denotes the resource type of the third party integration.
	AllMonitors ResourceType = 0
	Monitor     ResourceType = 2
	Tags        ResourceType = 3

	// Type of the Site24x7 resource.
	URL          MonitorType = "URL"
	HOMEPAGE     MonitorType = "HOMEPAGE"
	SSL_CERT     MonitorType = "SSL_CERT"
	RESTAPI      MonitorType = "RESTAPI"
	RESTAPISEQ   MonitorType = "RESTAPISEQ"
	AMAZON       MonitorType = "AMAZON"
	SERVER       MonitorType = "SERVER"
	CRON         MonitorType = "CRON"
	HEARTBEAT    MonitorType = "HEARTBEAT"
	DNS          MonitorType = "DNS"
	DOMAINEXPIRY MonitorType = "DOMAINEXPIRY"
	REALBROWSER  MonitorType = "REALBROWSER"
	FTP          MonitorType = "FTP"
	ISP          MonitorType = "ISP"
	PORT         MonitorType = "PORT"
	PING         MonitorType = "PING"
	SOAP         MonitorType = "SOAP"
	GCP          MonitorType = "GCP"
)
