# Load Testing Lightning Talk

Here you'll find code and slides from my lightning talk at Stavanger Software Developers Meetup 26. April 2022 on Load Testing.

Slides: https://github.com/StianOvrevage/ssdm-april22-loadtesting/blob/main/Slides%20-%20SSDM%20-%20Load%20Testing%20-%20April%202022.pdf

Video: https://www.youtube.com/watch?v=FKKxYGIePRM

# Installing all tools, golang, k6, locust, influxdb and grafana

My environment:
 - Windows 10
 - WSL2 (Windows Subsystem for Linux)
 - Ubuntu 20.04 in WSL2

If your environment differs substantially the installation script and commands might be different.
Refer to the documentation in the various subdirectories in the repo for more information.

Clone the repo:

    git clone git@github.com:StianOvrevage/ssdm-april22-loadtesting.git

Install all the tools (you might want to do a `sudo ls` in advance to avoid `sudo` in script asking for password and breaking it)

    bash installtools.sh

PS: This will update your golang install to 1.18 and remove older versions if you have it installed. Go is generally extremely backwards-compatible so should not pose an issue!

PS: The script does not add `go` to your `PATH` so either you need to add it to your `.bashrc` manually or execute `export PATH=$PATH:/usr/local/go/bin` whenever you open a new console.

# Start influxdb and grafana

    bash starttools.sh
