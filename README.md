# Aplikace EuroGoil - CTC

![logo.png](logo.png)

Příkaz pro sestavení a spuštění kontejneru, provedení simulace, uložení výsledků do souboru `output/output.yaml` a smazání kontejneru:
```docker
docker run --rm -it --mount type=bind,source="$(pwd)"/output,target=/app/output $(docker build -q .)
```