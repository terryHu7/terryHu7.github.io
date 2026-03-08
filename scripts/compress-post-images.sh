#!/usr/bin/env bash
set -euo pipefail

# Compress all images in a Hugo post bundle directory.
# - Cover image: crop to 16:9, resize to 1600x900, convert to JPEG
# - Other images: keep original aspect ratio, resize to max 1600px wide, convert to JPEG
# - Automatically updates all image references in index.md
# - Idempotent: skips directories that are already compressed
#
# Usage:
#   scripts/compress-post-images.sh <post_dir> [post_dir2 ...]
#   scripts/compress-post-images.sh content/posts/2026/feb-26-mar-8
#   scripts/compress-post-images.sh   # (no args) scans all posts under content/posts
#
# Note: Requires macOS 'sips' command.

SIZE_LIMIT=500000        # 500KB
COVER_WIDTH=1600
COVER_HEIGHT=900
MAX_WIDTH=1600           # max width for non-cover images
QUALITY_PRIMARY="normal"
QUALITY_FALLBACK="low"
ROOT="content/posts"

log() { echo "[compress] $*"; }
die() { echo "[compress] ERROR: $*" >&2; exit 1; }

if ! command -v sips >/dev/null 2>&1; then
  die "sips not found. This script requires macOS."
fi

# Get image dimensions
get_dimensions() {
  local file="$1"
  sips -g pixelWidth -g pixelHeight "$file" | awk '/pixelWidth/ {w=$2} /pixelHeight/ {h=$2} END{print w, h}'
}

# Crop to 16:9 center crop
crop_to_16x9() {
  local file="$1"
  local W H NEW_W NEW_H
  read W H < <(get_dimensions "$file")
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
  sips --cropToHeightWidth "$NEW_H" "$NEW_W" "$file" --out "$file" >/dev/null
}

# Convert to JPEG with size limit at exact dimensions (for cover)
to_jpeg_exact() {
  local in="$1" out="$2" w="$3" h="$4"
  sips -s format jpeg -s formatOptions "$QUALITY_PRIMARY" -z "$h" "$w" "$in" --out "$out" >/dev/null
  local size
  size=$(stat -f%z "$out")
  if (( size > SIZE_LIMIT )); then
    sips -s format jpeg -s formatOptions "$QUALITY_FALLBACK" -z "$h" "$w" "$in" --out "$out" >/dev/null
    size=$(stat -f%z "$out")
    if (( size > SIZE_LIMIT )); then
      sips -s format jpeg -s formatOptions "$QUALITY_FALLBACK" -z 720 1280 "$in" --out "$out" >/dev/null
    fi
  fi
}

# Convert to JPEG keeping aspect ratio, scale down if wider than MAX_WIDTH
to_jpeg_ratio() {
  local in="$1" out="$2"
  local W H
  read W H < <(get_dimensions "$in")

  if (( W > MAX_WIDTH )); then
    sips -s format jpeg -s formatOptions "$QUALITY_PRIMARY" --resampleWidth "$MAX_WIDTH" "$in" --out "$out" >/dev/null
  else
    sips -s format jpeg -s formatOptions "$QUALITY_PRIMARY" "$in" --out "$out" >/dev/null
  fi

  local size
  size=$(stat -f%z "$out")
  if (( size > SIZE_LIMIT )); then
    if (( W > MAX_WIDTH )); then
      sips -s format jpeg -s formatOptions "$QUALITY_FALLBACK" --resampleWidth "$MAX_WIDTH" "$in" --out "$out" >/dev/null
    else
      sips -s format jpeg -s formatOptions "$QUALITY_FALLBACK" "$in" --out "$out" >/dev/null
    fi
  fi
}

