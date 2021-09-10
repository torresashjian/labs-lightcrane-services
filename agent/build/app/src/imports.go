package main

import (
	_ "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/exec"
	_ "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/mapping"
	_ "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/log"
	_ "github.com/SteveNY-Tibco/labs-lightcrane-contrib/function/f1"
	_ "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/httpclient"
	_ "github.com/project-flogo/flow"
	_ "github.com/project-flogo/legacybridge"
	_ "github.com/project-flogo/flow/activity/subflow"
	_ "github.com/project-flogo/contrib/activity/actreturn"
	_ "github.com/project-flogo/contrib/trigger/timer"
	_ "github.com/project-flogo/contrib/function/coerce"
	_ "github.com/project-flogo/contrib/function/string"
	_ "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/jsondeserializer"
	_ "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/jsonserializer"
)
