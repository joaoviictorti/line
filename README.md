# line
line parse duplicate information
<p align="left">
	<a href="https://go.dev/"><img src="https://img.shields.io/badge/made%20with-go-blue"></a>
	<a href="#"><img src="https://img.shields.io/badge/platform-osx%2Flinux%2Fwindows-blueviolet"></a>
	<a href="https://github.com/joaoviictorti/line/releases"><img src="https://img.shields.io/github/release/joaoviictorti/line?color=blue"></a>
</p>

![line](/img/line.png)

line is a tool that aims to parse multiple lines. It is an excellent option to be used in penetration testing for application recognition and also in day-to-day activities. I designed `line` and maintained a consistently passive model to make it useful for pentest and IT professionals.

- [Installation](#installation)
- [features](#features)
- [Usage](#usage)
- [Details](#details)
- [Running line](#running-line)

---

# Features

- Perform duplicate row analysis
- Remove spaces between lines
- Remove empty lines
- Ignores case sensitivity

# Usage
- Sending the output to a file
```sh
cat urls.txt | line -f parse_urls.txt
```

- Ignoring case sensitivity
```sh
cat urls.txt | line -if parse_urls.txt
```

- Removing empty lines
```sh
cat urls.txt | line -ef parse_urls.txt
```

- Using all grouped flags
```sh
cat urls.txt | line -stif parse_urls.txt
```

# Details

![line](/img/help_line.png)


# Installation

line requires **golang** installed, to use:

```sh
go install -v github.com/joaoviictorti/line/cmd/line@latest
```

# Running line

![line](/img/exec.png)


<img width=100% src="https://capsule-render.vercel.app/api?type=waving&color=0000FF&height=120&section=footer"/>
