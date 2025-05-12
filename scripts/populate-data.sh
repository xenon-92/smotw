#!/bin/bash
set -e

# -----------------------------------
# Variables
# -----------------------------------
CLASS_LIB_NAME="Hyland.AcceptanceTests"
CLASS_LIB_EXTENSION=".csproj"
PROJECT_NAME="${CLASS_LIB_NAME}${CLASS_LIB_EXTENSION}"
VERSION="11.0.0"
AUTHOR="HYLAND SOFTWARE"
DESCRIPTION="A dynamically created NuGet package"
COMPANY="HYLAND"

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

if grep -q "<PackageId>" "$PROJECT_NAME"; then
  echo "ℹ️ Package metadata already exists in $PROJECT_NAME."
else
  echo "✍️ Inserting NuGet metadata..."
  sed -i "/<PropertyGroup>/a \
  <PackageId>${CLASS_LIB_NAME}</PackageId>\n  <Version>${VERSION}</Version>\n  <Authors>${AUTHOR}</Authors>\n  <Company>${COMPANY}</Company>" "$PROJECT_NAME"
  echo "✅ Metadata inserted into $PROJECT_NAME."
fi

dotnet pack "$PROJECT_NAME" -c Release -v:normal -o ./
ls
package_name=$(find . -name "*.nupkg" -printf "%f\n" | head -n 1)
echo "package_name=$package_name" >> "$GITHUB_OUTPUT"

cd ..
file_path=$(find ./generated-sdk -name "*.nupkg" | head -n 1)
echo "package_path=$file_path" >> "$GITHUB_OUTPUT"