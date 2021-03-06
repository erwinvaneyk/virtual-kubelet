package daemon

import (
	// Importing packages here only to make sure their init gets called and
	// therefore they register themselves to the logdriver factory.
	_ "github.com/hyperhq/hypercli/daemon/logger/awslogs"
	_ "github.com/hyperhq/hypercli/daemon/logger/fluentd"
	_ "github.com/hyperhq/hypercli/daemon/logger/gelf"
	_ "github.com/hyperhq/hypercli/daemon/logger/journald"
	_ "github.com/hyperhq/hypercli/daemon/logger/jsonfilelog"
	_ "github.com/hyperhq/hypercli/daemon/logger/splunk"
	_ "github.com/hyperhq/hypercli/daemon/logger/syslog"
)
