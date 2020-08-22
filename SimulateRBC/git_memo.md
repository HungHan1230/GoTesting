if the size of any file exceeds Github's file size limit of 100.00 MB, type the following command:
```
$ git filter-branch --tree-filter 'rm -rf SimulateRBC/nodes_jsons_sessions.json' HEAD
```

If you regret the commit you just made, type the following command to return to the previous commit:
```
# Find out what is the HEAD right now.
$ git log --oneline
95a0536 (HEAD -> master) After oral
fb62d09 complete most of functions
5f21102 (origin/master, origin/HEAD) complete the modularization
...

# After entering this command, the git HEAD will return to fb62d09
$ git reset 95a0536

```