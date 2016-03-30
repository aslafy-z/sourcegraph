package metricutil

import (
	"fmt"
	"os"
)

var srcEnvWhitelist = []string{
	"SG_API_DEFAULT_CACHE_MAX_AGE",
	"SG_APP_RAVEN_DSN",
	"SG_BUILD_LOG_DIR",
	"SG_CHECKING",
	"SG_DB_TEST_POOL",
	"SG_DIR",
	"SG_ELASTICSEARCH_URL",
	"SG_ENABLE_GITHUB_CLONE_PROXY",
	"SG_ENABLE_HSTS",
	"SG_ERROR",
	"SG_FEATURE_DISCUSSIONS",
	"SG_FEATURE_SEARCHNEXT",
	"SG_FORCE_HTTPS",
	"SG_GRAPHSTORE_ROOT",
	"SG_HTTP_TRACE",
	"SG_KIBANA_CLIENT_URL",
	"SG_KIBANA_USER_URL",
	"SG_LOG_FILE",
	"SG_LOG_URL_FOR_TAG",
	"SG_NAV_MSG",
	"SG_NOTICE",
	"SG_NO_SPACE_EXPECTED",
	"SG_NO_UPDATE_WATCHER",
	"SG_NUM_WORKERS",
	"SG_PREFIX",
	"SG_RAVEN_DSN",
	"SG_REQUIRE_SECRETS",
	"SG_RESULT",
	"SG_SPACE_EXPECTED",
	"SG_STRICT_HOSTNAME",
	"SG_SYSLOG_HOST",
	"SG_TRACEGUIDE_SERVICE_HOST",
	"SG_URL",
	"SG_USE_CSP",
	"SG_USE_PAPERTRAIL",
	"SRC_ALPHA",
	"SRC_ALPHA_SATURATE",
	"SRC_AMAZON_EC2",
	"SRC_COLOR",
	"SRC_COMMENT",
	"SRC_DIGITAL_OCEAN",
	"SRC_ENDPOINT",
	"SRC_FILES",
	"SRC_GOOGLE_COMPUTE_ENGINE",
	"SRC_HOSTNAME",
	"SRC_HREF_DQ",
	"SRC_HREF_IN_XML",
	"SRC_HREF_SQ",
	"SRC_LANGUAGE_GO",
	"SRC_LANGUAGE_JAVA",
	"SRC_NOT_SUPPORTED",
	"SRC_RGB",
	"SRC_TOKEN",
	"SRC_URL",
}

// getWhitelistedEnvironment returns the Sourcegraph env variables
// that are set in the current environment. The returned
// slice contains strings in the format "key=value".
// Env vars containing secret information, eg. SRC_ID_KEY_DATA,
// are not returned.
func getWhitelistedEnvironment() []string {
	var vars []string
	for _, key := range srcEnvWhitelist {
		if val := os.Getenv(key); val != "" {
			vars = append(vars, fmt.Sprintf("%s=%s", key, val))
		}
	}
	return vars
}