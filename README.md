# kubedeploy  [![Build Status](https://travis-ci.com/bialas1993/kubedeploy.svg?token=ehLcPVdjAYYLjAbuBSAe&branch=master)](https://travis-ci.com/bialas1993/kubedeploy)
Deploy k8s manager for versioning deployments


### Usage
- convert:
```kubedeploy convert --templates=tpl --output=out --env=stage```

- apply:
```kubedeploy apply -f ./out/stage-a12345/pod.yaml```

### Instalation
 - macos
    ```shell
    wget -i https://github.com/bialas1993/kubedeploy/releases/download/v1.0.0/kubedeploy-macos
    ```
 - linux
    ```shell
    wget -i https://github.com/bialas1993/kubedeploy/releases/download/v1.0.0/kubedeploy-linux
    ```
    
### Cleaner

Not implemented yet.