# Gopher 道場 #8 課題 1

## How to Build

```sh
make build
```

## Usage

```sh
> conv --help
```

```console
Usage of conv:
  -dir string
        Path to the target directory
  -from string
        Image extension before conversion (default "jpg")
  -to string
        Image extension after conversion (default "png")
  -verbose
        Show conversion logs
```

## 感想

- 変換を並列で処理できたら良さそうでした。
- `type ImageExtension string` を定義したが、`flags` から取得する拡張子が `string` なので、よしなにキャストする方法が分かりませんでした。