# Check if a post directory needs compression.
# Returns 0 (true) if there are non-JPEG images or oversized JPEGs.
needs_compression() {
  local dir="$1"
  # Check for non-JPEG images (png, webp, heic)
  for f in "$dir"/*.png "$dir"/*.PNG "$dir"/*.webp "$dir"/*.heic "$dir"/*.HEIC; do
    [ -f "$f" ] && return 0
  done
  # Check for oversized JPEGs
  for f in "$dir"/*.jpg "$dir"/*.jpeg "$dir"/*.JPG "$dir"/*.JPEG; do
    [ -f "$f" ] || continue
    local size
    size=$(stat -f%z "$f")
    if (( size > SIZE_LIMIT )); then
      return 0
    fi
  done
  return 1
}

# Process a single post directory
process_post_dir() {
  local POST_DIR="$1"
  local INDEX="$POST_DIR/index.md"

  if [ ! -d "$POST_DIR" ]; then
    log "Directory not found: $POST_DIR, skipping"
    return 0
  fi
  if [ ! -f "$INDEX" ]; then
    log "index.md not found in $POST_DIR, skipping"
    return 0
  fi

  # Idempotency: skip if nothing needs compression
  if ! needs_compression "$POST_DIR"; then
    log "Skipping $POST_DIR (already compressed)"
    return 0
  fi

  log "Processing $POST_DIR ..."

  # Detect cover filename from front matter
  local COVER_NAME=""
  COVER_NAME=$(awk '
    BEGIN{inside=0}
    /^---[[:space:]]*$/{ if(inside==0){inside=1;next} else{exit} }
    inside && /^cover:[[:space:]]*/{
      gsub(/^cover:[[:space:]]*/, "", $0);
      gsub(/\"|\047/, "", $0);
      gsub(/[[:space:]]+$/, "", $0);
      print $0; exit
    }
  ' "$INDEX")

  # Fallback: look for cover.* files
  if [ -z "$COVER_NAME" ]; then
    for ext in jpg jpeg png webp; do
      if [ -f "$POST_DIR/cover.$ext" ]; then
        COVER_NAME="cover.$ext"
        break
      fi
    done
  fi

  [ -n "$COVER_NAME" ] && log "  Cover image: $COVER_NAME" || log "  No cover image found"

  # Collect renames in a temp file (old_name|new_name per line) for bash 3.x compat
  local RENAME_LIST
  RENAME_LIST="$(mktemp -t compress-renames.XXXXXX)"

  # Process cover image
  if [ -n "$COVER_NAME" ] && [ -f "$POST_DIR/$COVER_NAME" ]; then
    local stem="${COVER_NAME%.*}"
    local ext="${COVER_NAME##*.}"
    local tmp_file
    tmp_file="$(mktemp -t compress-cover.XXXXXX)"
    cp "$POST_DIR/$COVER_NAME" "$tmp_file"

    crop_to_16x9 "$tmp_file"
    local out_file="$POST_DIR/${stem}.jpeg"
    to_jpeg_exact "$tmp_file" "$out_file" "$COVER_WIDTH" "$COVER_HEIGHT"

    # Remove original if different format
    if [ "$ext" != "jpeg" ] && [ -f "$POST_DIR/$COVER_NAME" ] && [ "$POST_DIR/$COVER_NAME" != "$out_file" ]; then
      rm -f "$POST_DIR/$COVER_NAME"
    fi

    echo "${COVER_NAME}|${stem}.jpeg" >> "$RENAME_LIST"
    local OW OH SH
    OW=$(sips -g pixelWidth "$out_file" | awk '/pixelWidth/ {print $2}')
    OH=$(sips -g pixelHeight "$out_file" | awk '/pixelHeight/ {print $2}')
    SH=$(ls -lh "$out_file" | awk '{print $5}')
    log "  [cover] ${stem}.jpeg => ${OW}x${OH} (${SH})"
    rm -f "$tmp_file"
  fi

  # Process all other images
  local img name stem ext out_file tmp_file OW OH SH
  for img in "$POST_DIR"/*.jpg "$POST_DIR"/*.jpeg "$POST_DIR"/*.png "$POST_DIR"/*.webp "$POST_DIR"/*.heic "$POST_DIR"/*.HEIC "$POST_DIR"/*.PNG "$POST_DIR"/*.JPG "$POST_DIR"/*.JPEG; do
    [ -f "$img" ] || continue

    name="$(basename "$img")"
    stem="${name%.*}"
    ext="${name##*.}"

    # Skip if this is the cover (already processed)
    [ "$name" = "$COVER_NAME" ] && continue
    # Skip if this is the already-processed cover output
    if [ -n "$COVER_NAME" ] && [ "$name" = "${COVER_NAME%.*}.jpeg" ] && [ "$name" != "$COVER_NAME" ]; then
      continue
    fi

    out_file="$POST_DIR/${stem}.jpeg"

    if [ "$img" = "$out_file" ]; then
      # Already jpeg — compress in place via temp file
      tmp_file="$(mktemp -t compress-img.XXXXXX)"
      cp "$img" "$tmp_file"
      to_jpeg_ratio "$tmp_file" "$out_file"
      rm -f "$tmp_file"
    else
      to_jpeg_ratio "$img" "$out_file"
      rm -f "$img"
    fi

    echo "${name}|${stem}.jpeg" >> "$RENAME_LIST"
    OW=$(sips -g pixelWidth "$out_file" | awk '/pixelWidth/ {print $2}')
    OH=$(sips -g pixelHeight "$out_file" | awk '/pixelHeight/ {print $2}')
    SH=$(ls -lh "$out_file" | awk '{print $5}')
    log "  ${stem}.jpeg => ${OW}x${OH} (${SH})"
  done

  # Update all references in index.md
  while IFS='|' read -r old_name new_name; do
    if [ "$old_name" != "$new_name" ]; then
      perl -pi -e "s/\Q${old_name}\E/${new_name}/g" "$INDEX"
      log "  index.md: $old_name -> $new_name"
    fi
  done < "$RENAME_LIST"

  rm -f "$RENAME_LIST"

  local count
  count=$(ls "$POST_DIR"/*.jpeg 2>/dev/null | wc -l | tr -d ' ')
  log "Done! $POST_DIR: ${count} images."
}

# Main
if [ "$#" -gt 0 ]; then
  for d in "$@"; do
    process_post_dir "$d"
  done
else
  if [ ! -d "$ROOT" ]; then
    die "Root directory not found: $ROOT"
  fi
  log "Scanning all posts under $ROOT ..."
  find "$ROOT" -name "index.md" -print0 | while IFS= read -r -d '' index_file; do
    process_post_dir "$(dirname "$index_file")"
  done
fi
