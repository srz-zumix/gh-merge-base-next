# Multiple Merge Base Test Repository

This repository contains a complex commit graph designed to test scenarios
where `git merge-base --all` returns multiple commits.

## Commit Graph

```text
*   H (branch2) Final merge creating multiple merge bases
|\
| * G Branch2 specific development
* | F Branch1 specific development  
|\ \
| * E Common work on path A
* | D Common work on path B
|\ |
| * C Branch A initial work
* | B Branch B initial work  
|/
* A Initial commit (root)
```

This creates a diamond pattern where both C and B could be considered
as merge bases between the final branches, resulting in multiple
merge base commits when using `git merge-base --all`.
</content>
</invoke>