#!/bin/sh
set -e

ollama serve &
OLLAMA_PID=$!

echo "Waiting for Ollama..."
until ollama list >/dev/null 2>&1; do
  sleep 1
done

echo "Pulling planner model: ${PLANNER_MODEL}"
ollama pull "${PLANNER_MODEL}"

wait "$OLLAMA_PID"