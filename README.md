# scrambler

Scrambler aims to provide a simple way to let you secure sensible
information such as credentials you don't want to put into source control.

## Install

```
go get github.com/dbldots/scrambler
```

## Usage

### Encrypt

To encrypt a value

```
scrambler encrypt --secret P@assw0rd --value [your-value]
```

.. then copy and paste the output into the file you want to share with others, e.g. a config file.

### Decrypt

To decrypt a file

```
scrambler decrypt --secret P@ssw0rd --file [your-file]
```

## Todo

* support secret environment variable
* support scrambler config file
* add support to manage (create/read/update) yaml keys
