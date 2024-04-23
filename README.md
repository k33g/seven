# seven
🤖7️⃣ The prompt runner - Resistance is futile


## Install Seven

```bash
VERSION="v0.4.2" OS="linux" ARCH="arm64"
wget -O capsule https://github.com/bots-garden/capsule/releases/download/${VERSION}/capsule-${VERSION}-${OS}-${ARCH}
chmod +x capsule
```



```bash
docker run \
--env SEVENCONFIG=./config/sevenconfig.yaml \
-v $(pwd)/robot:/robot \
-v $(pwd)/config:/config \
--rm k33g/seven:0.0.0 \
apply --manifest robot/01-simple.yaml
```
