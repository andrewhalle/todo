# Outline

keep track of items in a todo list, and sort them intelligently based on due date, time to complete, time already spent, etc. have a command-line git-like interface for adding, removing, marking complete/blocked/..., tasks.

# Design

store tasks as text files in a .todo directory, keep a metafile in .todo with the sorted order of the tasks according to the currently active scheduling algorithm.

# Ideas

 - if there is a .todo in the working directory, use that by default. Otherwise, look for a .todo directory in the user's home directory. If that doesn't exist, create one if it makes sense. there should be flags to explicitly look at the local or global .todo, which should error if it doesn't exist. Another flag should allow specification of a path to the .todo, which should error if it doesn't exist

 - todo list should print tasks in sorted order, and say which scheduling algorithm was used to sort them. random order should be an option, as well as alphabetical (for the truly crazy) and by date entered.

 - task dependency should be supported (not required), and when used should not list tasks in an order which would not allow them to be completed (or should warn the user when conflicts may arise)

