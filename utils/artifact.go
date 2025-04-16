package utils

import (
	"fmt"
	"strings"
	"time"
)

func GenerateArtifactName(project, env, version, build string) string {
	timestamp := time.Now().Format("060102") // YYMMDD
	return fmt.Sprintf("%s_%s_v%s+%s_%s", strings.ToLower(project), strings.ToLower(env), version, build, timestamp)
}
