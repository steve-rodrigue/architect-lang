FROM ollama/ollama:latest

ENV OLLAMA_HOST=0.0.0.0:11434
ENV PLANNER_MODEL=qwen2.5-coder:1.5b

COPY ./vllm/planner/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 11434

ENTRYPOINT ["/entrypoint.sh"]