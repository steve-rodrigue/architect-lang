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
  internal/infrastructure/antlr/grammars/Object.g4

antlr4 \
  -Dlanguage=Go \
  -package endpoint \
  -visitor \
  -no-listener \
  -Xexact-output-dir \
  -o internal/generated/endpoint \
  internal/infrastructure/antlr/grammars/Endpoint.g4
```