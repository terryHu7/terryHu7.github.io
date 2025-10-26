#!/usr/bin/env bash
set -euo pipefail

# Unify Hugo post cover images to 16:9 JPG at a target resolution and size limit.
# - Crops center to 16:9, resizes to WIDTHxHEIGHT (default 1600x900)
# - Exports to JPEG with quality fallback to meet SIZE_LIMIT (default 500000 bytes)
# - Converts PNG covers to JPG and updates front matter (cover: and images list)
# - Optionally keeps a backup of the original as <name>@orig.<ext> when --backup is set
#
# Usage:
#   scripts/unify-covers.sh [options] [post_dir ...]
# Options:
#   -r ROOT       Root directory to scan (default: content/posts)
#   -w WIDTH      Target width (default: 1600)
#   -h HEIGHT     Target height (default: 900)
#   -l LIMIT      Size limit in bytes (default: 500000)
#   -b            Keep a backup <name>@orig.<ext> before processing
#   -H            Show help
#
# Examples:
#   scripts/unify-covers.sh
#   scripts/unify-covers.sh -b
#   scripts/unify-covers.sh -w 1920 -h 1080 -l 500000
#   scripts/unify-covers.sh content/posts/sep-1-7 content/posts/oct-6-12
#
# Note: Requires macOS 'sips' command.

ROOT="content/posts"
TARGET_WIDTH=1600
TARGET_HEIGHT=900
SIZE_LIMIT=500000
BACKUP="false"

QUALITY_PRIMARY="normal"
QUALITY_FALLBACK="low"

log() { echo "[unify-covers] $*"; }
die() { echo "[unify-covers] ERROR: $*" >&2; exit 1; }

if ! command -v sips >/dev/null 2>&1; then
  die "sips not found. This script requires macOS."
fi

show_help() {
  sed -n '1,60p' "$0" | sed -n '1,60p'
}

while getopts ":r:w:h:l:bH" opt; do
  case "$opt" in
    r) ROOT="$OPTARG" ;;
    w) TARGET_WIDTH="$OPTARG" ;;
    h) TARGET_HEIGHT="$OPTARG" ;;
    l) SIZE_LIMIT="$OPTARG" ;;
    b) BACKUP="true" ;;
    H) show_help; exit 0 ;;
    \?) die "Invalid option: -$OPTARG (use -H for help)" ;;
    :) die "Option -$OPTARG requires an argument" ;;
  esac
done
shift $((OPTIND-1))

crop_to_16x9() { # in out
  local in="$1" out="$2"
  local W H NEW_W NEW_H
  read W H < <(sips -g pixelWidth -g pixelHeight "$in" | awk '/pixelWidth/ {w=$2} /pixelHeight/ {h=$2} END{print w, h}')
  if [ -z "${W:-}" ] || [ -z "${H:-}" ]; then
    die "Failed to read dimensions for $in"
  fi
  if (( 9*W > 16*H )); then
    NEW_W=$(( H*16/9 ))
    NEW_H=$H
  elif (( 9*W < 16*H )); then
    NEW_H=$(( W*9/16 ))
    NEW_W=$W
  else
    NEW_W=$W
    NEW_H=$H
  fi
  sips --cropToHeightWidth "$NEW_H" "$NEW_W" "$in" --out "$out" >/dev/null
}

ensure_jpeg_size_limit() { # in out
  local in="$1" out="$2"
  # First attempt: normal quality at target size
  sips -s format jpeg -s formatOptions "$QUALITY_PRIMARY" -z "$TARGET_HEIGHT" "$TARGET_WIDTH" "$in" --out "$out" >/dev/null
  local size
  size=$(stat -f%z "$out")
  if (( size > SIZE_LIMIT )); then
    # Fallback: low quality at target size
    sips -s format jpeg -s formatOptions "$QUALITY_FALLBACK" -z "$TARGET_HEIGHT" "$TARGET_WIDTH" "$in" --out "$out" >/dev/null
    size=$(stat -f%z "$out")
    if (( size > SIZE_LIMIT )); then
      # Last fallback: 1280x720, low quality
      sips -s format jpeg -s formatOptions "$QUALITY_FALLBACK" -z 720 1280 "$in" --out "$out" >/dev/null
    fi
  fi
}

