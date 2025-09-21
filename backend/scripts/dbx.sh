#!/bin/sh

program_name=$0
command=$1

# Get the directory where this script is located
script_dir=$(dirname "$0")
# Get the backend directory (parent of scripts directory)
backend_dir=$(dirname "$script_dir")

usage() {
  echo "Usage: $program_name COMMAND [ARGS...]"
  printf "\nCommands:\n"
  echo "  generate    - Run 'sqlc generate' in the backend directory"
  echo "  migrate     - Run goose migration commands"
  printf "\nExamples:\n"
  echo "  - Generate SQLC: $program_name generate"
  echo "  - Create migration: $program_name migrate create add_some_column -s sql"
  echo "  - Migrate up: $program_name migrate up"
  echo "  - Migrate down: $program_name migrate down"
  echo "  - Goose usage: $program_name migrate -h"
}

# Change to backend directory
cd "$backend_dir" || {
  echo "Error: Could not change to backend directory: $backend_dir"
  exit 1
}

if [ "$#" -lt 1 ]; then
  usage
  exit 1
fi

case "$command" in
  "generate")
    echo "Running 'sqlc generate' in $(pwd)..."
    if ! sqlc generate; then
      echo "Error: sqlc generate failed"
      exit 1
    fi
    echo "SQLC generation completed successfully"
    ;;
  "migrate")
    shift  # Remove 'migrate' from arguments
    echo "Running goose migration in $(pwd)/db/migrations..."
    if ! goose -dir "db/migrations" postgres \
      "host=$POSTGRES_HOST user=postgres password=$DB_PASSWORD dbname=$POSTGRES_DB port=5432" "$@"; then
      echo "Error: goose migration failed"
      exit 1
    fi
    ;;
  *)
    echo "Unknown command: $command"
    usage
    exit 1
    ;;
esac