import os
import re
import subprocess
import sys

def get_current_version():
    """Get the current version from Git tags."""
    try:
        version = subprocess.check_output(["git", "describe", "--tags", "--abbrev=0"]).strip().decode('utf-8')
        return version.lstrip('v')  # Remove the 'v' prefix for easier processing
    except subprocess.CalledProcessError:
        print("Error getting current version. Make sure you have at least one Git tag.")
        sys.exit(1)

def get_current_branch():
    """Get the current Git branch name."""
    try:
        branch = subprocess.check_output(["git", "rev-parse", "--abbrev-ref", "HEAD"]).strip().decode('utf-8')
        return branch
    except subprocess.CalledProcessError:
        print("Error getting current branch.")
        sys.exit(1)

def update_version_file(version_file, new_version):
    """Update the version.go file with the new version."""

    with open(version_file, 'r') as file:
        content = file.read()

    new_content = re.sub(r'const VERSION = ".*"', f'const VERSION = "{new_version}"', content)

    with open(version_file, 'w') as file:
        file.write(new_content)

def main(bump_type):
    version_file = "version.go"

    current_version = get_current_version()
    print(f"Current version: {current_version}")

    # Split the version into its components
    parts = current_version.split(".")

    if len(parts) != 3 and not parts[2].startswith("dev"):
        print(f"Unexpected version format: {current_version}")
        sys.exit(1)

    major, minor, patch = map(int, parts[0:3])

    # Check if the current version has a build number and extract it
    build_number = 0
    if "dev" in parts[2]:
        patch, build_number = parts[2].split("-dev.")
        patch = int(patch)
        build_number = int(build_number)
    else:
        build_number = 0  # Default build number if not present

    # Increment based on the bump type
    if bump_type == "major":
        major += 1
        minor = 0
        patch = 0
        build_number = 0  # Reset build number on major bump
    elif bump_type == "minor":
        minor += 1
        patch = 0
        build_number = 0  # Reset build number on minor bump
    elif bump_type == "patch":
        patch += 1
    else:
        print(f"Invalid bump type: {bump_type}. Use 'major', 'minor', or 'patch'.")
        sys.exit(1)

    # Get the current branch
    current_branch = get_current_branch()
    print(f"Current branch: {current_branch}")

    # Construct the new version
    if current_branch == "development":  # or whatever name you use for your development branch
        new_version = f"v{major}.{minor}.{patch}-dev.{build_number + 1}"
    else:
        new_version = f"v{major}.{minor}.{patch}"

    update_version_file(version_file, new_version)

    print(f"Updated version to: {new_version}")

    # Stage the changes
    subprocess.run(["standard-version", "--release-as", new_version])

    print("Pushed changes and tag to remote repository.")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python bump_version.py <bump_type>")
        print("bump_type: major, minor, or patch")
        sys.exit(1)

    bump_type = sys.argv[1]
    main(bump_type)
