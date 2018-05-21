What is this?
---

This repository demonstrates the registry pattern in Golang by using interface factories. Dynamic configuration is demonstrated with deferred TOML loading. TOML was chosen as it is commonly used for tools configuration, though any configuration that allows deferred decoding (i.e. json.RawMessage) to cascade the configuration down to the plugins. It also demonstrates a common build pattern with Docker multi-stage builds and building the resulting binary into the scratch Docker container.

Running the Example
---

Download Docker and Docker Compose https://docs.docker.com/install/

Get Dep https://github.com/golang/dep

Run ```dep ensure```

Run ```docker-compose up```
