<p align="center"><img src="img/dock.png"></p>

# Docker
> Docker is the leading Containers as a Service (CaaS) platform.

## Overview
Docker is the world’s leading software container platform. Developers use Docker to eliminate “works on my machine” problems when collaborating on code with co-workers. Operators use Docker to run and manage apps side-by-side in isolated containers to get better compute density. Enterprises use Docker to build agile software delivery pipelines to ship new features faster, more securely and with confidence for both Linux, Windows Server, and Linux-on-mainframe apps.

## What is a Container?

Containers are a way to package software in a format that can run isolated on a shared operating system. Unlike VMs, containers do not bundle a full operating system - only libraries and settings required to make the software work are needed. This makes for efficient, lightweight, self-contained systems and guarantees that software will always run the same, regardless of where it’s deployed.

## Docker For Developers

Docker automates the repetitive tasks of setting up and configuring development environments so that developers can focus on what matters: building great software.

Developers using Docker don’t have to install and configure complex databases nor worry about switching between incompatible language toolchain versions. When an app is dockerized, that complexity is pushed into containers that are easily built, shared and run. Onboarding a co-worker to a new codebase no longer means hours spent installing software and explaining setup procedures. Code that ships with Dockerfiles is simpler to work on: Dependencies are pulled as neatly packaged Docker images and anyone with Docker and an editor installed can build and debug the app in minutes.

## Autodock
For building:
> docker build -t <your username>/autodock
	
For running
> docker run -p 80:80 -e PORT=80 <your username>/autodock