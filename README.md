<img width=100% src="https://capsule-render.vercel.app/api?type=waving&color=0000FF&height=120&section=header"/>

![hunteheader](/img/line.png)

<p align="center">
	<a href="https://www.python.org/"><img src="https://img.shields.io/badge/made%20with-go-blue"></a>
	<a href="#"><img src="https://img.shields.io/badge/platform-osx%2Flinux%2Fwindows-blueviolet"></a>
	<a href="https://github.com/joaoviictorti/line/releases"><img src="https://img.shields.io/github/release/joaoviictorti/line?color=blue"></a>
</p>

<h4 align="center">line parse duplicate information</h4>


<p align="center">
  <a href="#características">Features</a> •
  <a href="#instalação">Installation</a> •
  <a href="#how-to-use"> How to use</a> •
  <a href="#details">Details</a> •
  <a href="#running-line">Running line</a>  
</p>

---


line is a tool that aims to parse multiple lines. It is an excellent option to be used in penetration testing for application recognition and also in day-to-day activities.

I designed `line` and maintained a consistently passive model to make it useful for pentest and IT professionals.

# Features

- Perform duplicate row analysis
- Remove spaces between lines
- Remove empty lines
- Ignores case sensitivity

# How to use

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
