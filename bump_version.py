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

    build_number = 0
    if len(parts) == 4:
        # The fourth part is build_number, and patch might be in the format "1-dev"
        major, minor = map(int, parts[0:2])
        
        # Handle patch if it contains a dash
        patch_parts = parts[2].split("-")
        patch = int(patch_parts[0])  # Extract the numeric part of patch
        
        build_number = int(parts[3])
    else:
        major, minor, patch = map(int, parts[0:3])

    # Get the current branch
    current_branch = get_current_branch()
    print(f"Current branch: {current_branch}")
    
    # Increment based on the bump type
    if bump_type == "major":
        major += 1
        minor = 0
        patch = 0
    
    if bump_type == "minor":
        minor += 1
        patch = 0
    
    if bump_type == "patch":
        patch += 1
    
    if current_branch == "development" or bump_type == "dev":
        build_number += 1

    # Construct the new version
    if current_branch == "development":  # or whatever name you use for your development branch
        new_version = f"{major}.{minor}.{patch}-dev.{build_number}"
    else:
        new_version = f"{major}.{minor}.{patch}"

    update_version_file(version_file, new_version)

    print(f"Updated version to: {new_version}")

    # Commit the changes
    subprocess.run(["git", "commit", "-am", f"Bump version to v{new_version}"])

    # Stage the changes
    subprocess.run(["standard-version", "--release-as", new_version])

    subprocess.run(["git", "push", "origin", "HEAD", "--tags"])

    print("Pushed changes and tag to remote repository.")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python bump_version.py <bump_type>")
        print("bump_type: major, minor, or patch")
        sys.exit(1)

    bump_type = sys.argv[1]
    main(bump_type)
