# User Agent Parser
## Overview
The Parser package is designed to extract device details from the user agent string for both Android and iOS devices. It can be used in applications to provide insights into the types of devices that are being used to access the application.

For **Android** devices, the `Parser` will return the following details:

- `OS Version`: This is the version of the Android operating system that the device is running. It will be extracted from the user agent string and returned as a string value.

- `Device Model Version`: This is the specific device model/version of the Android device that the user is using. It will be extracted from the user agent string and returned as a string value.

- `Build Version`: This is the build version of the Android operating system that the device is running. It will be extracted from the user agent string and returned as a string value.

For **iOS** devices, the `Parser` will return the following details:

- `CriOS Version`: This is the version of Google Chrome that is running on the iOS device. It will be extracted from the user agent string and returned as a string value.

- `Safari Version`: This is the version of the Safari browser that is running on the iOS device. It will be extracted from the user agent string and returned as a string value.

By using the Parser package, developers can gain a better understanding of the devices that are being used to access their applications. This information can be used to optimize the application for specific devices and to identify issues that are unique to certain devices

![Test](https://github.com/actionhero/node-resque/workflows/Test/badge.svg)

---

## Usage
Learn how to use the Parser package with an example.
```go
package main

import (
    "fmt"

    useragentparser "github.com/hardikshah197/user-agent-parser"
)

func main() {
    userAgentStrings := []string{
        // User agent string of Android
        "Mozilla/5.0 (Linux; Android 12; CPH2113 Build/SKQ1.210216.001;) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/111.0.5563.116 Mobile Safari/537.36 Zalo android/12100683 ZaloTheme/dark ZaloLanguage/vi",

        // User agent string of iOS
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/85 Version/11.1.1 Safari/605.1.15",
    }

    android_details := useragentparser.Parser(userAgentStrings[0])
    ios_details := useragentparser.Parser(userAgentStrings[1])

    fmt.Println("android device details: ", android_details)
    fmt.Println("ios device details: ", ios_details)
}
```

### Output:
```bash
android device details:  map[build_version:SKQ1.210216.001 model_version:CPH2113 os_version:12 safari:537.36]
ios device details:  map[crios:85 safari:605.1.15]
```

### Import Commond:
Package can easily be imported by running the following commond;
```sh
> go get github.com/hardikshah197/user-agent-parser
```

---

## Coverage
It has already covered all the case of User Agent String that the [google developer documentation](https://developer.chrome.com/docs/multidevice/user-agent/#chrome-for-ios) has provided till now.

___

## Implementation
This is a `Go` programming language code that provides a parser function which takes a user agent string as input and returns a map of details about the user agent. 

The function makes use of regular expressions to extract different types of information from the user agent string. 

Once the relevant information is extracted, it is stored in a map and returned as the output of the function. 

The details that can be extracted include the OS version, model version, build version, and details about the version of Safari or Chrome for iOS. 

Overall, the code provides a useful utility function for parsing and extracting information from user agent strings.

___

## Authors

[@hardikshah197](https://www.github.com/hardikshah197)