## Configuration keywords explained here

Currently [LFP](https://github.com/YuriyLisovskiy/lfp) supports only yaml config file format.
    
>*See license requirements for filed necessity. The default license notice template
requires <years>, <authors> and <license> fields.

Create a file, for example [`config.yml`](../sample/config.yml).
> Keywords can be added in any other order.

Program's name (some licenses require program name, see license 
requirements [here](licenses.md)). If this field is not provided, program name 
will be set to the project root directory name **_optional_**:
```yaml
program_name: Skynet
```

The program's name and a brief idea of what it does (some licenses require program description
instead of the program name, see license requirements [here](licenses.md))*:
```yaml
program_description: Sowftware 'Skynet' is developed for human extermination (nope)
```

Name(s) of a project's author(s)*:
```yaml
authors:
  - John Smith
  - Clint Eastwood
```

Year(s) when sources were developed*:
```yaml
years:
  - 2000-2001
```

License type, the list of licenses marking is available [here](licenses.md)*:
```yaml
license: agpl-3.0
```

Set up project root path*:
```yaml
project_root: /home/root/projects/skynet
```

Set up path(s) to file(s) that need to be licensed, `folder/...` means "include all 
children from the `folder` directory"*:
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

> * Setup at least one of `add_license_file` or `add_license_notice` options to `true`.
