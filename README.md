# ymldiff

Creates a diff report between two yaml files

## Usage

Download the latest release from the releases page

```bash
$ ymldiff from.yml to.yml
```

from.yml

```yml
- one
- two
- three
```

to.yml

```yml
- one
- two
- three
- four
```

will result in

```yml
type: replace
path: /3:before
value: 
  four
```

## Formatters

1. bosh
2. yaml
3. json