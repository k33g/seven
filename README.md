# 🤖7️⃣ Seven
> Define **Ollama** prompts with yaml, run them with **Seven**

> *Resistance is futile*

## Install Seven

Download the binariy version:
```bash
VERSION="v0.0.0" OS="darwin" ARCH="arm64"
wget -O seven https://github.com/k33g/seven/releases/download/${VERSION}/seven-${OS}-${ARCH}
chmod +x seven
```
> - Find a version for a specific target architecture on the release page: https://github.com/k33g/seven/releases/tag/v0.0.0
> - Seven exists for `darwin/amd64`, `darwin/arm64`, `linux/amd64`, `linux/arm64`

Or use docker: `docker pull k33g/seven:0.0.0`

## Configuration

Create a yaml configuration file. For example:

> `./config/sevenconfig.yaml`
```yaml
config:
  ollama:
    url: http://host.docker.internal:11434
```
> - If you run **seven** from a container and Ollama from the host, the url of the Ollama API is `http://host.docker.internal:11434`
> - If you run **seven** locally and Ollama locally (from the host), the url of the Ollama API is `http://localhost:11434`


## Run a yaml prompt with Seven

### Create a prompt manifest

> `./robot/01-simple.yaml`
```yaml
model:
  name: phi3
  memory: false
  stream: true

prompt:
  settings:
    temperature: 1

  human: |
    Who is James T. Kirk?
```

If needed, pull the LLM:
```bash
ollama pull phi3
```

Then apply the prompt manifest with Seven:
```bash
export SEVENCONFIG=$(pwd)/config/sevenconfig.yaml
./seven apply --manifest robot/01-simple.yaml
```

Or you can apply it with the Seven Docker image:
```bash
docker run \
--env SEVENCONFIG=./config/sevenconfig.yaml \
-v $(pwd)/robot:/robot \
-v $(pwd)/config:/config \
--rm k33g/seven:0.0.0 \
apply --manifest robot/01-simple.yaml
```

#### Remarks

- You can override the human question of the prompt without modifying the  manifest with the `--question` option flag:
  ```bash
  ./seven apply --manifest robot/01-simple.yaml --question "Who is Jean-Luc Picard?"
  ```
- You can specify the path of the config file with the `--config` option flag.
- You can create a text file with the result of the completion with the `--output` option flag.

## Add a system message to the manifest

> `./robot/02-prompt-system.yaml`
```yaml
model:
  name: phi3
  memory: false
  stream: true

prompt:
  system: |
    You are an AI assistant. Your name is Seven. 
    You are an expert in Star Trek.
    All questions are about Star Trek.
  human: |
    Who is James T. Kirk?
```

Apply the prompt manifest:
```bash
./seven apply --manifest robot/02-prompt-system.yaml
```

### Remarks

- You can override the system message of the manifest with the `--system` option flag.

```bash
./seven apply --manifest robot/02-prompt-system.yaml \
--question "Who is Spiderman?"

# as the AI agent is a Star Trek expert, 
# "he" will answer that he doesn't know Spiderman, 
# so change the system message:
./seven apply --manifest robot/02-prompt-system.yaml \
--system "you are a Marvel expert" \
--question "Who is Spiderman?"
```

## 🚧 This is a work in progress

More documentation to come. But you can have a look to the `/robot` directory to get more examples.
