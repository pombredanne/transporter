package adaptor

import (
	_ "github.com/compose/transporter/pkg/message/adaptor/elasticsearch"
	_ "github.com/compose/transporter/pkg/message/adaptor/file"
	_ "github.com/compose/transporter/pkg/message/adaptor/influxdb"
	_ "github.com/compose/transporter/pkg/message/adaptor/mongodb"
	_ "github.com/compose/transporter/pkg/message/adaptor/transporter"
)
