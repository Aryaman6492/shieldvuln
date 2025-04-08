## Docker Build

### Build your own Docker image

1. Clone Project
```
git clone https://github.com/Aryaman6492/shieldvuln.git shieldvuln && cd "$_"
```

2. Build
```
docker buildx build -t shieldvuln -f build/Dockerfile --load .
```
