## Create custom license notice

Create file which will contain custom license notice, for example `license-notice.txt`,
and setup `custom_license_notice` field in configuration file. All text from `license-notice.txt`
will be added to the all files, provided in `paths` field in configuration file.

This is the default license notice which will be used if `custom_license_notice` field
are not provided*:
```text
// Copyright (c) <years> <authors>
// Distributed under the <license name>,
// see the accompanying file LICENSE or <license link>
```

##### List of available keywords:
* `<years>`
* `<authors>`
* `<comment>`
* `<program name>`
* `<license name>`
* `<license link>`
* `<program description>`

>In this example C-language comments are used temporarily. The program automatically
determine what the language is used and choose suitable comments for license notice.
>
>NOTE: this feature is in development now!
