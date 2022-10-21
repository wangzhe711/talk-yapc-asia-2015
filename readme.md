## How to use

```bash
export tag=profiling
docker build -t=$tag . && docker run -it --rm -p=8082:8082 $tag zsh 

...
 ⚡ root  /demo   master ✚ ● ?  cd demo 
 ⚡ root  demo   master ✚ ● ?  web

```