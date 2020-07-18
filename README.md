<p align="center">
<a href="https://dscvit.com">
	<img src="https://user-images.githubusercontent.com/30529572/72455010-fb38d400-37e7-11ea-9c1e-8cdeb5f5906e.png" />
</a>
	<h2 align="center"> AutoDock </h2>
	<h4 align="center"> Continuous Docker Deployment Integration. <h4>
</p>

---

## What is Continuous Deployment?

Continuous Deploy is the process of automating publications in the production environment. The goal is not to accumulate code in staging.

## Instructions to run
```
$ https://github.com/GDGVIT/autodock.git
$ cd autodock
$ docker build -t <your username>/autodock .
$ docker run -p 80:80 -e PORT=80 <your username>/autodock
```

## Contributors
<table>
<tr align="center">


<td>

Balaji Jinnah

<p align="center">
<img src = "https://avatars3.githubusercontent.com/u/10010900?s=460&u=c95bb661dd058ef1d0759fb374321e68d257beab&v=4" width="150" height="150" alt="Balaji Jinnah">
</p>
</td>

<td>

Marcelo Cristiano

<p align="center">
<img src = "https://avatars1.githubusercontent.com/u/11621153?s=460&u=bebcd99a8d65cb537becc24ead2c926abb51f4c4&v=4" width="150" height="150" alt="Marcelo Cristiano">
</p>
</td>
</tr>
  </table>
  
## License
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

<p align="center">
	Made with :heart: by <a href="https://dscvit.com">DSC VIT</a>
</p>

