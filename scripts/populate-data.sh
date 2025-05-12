#!/bin/bash
set -e

# -----------------------------------
# Variables
# -----------------------------------
CLASS_LIB_NAME="Hyland.AcceptanceTests"
CLASS_LIB_EXTENSION=".csproj"
PROJECT_NAME="${CLASS_LIB_NAME}${CLASS_LIB_EXTENSION}"
VERSION="1.0.0"
AUTHOR="HYLAND SOFTWARE"
DESCRIPTION="A dynamically created NuGet package"
COMPANY="HYLAND"
OUTPUT_DIR="./nupkgs"

mkdir -p ./generated-sdk

echo "📦 Creating class library: $CLASS_LIB_NAME"

# -----------------------------------
# Create Class Library Project
# -----------------------------------
dotnet new classlib -n "$CLASS_LIB_NAME" --output ./generated-sdk

cd ./generated-sdk

if [ ! -f "$PROJECT_NAME" ]; then
  echo "❌ $PROJECT_NAME not found."
  exit 1
fi

echo "🔍 Checking for existing NuGet metadata..."

dotnet pack
ls ./bin/Release/