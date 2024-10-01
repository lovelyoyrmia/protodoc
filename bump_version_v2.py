import subprocess
import re
import sys

def get_git_tags():
    """Fetch all git tags."""
    subprocess.run(["git", "fetch", "--tags"], check=True)
    result = subprocess.run(
        ["git", "tag"],
        stdout=subprocess.PIPE,
        text=True,
        check=True,
    )
    return result.stdout.splitlines()

def parse_version(version_str):
    """Parse version string into a tuple of (major, minor, patch, pre-release)."""
    match = re.match(r'^[vV]?(\d+)\.(\d+)\.(\d+)(.*)', version_str)
    if match:
        major, minor, patch, prerelease = match.groups()
        # Strip leading '.' and split pre-release by '.' if it exists
        prerelease = prerelease.lstrip('.')
        prerelease_parts = prerelease.split('.') if prerelease else []
        return (int(major), int(minor), int(patch), prerelease_parts)
    else:
        return None

def compare_versions(v1, v2):
    """Compare two version tuples."""
    # Compare major, minor, and patch first
    for a, b in zip(v1[:3], v2[:3]):
        if a != b:
            return (a > b) - (a < b)
    
    # If the major, minor, and patch are equal, compare pre-release parts
    if v1[3] and not v2[3]:
        return -1  # v1 is pre-release, v2 is not
    if not v1[3] and v2[3]:
        return 1   # v1 is not pre-release, v2 is
    # Compare pre-release parts
    return compare_pre_release(v1[3], v2[3])

def compare_pre_release(prerelease1, prerelease2):
    """Compare pre-release version strings."""
    len1, len2 = len(prerelease1), len(prerelease2)
    for i in range(max(len1, len2)):
        part1 = prerelease1[i] if i < len1 else ""
        part2 = prerelease2[i] if i < len2 else ""
        if part1.isdigit() and part2.isdigit():
            # Compare numeric parts
            if int(part1) != int(part2):
                return (int(part1) > int(part2)) - (int(part1) < int(part2))
    
        if part1 != part2:
            return (part1 > part2) - (part1 < part2)
    return 0  # Equal pre-release parts

def find_max_version(tags):
    """Find the maximum version from the list of tags."""
    max_version = None

    for tag in tags:
        version = parse_version(tag)
        if version and (max_version is None or compare_versions(version, max_version) > 0):
            max_version = version

    return max_version

def count_dev_tags(tags, base_version):
    """Count the number of development tags based on the base version."""
    pattern = re.compile(rf'^v{base_version}-dev\.(\d+)$')
    build_numbers = []

    for tag in tags:
        match = pattern.match(tag)
        if match:
            build_numbers.append(int(match.group(1)))

    return max(build_numbers) + 1 if build_numbers else 1

def generate_custom_version(bump_type, major, minor, patch, branch):
    """Generate a custom version based on the maximum version and the branch name."""
    if branch == 'main' or branch == 'master':
        if bump_type == "major":
            major += 1
            minor = 0
            patch = 0
        
        if bump_type == "minor":
            minor += 1
            patch = 0
        
        if bump_type == "patch":
            patch += 1
        
        base_version = f"{major}.{minor}.{patch}"
        return f"v{base_version}"
    
    base_version = f"{major}.{minor}.{patch}"
    build_number = count_dev_tags(get_git_tags(), base_version)
    return f"v{base_version}-dev.{build_number}"

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
    tags = get_git_tags()
    if not tags:
        print("No valid tags found.")
        return

    max_version = find_max_version(tags)

    if max_version is None:
        print("No valid version tags found.")
        return
    
    major = max_version[0]
    minor = max_version[1]
    patch = max_version[2]
    
    current_branch = get_current_branch()
    custom_version = generate_custom_version(bump_type, major, minor, patch, current_branch)
   
    print(f"Generated Version: {custom_version}")

    update_version_file(version_file, custom_version)

    print(f"Updated version to: {custom_version}")

    # Commit the changes
    subprocess.run(["git", "commit", "-am", f"Bump version to {custom_version}"])

    # Stage the changes
    subprocess.run(["standard-version", "--release-as", custom_version])

    subprocess.run(["git", "push", "origin", "HEAD", "--tags"])

    print("Pushed changes and tag to remote repository.")

if __name__ == "__main__":

    if len(sys.argv) != 2:
        bump_type = "patch"
    else:
        bump_type = sys.argv[1]
    main(bump_type)