extract_cover_from_front_matter() { # index.md -> prints file name (may be empty)
  local index="$1"
  awk '
    BEGIN{inside=0}
    /^---[[:space:]]*$/{ if(inside==0){inside=1;next} else{inside=0} }
    inside && /^cover:[[:space:]]*/{
      gsub(/^cover:[[:space:]]*/, "", $0);
      gsub(/\"|'\''/, "", $0);
      print $0; exit
    }
  ' "$index"
}

update_front_matter_refs() { # index.md stem old_ext new_ext
  local index="$1" stem="$2" old_ext="$3" new_ext="$4"
  # Update 'cover:' field
  sed -i '' -E "s/^cover:[[:space:]]*${stem}\\.${old_ext}/cover: ${stem}.${new_ext}/" "$index" || true
  # Update images list entries referencing the cover (simple replacement)
  sed -i '' -E "s/^[[:space:]]*-[[:space:]]*${stem}\\.${old_ext}/- ${stem}.${new_ext}/" "$index" || true
}

process_cover_file() { # dir name
  local dir="$1" name="$2"
  local src="$dir/$name"
  if [ ! -f "$src" ]; then
    log "skip: $src not found"
    return 0
  fi
  local ext="${name##*.}"
  local stem="${name%.*}"
  local tmp_work tmp_crop
  tmp_work="$(mktemp -t unify-cover-work.XXXXXX)"
  tmp_crop="$(mktemp -t unify-cover-crop.XXXXXX)"
  cp "$src" "$tmp_work"
  if [ "$BACKUP" = "true" ]; then
    cp "$src" "$dir/${stem}@orig.$ext"
  fi
  crop_to_16x9 "$tmp_work" "$tmp_crop"
  local out_jpg="$dir/${stem}.jpg"
  ensure_jpeg_size_limit "$tmp_crop" "$out_jpg"
  # Remove original PNG if we converted
  if [ "$ext" = "png" ]; then
    rm -f "$src"
  fi
  local OW OH SH
  OW=$(sips -g pixelWidth "$out_jpg" | awk '/pixelWidth/ {print $2}')
  OH=$(sips -g pixelHeight "$out_jpg" | awk '/pixelHeight/ {print $2}')
  SH=$(ls -lh "$out_jpg" | awk '{print $5}')
  log "$out_jpg => ${OW}x${OH} (${SH})"
  rm -f "$tmp_work" "$tmp_crop"
}

process_post_dir() { # dir
  local dir="$1"
  local index="$dir/index.md"
  if [ ! -f "$index" ]; then
    log "skip: $dir (index.md not found)"
    return 0
  fi
  local cover_rel cover_name
  cover_rel="$(extract_cover_from_front_matter "$index" || true)"
  if [ -n "${cover_rel:-}" ] && [ -f "$dir/$cover_rel" ]; then
    cover_name="$cover_rel"
  else
    if [ -f "$dir/cover.jpg" ]; then
      cover_name="cover.jpg"
    elif [ -f "$dir/cover.png" ]; then
      cover_name="cover.png"
    else
      log "no cover found in $dir (front matter or cover.{jpg,png})"
      return 0
    fi
  fi
  local old_ext="${cover_name##*.}"
  local stem="${cover_name%.*}"
  process_cover_file "$dir" "$cover_name"
  update_front_matter_refs "$index" "$stem" "$old_ext" "jpg"
}

main() {
  if [ "$#" -gt 0 ]; then
    for d in "$@"; do
      [ -d "$d" ] || { die "Not a directory: $d"; }
      process_post_dir "$d"
    done
  else
    if [ ! -d "$ROOT" ]; then
      die "Root directory not found: $ROOT"
    fi
    for d in "$ROOT"/*; do
      [ -d "$d" ] || continue
      process_post_dir "$d"
    done
  fi
  log "Done."
}

main "$@"