# Email Sender

![alt](/docs/images/gopher-email.jpeg?raw=true)
<p align="center">
<sup>If you're the owner of this image, please send us an email an we will remove it (if that's what you want) asap</sup>
</p>

---

## Index
1. [Introduction](#introduction)
2. [Setup](#setup)
3. [Environment Variables](#environmentvariables)
4. [API Usage](#apiusage)


---

## Introduction
This is just self-explanatory, `POST` an `/email` and then send it.

## Setup
1. Install Go (duh)
2. Set environment variables

## Environment Variables

These are the environment variables you need to make this awesome service work
``` shell
    EMAILSENDER_PORT    <PORT>
    EMAIL_USERNAME      <USERNAME>
    EMAIL_PASSWORD      <PASSWORD>
    EMAIL_SMTP          <SMTP>
```

#### Windows

Open up your command prompt and type
```pwsh
    setx ENVIRONMENT_VARIABLE_NAME  <VALUE>
```

#### Linux

Open up your `~/.bashrc` or `~/.bash_profile` (if you're using zsh maybe you should edit `~/.zshrc`) and add a line for each variable like this

``` shell
    export ENVIRONMENT_VARIABLE_NAME  <VALUE>
```

#### macOS

Same as Linux

## API Usage

Send a HTTP request using a `POST` method with the following body

```json
{
  "to": ["john@doe.com"],
  "template": "confirm-email",
  "data": {
    "name": "John Doe",
    "token": "F4K3-T0K3N"
  }
}
```

Data is a dynamic type, which will be created in `templatebodies.go` and later on reflect will parse the request data into the desired template
