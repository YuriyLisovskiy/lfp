## Configuration keywords explained here

Currently [lfp](https://github.com/YuriyLisovskiy/lfp) supports only yaml config file format.

Create a file, for example [`config.yml`](../sample/config.yml).
> Keywords can be added in any other order. 

Program's name (some licenses require program name, see license 
requirements [here](licenses.md)). If this field are not provided program name 
will be set to the project root directory name, **_optional_**:
```yaml
program_name: Skynet
```

Name(s) of a project's author(s), separated by comma(s), **_required_**:
```yaml
author: John Smith, Clint Eastwood
```

Year(s) when sources were developed, separated by comma(s), **_required_**:
```yaml
year: 2000, 2001
```

License type, the list of licenses marking is available [here](licenses.md), **_required_**:
```yaml
license: agpl-3.0
```

Set up project root path, **_required_**:
```yaml
project_root: /home/root/projects/skynet
```

Set up path(s) to file(s) that need to be licensed, `folder/...` means "include all 
children from the `folder` directory", **_required_**:
```yaml
paths:
  - parser/...
  - ai/...
  - execute.c
  - main.c
```

Add the path to custom license notice. If this field are not provided, default template 
will be used, custom license notice tutorial available [here](custom-notice.md), **_optional_**:
```yaml
custom_license_notice: /home/root/notices/license_notice_for_skynet.txt
```

Set `add_license_file` option to `true`, if it is required to add LICENSE file to the project root directory,
otherwise set it to `false`, **_optional_***:
```yaml
add_license_file: false
```

Set up `add_license_notice` to `true`, if you need to add license notice to each file,
provided in `paths`, otherwise set it to `false`, **_optional_***:
```yaml
add_license_notice: true
```

> * Setup at least one of `add_license_file` or `add_license_notice` options.
