# scrambler

Scrambler aims to provide a simple way to let you secure sensible
information such as credentials you don't want to put into source control.

## Install

```
go get github.com/dbldots/scrambler
```

## Usage

## Secret config

There are to ways to let scrambler know what secret to use. See below.

### 1. Environment variable

```
export SCRAMBLER_SECRET=P@ssw0rd
```

### 2. Via config file

To use a config file for scrambler, create a file in your home directory called `.scrambler.yml`:

```
secret: P@ssw0rd
```

### Encrypt

To encrypt a value

```
scrambler encrypt [your-value]
```

This will print a line with your encrypted value. Copy and paste into your favorite editor.

### Decrypt

To decrypt (read) a file

```
scrambler read [your-file]
```

### Edit

#### Encrypt single values

To change values of a file that contains (or should contain) encrypted values:

```
scrambler edit [your-file]
```

This will open up the editor specified in EDITOR. An example:

the config file `config.yml` with content

```
one: one
two: SCRAMBLED:LIsPNA==
three: SCRAMBLED:L4twOw==
four:
  five: SCRAMBLED:L4twdA==
```

will be decrypted and opened in the editor in the following format:

```
one: one
two: SCRAMBLE:foo
three: SCRAMBLE:bar
four:
  five: SCRAMBLE:baz
```

Change anything after `SCRAMBLE:` to alter the values.
You can also encrypt existing values or add new values by using the `SCRAMBLE:[value]` pattern.

#### Encrypt whole file

Alternatively you can also choose to encrypt the entire content of a file (e.g. keyfiles).

Compared to the example for single values the `config.yml` file, if encrypted as a whole, looks like this:

```
:SCRAMBLED
FNQDLqFAbbdYvhYkSL7f5spxSJGr+94GXZ97rA6W7X5QHIb7j41Mp5qRB0cYl+bvrR8hsuoINk7gupO5esgPUA==%
```

You can edit the contents of the file via `scrambler edit [your-file]`, still.
