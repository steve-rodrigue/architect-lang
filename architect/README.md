# Antlr4 command
Generate the Go parser from the grammar:

```bash
antlr4 \
  -Dlanguage=Go \
  -package object \
  -visitor \
  -no-listener \
  -Xexact-output-dir \
  -o internal/generated/object \
  internal/infrastructure/antlr/Object.g4
```