[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b07a461e6a9a48fc84226baefff06423)](https://www.codacy.com/app/OpenDevSecOps/go-ses2email?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=opendevsecops/go-ses2email&amp;utm_campaign=Badge_Grade)
[![Follow on Twitter](https://img.shields.io/twitter/follow/opendevsecops.svg?logo=twitter)](https://twitter.com/opendevsecops)

# go-ses2email

Simple utility to send emails via AWS SES.

## Rationale

Although this task can be done with cURL, this command is meant to save you some troubles with the API and future-proof it.

## Getting Started

To install ses2email simply do:

```sh
$ go get github.com/opendevsecops/go-ses2email
```

Once the command is installed in your home go/bin folder execute it like this:

```sh
$ ~/go/bin/go-ses2email --help
```

To send message to slack simply do:

```sh
$ ~/go/bin/go-ses2email --from "from@domain.com" --to "to@domain.com" --subject "Your subject" --text "Your text message!"
```
