# 🍒 cherrybomb: a nicer way to cherry-pick commits

Cherrybomb is a simple CLI utility designed to help you cherry-pick all commits you have made on an upstream branch and bring them into your main branch instantly.

To do this, we introduce a concept of "cherrybombing" which involves the following key steps:

1. First, we gather all commits you have authored inside of an upstream branch, except for merge commits. This assumes you have authored all commits on a particular feature or fix branch and such commits will contain the necessary changes you wish to migrate to your target branch

2. Commits are previewed and you will be able to manually verify they are as expected.

3. The cherrybomb begins - each commit hash in the preview is cherrypicked and you are notified of any conflicts in the process.

Running cherrybomb is as simple as `cherrybomb <Source Branch in Upstream>`
