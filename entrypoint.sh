#!/bin/sh
# Set permissions for the Loki data directory
chmod -R 777 /tmp/loki

# Start Loki with the provided command
exec "$@"