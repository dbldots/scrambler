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
scrambler encrypt --secret P@assw0rd [your-value]
```

.. then copy and paste the output into the file you want to share with others, e.g. a config file.

### Decrypt

To decrypt a file

```
scrambler decrypt --secret P@ssw0rd [your-file]
```

## Secret config

Instead of passing the secret via the `--secret` flag you can also provide the secret as..

### Environment variable

```
export SCRAMBLER_SECRET=P@ssw0rd
```

### Via config file

To use a config file for scrambler, create a file in your home directory called `.scrambler.yml`:

```
secret: P@ssw0rd
```

## Todo

* support secret environment variable
* support scrambler config file
* add support to manage (create/read/update) yaml keys
