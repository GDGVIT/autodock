# AutoDock

Continuous Docker Deployment Integration.

#### What is Continuous Deployment?

Continuous Deploy is the process of automating publications in the production environment. The goal is not to accumulate code in staging.

## Developing AutoDock

The easiest way to run autodock is with docker.

You can learn how to install docker here:
    
    https://docs.docker.com/engine/installation/

Now, with docker:

    https://github.com/GDGVIT/autodock.git
    cd autodock
    
    docker build -t <your username>/autodock .
    docker run -p 80:80 -e PORT=80 <your username>/autodock
