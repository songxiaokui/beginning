#!/bin/bash
set -euo pipefail
. ./__env.file

singularity run --contain \
				--bind "${PWD}" \
				--bind /tmp \
				--cleanenv \
				--env-file "" \
				--pwd "${PWD}" \
				"$1" \
				bash -euo pipefail "__cmder.sh"

