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
  -package object \
  -visitor \
  -no-listener \
  -Xexact-output-dir \
  -o internal/generated/common \
  internal/infrastructure/antlr/grammars/Common.g4

antlr4 \
  -Dlanguage=Go \
  -package workflow \
  -visitor \
  -no-listener \
  -Xexact-output-dir \
  -o internal/generated/workflow \
  internal/infrastructure/antlr/grammars/Workflow.g4

antlr4 \
  -Dlanguage=Go \
  -package endpoint \
  -visitor \
  -no-listener \
  -Xexact-output-dir \
  -o internal/generated/endpoint \
  internal/infrastructure/antlr/grammars/Endpoint.g4

antlr4 \
  -Dlanguage=Go \
  -package consumer \
  -visitor \
  -no-listener \
  -Xexact-output-dir \
  -o internal/generated/consumer \
  internal/infrastructure/antlr/grammars/Consumer.g4

antlr4 \
  -Dlanguage=Go \
  -package application \
  -visitor \
  -no-listener \
  -Xexact-output-dir \
  -o internal/generated/application \
  internal/infrastructure/antlr/grammars/Application.g4

antlr4 \
  -Dlanguage=Go \
  -package service \
  -visitor \
  -no-listener \
  -Xexact-output-dir \
  -o internal/generated/service \
  internal/infrastructure/antlr/grammars/Service.g4

antlr4 \
  -Dlanguage=Go \
  -package deployment \
  -visitor \
  -no-listener \
  -Xexact-output-dir \
  -o internal/generated/deployment \
  internal/infrastructure/antlr/grammars/Deployment.g4
```