package useragentparser

import (
	"regexp"
	"strings"
)

var (
	BUILD_REGEX = `Build/([\w\.]+)`
	CRIOS_REGEX = `CriOS/([\w\.]+)`
	SAFARI_REGEX = `Safari/([\w\.]+)`
	WEBVIEW_REGEX = `^\w{2}-\w{2};\s*`
	MODEL_REGEX = `\b(Linux|Android)[^;]+; ([^/]+)/`
	ANDROID_VERSION_REGEX = `Android (\d+(?:\.\d+)*);`
)

func Parser(userAgent string) map[string]string {
	var version, model, build, crios, safari string
	
	androidMatches := regexp.MustCompile(ANDROID_VERSION_REGEX).FindStringSubmatch(userAgent);
	if len(androidMatches) >= 2 {
		version = androidMatches[1]
	}
  
	criosMatches := regexp.MustCompile(CRIOS_REGEX).FindStringSubmatch(userAgent)
	if len(criosMatches) >= 2 {
	  crios = criosMatches[1]  
	}
  
	safariMatches := regexp.MustCompile(SAFARI_REGEX).FindStringSubmatch(userAgent)
	if len(safariMatches) >= 2 {
	  safari = safariMatches[1]
	}
	
	buildMatches := regexp.MustCompile(BUILD_REGEX).FindStringSubmatch(userAgent)
	if len(buildMatches) >= 2 {
		build = buildMatches[1]
	}
  
	modelMatches := regexp.MustCompile(MODEL_REGEX).FindStringSubmatch(userAgent);
	if len(modelMatches) >= 2 {
		modelVersion := strings.ReplaceAll(modelMatches[2], "SAMSUNG ", "")
		modelVersion = strings.ReplaceAll(modelVersion, "wv ", "")
		modelVersion = strings.ReplaceAll(modelVersion, " wv", "")
		
		bracketIdx := strings.Index(modelVersion, ")")
		if bracketIdx != -1 {
			model = string(modelVersion[:bracketIdx])
		}

		buildIdx := strings.Index(modelVersion, " Build")
		if buildIdx != -1 {
			model = string(modelVersion[:buildIdx])
		}

		re := regexp.MustCompile(WEBVIEW_REGEX)
		model = re.ReplaceAllString(model, "")
		model = strings.ReplaceAll(model, ";", "")

		if len(model) < 4 {
			model = ""
		}
	}
  
	details := make(map[string]string, 0)
	if version != "" {
		details["os_version"] = version
	}

	if model != "" {
		details["model_version"] = model
	}

	if build != "" {
		details["build_version"] = build
	}

	if crios != "" {
		details["crios"] = crios
	}

	if safari != "" {
		details["safari"] = safari
	}

	return details
}