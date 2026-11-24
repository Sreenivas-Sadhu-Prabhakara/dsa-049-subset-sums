# All Subset Sums

> Striver SDE Sheet 049 · Day 9 Recursion — the take/skip recursion in its purest form: 2^n leaves, one sum each. Go, stdlib only, tested.

## Problem

Produce the sum of every subset of an array (all 2ⁿ of them, duplicates included), sorted. This is the warm-up that establishes the recursion pattern the rest of Day 9–10 builds on.

## Approach

At each index, branch twice: include the element in the running sum, or don't. The recursion tree therefore has exactly 2ⁿ leaves, each corresponding to one subset, and the running-sum parameter means no subset is ever materialized — only its sum reaches the leaf. Everything later in the recursion days (subsets with dedupe, combination sums, partitioning) is this skeleton plus pruning or bookkeeping.

**Complexity:** O(2ⁿ) time (inherent — that's the output size), O(n) recursion depth.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
